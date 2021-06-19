package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-basic-exam/go_exam_4/internal"
	"go-basic-exam/go_exam_4/internal/employee"
	"go-basic-exam/go_exam_4/internal/health"
	"net/http"
)

type CloseFunc func() error

type route struct {
	Group          string
	Path           string
	HttpMethod     string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

func NewRoutes(e *echo.Echo, cv *internal.Configs) ([]CloseFunc, error) {
	closeFuncs := make([]CloseFunc, 0)

	emp := employee.NewEndpoint(cv)
	routes := []route{
		{
			Group:          "",
			Path:           "/health_check",
			HttpMethod:     http.MethodGet,
			HandlerFunc:    health.HealthCheck,
			MiddlewareFunc: nil,
		},
		{
			Group:          "employee",
			Path:           "/byId",
			HttpMethod:     http.MethodPost,
			HandlerFunc:    emp.GetEmployeeById,
			MiddlewareFunc: nil,
		},
	}

	// http connection
	for _, rt := range routes {
		mw := []echo.MiddlewareFunc{
			middleware.BodyDumpWithConfig(BodyDumpConfig()),
		}
		mw = append(mw, rt.MiddlewareFunc...)
		e.Group(rt.Group).Add(rt.HttpMethod, rt.Path, rt.HandlerFunc, mw...)
	}

	return closeFuncs, nil
}
