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
func (UnimplementedHandler) CrateTask(ctx context.Context, req OptTaskRequest) error {
	return ht.ErrNotImplemented
}

// DeleteTaskByID implements delete-task-by-id operation.
//
// Delete Task by ID.
//
// DELETE /tasks/{taskId}
func (UnimplementedHandler) DeleteTaskByID(ctx context.Context, params DeleteTaskByIDParams) error {
	return ht.ErrNotImplemented
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
func (UnimplementedHandler) GetTasks(ctx context.Context) (r []TaskResponse, _ error) {
	return r, ht.ErrNotImplemented
}