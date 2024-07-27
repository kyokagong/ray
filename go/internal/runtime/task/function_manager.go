package task

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

var funcManagerInitOnce sync.Once

// Function descriptor for GO.
type GoFunctionDescriptor struct {
	functionName string
}

// Represents a Ray function
type RayFunction struct {
	FunctionDescriptor GoFunctionDescriptor
	FunctionHolder     RemoteFunctionHolder
}

func (rf *RayFunction) Call(args ...interface{}) (interface{}, error) {
	var inArgs []reflect.Value
	argsType := rf.FunctionHolder.FuncWrapper.InArgsType
	for i, arg := range args {
		inType := argsType[i]
		var inArg reflect.Value
		switch inType.Kind() {
		case reflect.String:
			inArg = reflect.ValueOf(arg.(string))
		case reflect.Int:
			inArg = reflect.ValueOf(arg.(int))
		case reflect.Int8:
			inArg = reflect.ValueOf(arg.(int8))
		case reflect.Int16:
			inArg = reflect.ValueOf(arg.(int16))
		case reflect.Int32:
			inArg = reflect.ValueOf(arg.(int32))
		case reflect.Int64:
			inArg = reflect.ValueOf(arg.(int64))
		case reflect.Bool:
			inArg = reflect.ValueOf(arg.(bool))
		default:
			inArg = reflect.ValueOf(arg)
		}
		inArgs = append(inArgs, inArg)
	}

	// 调用 f 并获取返回值
	result := rf.FunctionHolder.FuncWrapper.AnyFunc(inArgs)
	if result.IsValid() {
		return result.Interface(), nil
	}
	// 将返回的 reflect.Value 转换为 interface{}
	return nil, errors.New(fmt.Sprintf("func: %s, return is not valid", rf.FunctionHolder.FuncWrapper.FuncName))
}

// Go FunctionManager that manages function create a
type FunctionManager struct {
	funcHolderMap     map[string]RemoteFunctionHolder
	funcDescriptorMap map[string]GoFunctionDescriptor
}

type FunctionWrapper struct {
	FuncName   string
	AnyFunc    FuncType1
	InArgsType []reflect.Type
}

func NewFunctionWrapper(fn interface{}) FunctionWrapper {
	return FunctionWrapper{
		GetFunctionName(fn),
		BuildAnyFunc(fn),
		GetFunInArgsType(fn),
	}
}

func BuildAnyFunc(fn interface{}) FuncType1 {
	vfn := reflect.ValueOf(fn)
	return func(args []reflect.Value) reflect.Value {
		ret := vfn.Call(args)
		if ret != nil {
			return ret[0]
		}
		return reflect.Value{}
	}
}

func GetFunInArgsType(fn interface{}) []reflect.Type {
	funcType := reflect.TypeOf(fn)

	numIn := funcType.NumIn()

	var argsType []reflect.Type
	for i := 0; i < numIn; i++ {
		inType := funcType.In(i)
		argsType = append(argsType, inType)
	}
	return argsType
}

func (functionManager FunctionManager) GetFunction(funcName string) RemoteFunctionHolder {
	return RemoteFunctionHolder{}
}

// Get a function name
func GetFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

// Register remote function to FunctionManager
func (functionManager *FunctionManager) RegisterRemoteFunction(funcWrapper FunctionWrapper) {
	funcName := funcWrapper.FuncName
	if _, found := functionManager.funcHolderMap[funcName]; !found {
		funcHolder := BuildRemoteFunctionHolder(funcWrapper)
		functionManager.funcHolderMap[funcHolder.FuncWrapper.FuncName] = funcHolder
		functionManager.funcDescriptorMap[funcName] = functionManager.BuildFunctionDescriptor(funcWrapper)
	}
}

// Serialize user function
func (functionManager *FunctionManager) GetRayFunction(funcName string) (RayFunction, error) {
	funcHolder, found := functionManager.funcHolderMap[funcName]
	if !found {
		return RayFunction{}, errors.New(fmt.Sprintf("funcName %s is not found", funcName))
	}
	functionDescriptor := functionManager.funcDescriptorMap[funcHolder.FuncWrapper.FuncName]
	return RayFunction{
		FunctionDescriptor: functionDescriptor,
		FunctionHolder:     funcHolder,
	}, nil
}

// Get GoFunctionDescriptor
func (functionManager *FunctionManager) BuildFunctionDescriptor(funcWrapper FunctionWrapper) GoFunctionDescriptor {
	functionDescriptor := GoFunctionDescriptor{funcWrapper.FuncName}
	return functionDescriptor
}

type FuncType1 func(args []reflect.Value) reflect.Value

type RemoteFunctionHolder struct {
	FuncWrapper FunctionWrapper
}

func BuildRemoteFunctionHolder(funcWrapper FunctionWrapper) RemoteFunctionHolder {
	holder := RemoteFunctionHolder{funcWrapper}
	return holder
}

var (
	funcManager *FunctionManager
)

func GetFunctionManager() *FunctionManager {
	funcManagerInitOnce.Do(func() {
		funcManager = &FunctionManager{
			make(map[string]RemoteFunctionHolder),
			make(map[string]GoFunctionDescriptor),
		}
	})
	return funcManager
}
