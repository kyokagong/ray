package task

import (
	"ray/internal/runtime/common"
	"ray/internal/runtime/generated"
)

type LocalTaskSubmitter struct {
	rayFunc RayFunction
	args    []interface{}
	// waitingTasks Map[api.ObjectRef]TaskSpec
}

// Create task spec and submit
func (localTaskSubmitter LocalTaskSubmitter) SubmitTask() string {
	taskSpec := buildTaskSpec(common.TaskType_NORMAL_TASK, localTaskSubmitter.rayFunc, localTaskSubmitter.args)
	localTaskSubmitter.submitTaskSpec(taskSpec)
	return localTaskSubmitter.GetReturnIds(taskSpec)
}

func (localTaskSubmitter LocalTaskSubmitter) submitTaskSpec(taskSpec generated.TaskSpec) {

}

func (localTaskSubmitter LocalTaskSubmitter) GetReturnIds(taskSpec generated.TaskSpec) string {
	return "test"
}

func CreateLocalTaskSubmitter(rayFunc RayFunction, args []interface{}) LocalTaskSubmitter {
	return LocalTaskSubmitter{rayFunc, args}
}

// Build task spec
func buildTaskSpec(taskType common.TaskType, rayFunc RayFunction, args []interface{}) generated.TaskSpec {

	jobId := []byte("testJobId")
	taskId := []byte("testTaskId")
	skipException := false
	return generated.TaskSpec{
		generated.TaskType_NORMAL_TASK,
		rayFunc.functionDescriptor.functionName,
		generated.Language_GO,
		generated.GoFunctionDescriptor{rayFunc.functionDescriptor.functionName},
		jobId,
		taskId,
	ParentTaskId                    []byte                 `protobuf:"bytes,7,opt,name=parent_task_id,json=parentTaskId,proto3" json:"parent_task_id,omitempty"`
	ParentCounter                   uint64                 `protobuf:"varint,8,opt,name=parent_counter,json=parentCounter,proto3" json:"parent_counter,omitempty"`
	CallerId                        []byte                 `protobuf:"bytes,9,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	CallerAddress                   *Address               `protobuf:"bytes,10,opt,name=caller_address,json=callerAddress,proto3" json:"caller_address,omitempty"`
		generated.TaskArg{},
		1,
	RequiredResources               map[string]float64     `protobuf:"bytes,13,rep,name=required_resources,json=requiredResources,proto3" json:"required_resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	RequiredPlacementResources      map[string]float64     `protobuf:"bytes,14,rep,name=required_placement_resources,json=requiredPlacementResources,proto3" json:"required_placement_resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	ActorCreationTaskSpec           *ActorCreationTaskSpec `protobuf:"bytes,15,opt,name=actor_creation_task_spec,json=actorCreationTaskSpec,proto3" json:"actor_creation_task_spec,omitempty"`
	ActorTaskSpec                   *ActorTaskSpec         `protobuf:"bytes,16,opt,name=actor_task_spec,json=actorTaskSpec,proto3" json:"actor_task_spec,omitempty"`
		1,
	PlacementGroupId                []byte                 `protobuf:"bytes,18,opt,name=placement_group_id,json=placementGroupId,proto3" json:"placement_group_id,omitempty"`
	PlacementGroupBundleIndex       int64                  `protobuf:"varint,19,opt,name=placement_group_bundle_index,json=placementGroupBundleIndex,proto3" json:"placement_group_bundle_index,omitempty"`
	PlacementGroupCaptureChildTasks bool                   `protobuf:"varint,20,opt,name=placement_group_capture_child_tasks,json=placementGroupCaptureChildTasks,proto3" json:"placement_group_capture_child_tasks,omitempty"`
	OverrideEnvironmentVariables    map[string]string      `protobuf:"bytes,21,rep,name=override_environment_variables,json=overrideEnvironmentVariables,proto3" json:"override_environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
		skipException,
	DebuggerBreakpoint              []byte                 `protobuf:"bytes,23,opt,name=debugger_breakpoint,json=debuggerBreakpoint,proto3" json:"debugger_breakpoint,omitempty"`
	SerializedRuntimeEnv            string     
	}
}
