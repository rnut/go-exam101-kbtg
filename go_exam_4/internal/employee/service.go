package employee

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"go-basic-exam/go_exam_4/internal"
	"go-basic-exam/go_exam_4/internal/models"
)

type employeeRepo interface {
	GetEmployeeById(ctx context.Context, empId string) ([]models.Employee, error)
}

type Service struct {
	cv   *internal.Configs
	repo employeeRepo
}

func NewService(cv *internal.Configs) *Service {
	return &Service{
		cv:   cv,
		repo: nil,
	}
}

func (s Service) GetEmployeeById(c context.Context, employeeId string) ([]models.Employee, error) {

	if len(employeeId) != 10 {
		return []models.Employee{}, errors.New("employee id length require 10 digits")
	}

	log.Infof("service query: %#v", employeeId)

	return s.repo.GetEmployeeById(c, employeeId)
}
