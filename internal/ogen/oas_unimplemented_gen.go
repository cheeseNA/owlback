// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CrateTask implements crate-task operation.
//
// Create Task.
//
// POST /tasks
func (UnimplementedHandler) CrateTask(ctx context.Context, req OptTaskRequest) (r CrateTaskRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteTaskByID implements delete-task-by-id operation.
//
// Delete Task by ID.
//
// DELETE /tasks/{taskId}
func (UnimplementedHandler) DeleteTaskByID(ctx context.Context, params DeleteTaskByIDParams) (r DeleteTaskByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTaskByID implements get-task-by-id operation.
//
// Get Task by ID.
//
// GET /tasks/{taskId}
func (UnimplementedHandler) GetTaskByID(ctx context.Context, params GetTaskByIDParams) (r GetTaskByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTasks implements get-tasks operation.
//
// Get Tasks.
//
// GET /tasks
func (UnimplementedHandler) GetTasks(ctx context.Context) (r GetTasksRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTasksOfUser implements get-tasks-of-user operation.
//
// Get public tasks of {userId}.
// If {userId} is me it will return tasks of logged in user.
// If {userId} is the same as logged in user, also return private tasks.
//
// GET /users/{userId}/tasks
func (UnimplementedHandler) GetTasksOfUser(ctx context.Context, params GetTasksOfUserParams) (r GetTasksOfUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// Healthz implements healthz operation.
//
// Get health state.
//
// GET /healthz
func (UnimplementedHandler) Healthz(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// PostCronWrpouiqjflsadkmxcvz780923 implements post-cron-wrpouiqjflsadkmxcvz780923 operation.
//
// Execute crawl.
//
// POST /cron-wrpouiqjflsadkmxcvz780923
func (UnimplementedHandler) PostCronWrpouiqjflsadkmxcvz780923(ctx context.Context) error {
	return ht.ErrNotImplemented
}
