package task

import "github.com/ray-project/ray/go/internal/runtime/object"

type TaskSubmitter interface {
	SubmitTask(*RemoteFunctionHolder, []interface{}) object.ObjectRef
}
