package wait

import (
	"testing"

	"context"

	"github.com/fake-eta-task/mocks"
	"github.com/golang/mock/gomock"
)

func TestWaitServiceImplement(t *testing.T) {
	// just check implementing interface at build time.
	var ws WaitService = &WaitServiceImpl{}
	_ = ws
}

func TestWaitServiceCalculateBestWaitingTime(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	wait := mocks.NewMockWaitService(mockCtrl)
	expectBestTime := int64(1)

	bestTime, err := wait.CalculateBestWaitingTime(context.Background(), []int64{1, 2, 3, 4, 5})
	if err != nil {
		t.Fatal(err)
	}

	if bestTime != expectBestTime {
		t.Errorf("different bestTime for calculate best waiting time, expect: %v, got: %v", bestTime, expectBestTime)
	}
}
