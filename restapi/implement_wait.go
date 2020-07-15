package restapi

import (
	"github.com/fake-eta-task/internal"
	carModels "github.com/fake-eta-task/internal/car/models"
	predictModels "github.com/fake-eta-task/internal/predict/models"
	"github.com/fake-eta-task/models"
	"github.com/fake-eta-task/restapi/operations"
	"github.com/fake-eta-task/wait"
	middleware "github.com/go-openapi/runtime/middleware"
)

type RESTHandler struct {
	wait     wait.WaitService
	internal internal.InternalService
}

func NewRESTHandler(wait wait.WaitService, internal internal.InternalService) *RESTHandler {
	return &RESTHandler{
		wait:     wait,
		internal: internal,
	}
}

func (s *RESTHandler) CalculateWaitingTimeImpl(params operations.CalculateWaitingTimeParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	cars, err := s.internal.GetCars(ctx, params.Position)
	if err != nil {
		return operations.NewCalculateWaitingTimeDefault(500).WithPayload(&models.Error{Error: err.Error()})
	}

	if len(cars) == 0 {
		return operations.NewCalculateWaitingTimeDefault(404).WithPayload(&models.Error{Error: "К сожалению, доступных водителей в Вашем районе не найдено"})
	}

	waitingTimes, err := s.internal.PredictWaitingTimes(ctx, getPositions(cars), predictModels.Position(params.Position))
	if err != nil {
		return operations.NewCalculateWaitingTimeDefault(500).WithPayload(&models.Error{Error: err.Error()})
	}

	bestTime, err := s.wait.CalculateBestWaitingTime(ctx, waitingTimes)
	if err != nil {
		return operations.NewCalculateWaitingTimeDefault(500).WithPayload(&models.Error{Error: err.Error()})
	}

	return operations.NewCalculateWaitingTimeOK().WithPayload(bestTime)
}

func getPositions(cars []carModels.Car) (positions []predictModels.Position) {
	for _, car := range cars {
		positions = append(positions, predictModels.Position{Lat: car.Lat, Lng: car.Lng})
	}
	return
}
