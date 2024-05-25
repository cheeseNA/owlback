// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteTaskByIDParams is parameters of delete-task-by-id operation.
type DeleteTaskByIDParams struct {
	// Task ID.
	TaskId uuid.UUID
}

func unpackDeleteTaskByIDParams(packed middleware.Parameters) (params DeleteTaskByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "taskId",
			In:   "path",
		}
		params.TaskId = packed[key].(uuid.UUID)
	}
	return params
}

func decodeDeleteTaskByIDParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteTaskByIDParams, _ error) {
	// Decode path: taskId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "taskId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.TaskId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "taskId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetTaskByIDParams is parameters of get-task-by-id operation.
type GetTaskByIDParams struct {
	// Task ID.
	TaskId uuid.UUID
}

func unpackGetTaskByIDParams(packed middleware.Parameters) (params GetTaskByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "taskId",
			In:   "path",
		}
		params.TaskId = packed[key].(uuid.UUID)
	}
	return params
}

func decodeGetTaskByIDParams(args [1]string, argsEscaped bool, r *http.Request) (params GetTaskByIDParams, _ error) {
	// Decode path: taskId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "taskId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.TaskId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "taskId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
