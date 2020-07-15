package wait

import (
	"context"
	"math"
)

//go:generate mockgen -destination=../mocks/mock_wait_service.go -package=mocks github.com/fake-eta-task/wait WaitService

type WaitService interface {
	CalculateBestWaitingTime(ctx context.Context, waitingTimes []int64) (int64, error)
}

type WaitServiceImpl struct{}

func NewWaitService() *WaitServiceImpl {
	return &WaitServiceImpl{}
}

func (s *WaitServiceImpl) CalculateBestWaitingTime(ctx context.Context, waitingTimes []int64) (int64, error) {
	minTime := int64(math.MaxInt64)
	for _, time := range waitingTimes {
		if time < minTime {
			minTime = time
		}
	}
	return minTime, nil
}
