package employee

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-basic-exam/go_exam_4/internal"
	"go-basic-exam/go_exam_4/internal/models"
)

type employeeService interface {
	GetEmployeeById(ctx context.Context, empId string) ([]models.Employee, error)
}

type Endpoint struct {
	cv  *internal.Configs
	srv employeeService //Service Interface Tier
}

func NewEndpoint(cv *internal.Configs) *Endpoint {
	return &Endpoint{cv: cv, srv: NewService(cv)}
}

func (e Endpoint) GetEmployeeById(c echo.Context) error {
	return nil
}
