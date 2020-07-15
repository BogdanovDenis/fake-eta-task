package internal

import "testing"

func TestInternalServiceImplement(t *testing.T) {
	// just check implementing interface at build time.
	var is InternalService = &InternalServiceImpl{}
	_ = is
}
