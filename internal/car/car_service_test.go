package car

import "testing"

func TestCarServiceImplement(t *testing.T) {
	// just check implementing interface at build time.
	var cs CarService = &CarServiceImpl{}
	_ = cs
}
