// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "cron-wrpouiqjflsadkmxcvz780923"
				origElem := elem
				if l := len("cron-wrpouiqjflsadkmxcvz780923"); len(elem) >= l && elem[0:l] == "cron-wrpouiqjflsadkmxcvz780923" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handlePostCronWrpouiqjflsadkmxcvz780923Request([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

				elem = origElem
			case 'h': // Prefix: "healthz"
				origElem := elem
				if l := len("healthz"); len(elem) >= l && elem[0:l] == "healthz" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHealthzRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 't': // Prefix: "tasks"
				origElem := elem
				if l := len("tasks"); len(elem) >= l && elem[0:l] == "tasks" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleGetTasksRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCrateTaskRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "taskId"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "DELETE":
							s.handleDeleteTaskByIDRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleGetTaskByIDRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'u': // Prefix: "users/"
				origElem := elem
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "userId"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/tasks"
					origElem := elem
					if l := len("/tasks"); len(elem) >= l && elem[0:l] == "/tasks" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetTasksOfUserRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "cron-wrpouiqjflsadkmxcvz780923"
				origElem := elem
				if l := len("cron-wrpouiqjflsadkmxcvz780923"); len(elem) >= l && elem[0:l] == "cron-wrpouiqjflsadkmxcvz780923" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = "PostCronWrpouiqjflsadkmxcvz780923"
						r.summary = "Execute crawl"
						r.operationID = "post-cron-wrpouiqjflsadkmxcvz780923"
						r.pathPattern = "/cron-wrpouiqjflsadkmxcvz780923"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'h': // Prefix: "healthz"
				origElem := elem
				if l := len("healthz"); len(elem) >= l && elem[0:l] == "healthz" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = "Healthz"
						r.summary = "healthz"
						r.operationID = "healthz"
						r.pathPattern = "/healthz"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 't': // Prefix: "tasks"
				origElem := elem
				if l := len("tasks"); len(elem) >= l && elem[0:l] == "tasks" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "GetTasks"
						r.summary = "Get Tasks"
						r.operationID = "get-tasks"
						r.pathPattern = "/tasks"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CrateTask"
						r.summary = "Create Task"
						r.operationID = "crate-task"
						r.pathPattern = "/tasks"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "taskId"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "DELETE":
							r.name = "DeleteTaskByID"
							r.summary = "Delete Task by ID"
							r.operationID = "delete-task-by-id"
							r.pathPattern = "/tasks/{taskId}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "GetTaskByID"
							r.summary = "Get Task by ID"
							r.operationID = "get-task-by-id"
							r.pathPattern = "/tasks/{taskId}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 'u': // Prefix: "users/"
				origElem := elem
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "userId"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/tasks"
					origElem := elem
					if l := len("/tasks"); len(elem) >= l && elem[0:l] == "/tasks" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "GetTasksOfUser"
							r.summary = "Get tasks of user"
							r.operationID = "get-tasks-of-user"
							r.pathPattern = "/users/{userId}/tasks"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
