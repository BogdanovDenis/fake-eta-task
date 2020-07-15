package predict

import "testing"

func TestPredictServiceImplement(t *testing.T) {
	// just check implementing interface at build time.
	var ps PredictService = &PredictServiceImpl{}
	_ = ps
}
