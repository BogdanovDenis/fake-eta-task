package car

import (
	"context"

	carModels "github.com/fake-eta-task/internal/car/models"
	"github.com/fake-eta-task/internal/car/operations"
	"github.com/fake-eta-task/models"
)

//go:generate mockgen -destination=../../mocks/mock_car_service.go -package=mocks github.com/fake-eta-task/internal/car CarService

type CarService interface {
	GetCars(ctx context.Context, userPosition models.Position) ([]carModels.Car, error)
}

type CarServiceImpl struct {
	url string
}

func NewCarServiceImpl(url string) *CarServiceImpl {
	return &CarServiceImpl{url: url}
}

func (s *CarServiceImpl) GetCars(ctx context.Context, userPosition models.Position) ([]carModels.Car, error) {
	cfg := DefaultTransportConfig().WithHost(s.url).WithBasePath("/fake-eta")
	client := NewHTTPClientWithConfig(nil, cfg)
	params := operations.NewGetCarsParams().
		WithContext(ctx).
		WithLat(userPosition.Lat).
		WithLng(userPosition.Lng).
		WithLimit(5)

	res, err := client.Operations.GetCars(params)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
