package task

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"runtime"
)

// Function descriptor for GO.
type GoFunctionDescriptor struct {
	functionName    string
	functionContent []byte
}

// Represents a Ray function
type RayFunction struct {
	functionDescriptor GoFunctionDescriptor
}

// Go FunctionManager that manages function create a
type FunctionManager struct {
	funcNameMap map[string]FuncType1
}

type FunctionWrapper struct {
	recall func(interface{}) interface{}
}

// Return static FunctionManager
func GetFunctionManager() FunctionManager {
	return FunctionManager{}
}

func (functionManager FunctionManager) GetFunction(funcName string) RemoteFunctionHolder {
	return RemoteFunctionHolder{}
}

// Get a function name
func (functionManager *FunctionManager) GetFunctionName(function FuncType1) string {
	return runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
}

// Register remote function to FunctionManager
func (functionManager *FunctionManager) RegisterRemoteFunction(funHolder RemoteFunctionHolder) {
	if _, found := functionManager.funcNameMap[funHolder.funcName]; !found {
		functionManager.funcNameMap[funHolder.funcName] = funHolder.function
	}
}

// Serialize user function
func (functionManager *FunctionManager) GetRayFunction(fn func(interface{}) interface{}) RayFunction {
	functionDescriptor := functionManager.GetFunctionDescriptor(fn)
	return RayFunction{functionDescriptor}
}

// Get GoFunctionDescriptor
func (functionManager *FunctionManager) GetFunctionDescriptor(fn func(interface{}) interface{}) GoFunctionDescriptor {
	wrapper := FunctionWrapper{fn}
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	e.Encode(wrapper)
	funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	functionDescriptor := GoFunctionDescriptor{funcName, b.Bytes()}
	return functionDescriptor
}

type FuncType1 func(...interface{}) interface{}

type RemoteFunctionHolder struct {
	function FuncType1
	funcName string
}

func BuildRemoteFunctionHolder(function FuncType1) RemoteFunctionHolder {
	funcManager := GetFunctionManager()
	funcName := funcManager.GetFunctionName(function)
	holder := RemoteFunctionHolder{function, funcName}
	funcManager.RegisterRemoteFunction(holder)
	return holder
}

var (
	funcManager *FunctionManager
)

func GetFunctionManagerInstance() *FunctionManager {
	if funcManager == nil {
		funcNameMap := map[string]FuncType1{}
		funcManager = &FunctionManager{funcNameMap}
	}
	return funcManager
}
