package task

import (
	"bytes"
	"sync"

	"github.com/ray-project/ray/go/internal/logger"
	"github.com/ray-project/ray/go/internal/runtime/generated"
	"github.com/ray-project/ray/go/internal/runtime/object"
	"github.com/ray-project/ray/go/internal/runtime/serializer"
)

type LocalTaskSubmitter struct {
	TaskExecutor   *LocalTaskExecutor
	TaskObjRefMap  map[string]object.ObjectRef
	ExecuteService *GoroutinePool
	ObjectStore    *object.LocalObjectStore
}

type RunnableTask func()

// GoroutinePool 是一个goroutine池
type GoroutinePool struct {
	taskChan chan RunnableTask
	wg       sync.WaitGroup
}

// NewGoroutinePool 创建一个新的goroutine池
func NewGoroutinePool(maxGoroutines int) *GoroutinePool {
	pool := &GoroutinePool{
		taskChan: make(chan RunnableTask),
	}
	pool.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			defer pool.wg.Done()
			for task := range pool.taskChan {
				task()
			}
		}()
	}
	return pool
}

// Run 提交一个任务到goroutine池
func (p *GoroutinePool) Run(task RunnableTask) {
	p.taskChan <- task
}

// Shutdown 等待所有goroutine停止
func (p *GoroutinePool) Shutdown() {
	close(p.taskChan)
	p.wg.Wait()
}

// Create task spec and submit
func (submitter *LocalTaskSubmitter) SubmitTask(remoteFunctionHolder *RemoteFunctionHolder, args []interface{}) (object.ObjectRef, error) {
	rayFunction, err := GetFunctionManager().GetRayFunction(remoteFunctionHolder.FuncWrapper.FuncName)
	if err != nil {
		return object.ObjectRef{}, err
	}
	taskSpec, err := buildTaskSpec(generated.TaskType_NORMAL_TASK, rayFunction, args)
	if err != nil {
		return object.ObjectRef{}, err
	}
	submitter.submitTaskSpec(taskSpec)
	return submitter.GetReturnObjefRef(taskSpec), nil
}

func (submitter *LocalTaskSubmitter) submitTaskSpec(taskSpec *generated.TaskSpec) {
	runnableTask := func() {
		funcName := taskSpec.GetFunctionDescriptor().GetGoFunctionDescriptor().FunctionName

		var args []interface{}
		for _, tArg := range taskSpec.Args {
			arg, err := serializer.Deserialize(tArg.GetData())
			if err != nil {
				logger.Errorf("runnableTask Deserialize error: %v", err)
				// put error in object store
			}
			args = append(args, arg)
		}
		result := submitter.TaskExecutor.ExecuteTask(funcName, args)
		objRef := submitter.GetReturnObjefRef(taskSpec)
		submitter.ObjectStore.PutRaw(objRef, result.Value)
	}
	submitter.ExecuteService.Run(runnableTask)
}

func (submitter *LocalTaskSubmitter) GetReturnObjefRef(taskSpec *generated.TaskSpec) object.ObjectRef {
	strTaskId := string(taskSpec.GetTaskId())
	objRef, found := submitter.TaskObjRefMap[strTaskId]
	if !found {
		objRef = object.NewObjectRef()
		submitter.TaskObjRefMap[strTaskId] = objRef
	}
	return objRef
}

func NewLocalTaskSubmitter(objectStore *object.LocalObjectStore) LocalTaskSubmitter {
	executeService := NewGoroutinePool(10)
	return LocalTaskSubmitter{
		NewLocaLocalTaskExecutor(objectStore),
		make(map[string]object.ObjectRef),
		executeService,
		objectStore,
	}
}

// Build task spec
func buildTaskSpec(taskType generated.TaskType, rayFunc RayFunction, args []interface{}) (*generated.TaskSpec, error) {

	jobId := []byte("testJobId")
	taskId := []byte("testTaskId")
	var taskArgs []*generated.TaskArg
	for _, arg := range args {
		datum, err := serializer.Serialize(arg)
		if err != nil {
			return nil, err
		}
		taskArgs = append(taskArgs, &generated.TaskArg{
			ObjectRef: &generated.ObjectReference{},
			Data:      datum,
			Metadata:  new(bytes.Buffer).Bytes(),
		})
	}

	funcDescripter := generated.FunctionDescriptor{
		FunctionDescriptor: &generated.FunctionDescriptor_GoFunctionDescriptor{
			GoFunctionDescriptor: &generated.GoFunctionDescriptor{
				FunctionName: rayFunc.FunctionDescriptor.functionName,
			},
		},
	}
	return &generated.TaskSpec{
		Type:                       taskType,
		Name:                       "",
		Language:                   generated.Language_GO,
		FunctionDescriptor:         &funcDescripter,
		JobId:                      jobId,
		TaskId:                     taskId,
		ParentTaskId:               new(bytes.Buffer).Bytes(),
		ParentCounter:              0,
		CallerId:                   new(bytes.Buffer).Bytes(),
		CallerAddress:              &generated.Address{},
		Args:                       taskArgs,
		NumReturns:                 1,
		RequiredResources:          make(map[string]float64),
		RequiredPlacementResources: make(map[string]float64),
		ActorCreationTaskSpec:      nil,
		ActorTaskSpec:              nil,
		MaxRetries:                 3,
	}, nil
}
