package internal

import (
	"context"

	"github.com/fake-eta-task/internal/car"
	carModels "github.com/fake-eta-task/internal/car/models"
	"github.com/fake-eta-task/internal/predict"
	predictModels "github.com/fake-eta-task/internal/predict/models"
	"github.com/fake-eta-task/models"
)

//go:generate mockgen -destination=../mocks/mock_intelnal_service.go -package=mocks github.com/fake-eta-task/internal InternalService

type InternalService interface {
	GetCars(ctx context.Context, userPosition models.Position) ([]carModels.Car, error)
	PredictWaitingTimes(ctx context.Context, carPositions []predictModels.Position, target predictModels.Position) ([]int64, error)
}

type InternalServiceImpl struct {
	car.CarService
	predict.PredictService
}

func NewInternalService(car car.CarService, predict predict.PredictService) *InternalServiceImpl {
	return &InternalServiceImpl{
		car,
		predict,
	}
}
