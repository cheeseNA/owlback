// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

func decodeCrateTaskResponse(resp *http.Response) (res *CrateTaskCreated, _ error) {
	switch resp.StatusCode {
	case 201:
		// Code 201.
		return &CrateTaskCreated{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeDeleteTaskByIDResponse(resp *http.Response) (res *DeleteTaskByIDOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &DeleteTaskByIDOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetTaskByIDResponse(resp *http.Response) (res GetTaskByIDRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response TaskResponse
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		// Code 404.
		return &GetTaskByIDNotFound{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetTasksResponse(resp *http.Response) (res []TaskResponse, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response []TaskResponse
			if err := func() error {
				response = make([]TaskResponse, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem TaskResponse
					if err := elem.Decode(d); err != nil {
						return err
					}
					response = append(response, elem)
					return nil
				}); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if response == nil {
					return errors.New("nil is invalid value")
				}
				var failures []validate.FieldError
				for i, elem := range response {
					if err := func() error {
						if err := elem.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeHealthzResponse(resp *http.Response) (res *HealthzOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &HealthzOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
