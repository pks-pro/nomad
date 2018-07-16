package interfaces

import (
	"context"

	"github.com/hashicorp/nomad/client/driver/env"
	"github.com/hashicorp/nomad/nomad/structs"
)

/*
     pre-run   post-run      pre-stop               post-stop
        |        |              |                     |
        |        |              |                     |
 --------> run ------> exited ----------> not restart ---------> garbage collect

*/

// TaskHook is a lifecycle hook into the life cycle of a task runner.
type TaskHook interface {
	Name() string
}

type TaskPrerunRequest struct {
	// HookData is previously set data by the hook
	HookData map[string]string

	// Task is the task to run
	Task *structs.Task

	// Vault token may optionally be set if a Vault token is available
	VaultToken string

	// TaskDir is the task's directory on the host
	TaskDir string

	// TaskEnv is the task's environment
	TaskEnv *env.TaskEnv
}

type TaskPrerunResponse struct {
	// Env is the environment variables to set for the task
	Env map[string]string

	// HookData allows the hook to emit data to be passed in the next time it is
	// run
	HookData map[string]string

	// Done lets the hook indicate that it should only be run once
	Done bool
}

type TaskPrerunHook interface {
	TaskHook
	Prerun(context.Context, *TaskPrerunRequest, *TaskPrerunResponse) error
}

// XXX If we want consul style hooks, need to have something that runs after the
// tasks starts
type TaskPostrunRequest struct {
	// Network info
}
type TaskPostrunResponse struct{}

type TaskPostrunHook interface {
	TaskHook
	Postrun() error
	//Postrun(context.Context, *TaskPostrunRequest, *TaskPostrunResponse) error
}

type TaskPoststopRequest struct{}
type TaskPoststopResponse struct{}

type TaskPoststopHook interface {
	TaskHook
	Postrun(context.Context, *TaskPostrunRequest, *TaskPostrunResponse) error
}

type TaskDestroyRequest struct{}
type TaskDestroyResponse struct{}

type TaskDestroyHook interface {
	TaskHook
	Destroy() error
	//Destroy(context.Context, *TaskDestroyRequest, *TaskDestroyResponse) error
}

type TaskUpdateRequest struct {
	VaultToken string
}
type TaskUpdateResponse struct{}

type TaskUpdateHook interface {
	TaskHook
	Update(context.Context, *TaskUpdateRequest, *TaskUpdateResponse) error
}