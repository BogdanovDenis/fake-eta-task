package predict

import (
	"context"

	"github.com/fake-eta-task/internal/predict/models"
	"github.com/fake-eta-task/internal/predict/operations"
)

//go:generate mockgen -destination=../../mocks/mock_predict_service.go -package=mocks github.com/fake-eta-task/internal/predict PredictService

type PredictService interface {
	PredictWaitingTimes(ctx context.Context, carPositions []models.Position, target models.Position) ([]int64, error)
}

type PredictServiceImpl struct {
	url string
}

func NewPredictServiceImpl(url string) *PredictServiceImpl {
	return &PredictServiceImpl{url: url}
}

func (s *PredictServiceImpl) PredictWaitingTimes(ctx context.Context, carPositions []models.Position, target models.Position) ([]int64, error) {
	cfg := DefaultTransportConfig().WithHost(s.url).WithBasePath("/fake-eta")
	client := NewHTTPClientWithConfig(nil, cfg)
	params := operations.NewPredictParams().
		WithContext(ctx).
		WithPositionList(operations.PredictBody{Source: carPositions, Target: target})

	res, err := client.Operations.Predict(params)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
