package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// BellShaped3D ...
type BellShaped3D struct {
	a, b, c []float64
}

// NewBellShaped3D ...
func NewBellShaped3D(params [][]float64) (*BellShaped3D, error) {

	t := &BellShaped3D{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t BellShaped3D) GetParams() [][]float64 {

	return [][]float64{t.a, t.b, t.c}

}

// SetParams ...
func (t BellShaped3D) SetParams(params [][]float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t BellShaped3D) Eval(xv []float64) float64 {

	a, b, c := t.a, t.b, t.c

	e := 0.0

	for i, x := range xv {

		e1 := x - a[i]

		e += math.Pow(math.Abs(e1/b[i]), 2*c[i])

	}

	return 1 / (1 + e)

}

// validateAndSetParams ...
func (t *BellShaped3D) validateAndSetParams(params [][]float64) error {

	var a, b, c []float64

	mp := 3

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64Slice(params, &a, &b, &c)

	if len(a) < 2 {
		return fmt.Errorf("параметер a має містит два елемента: (a=%v)", a)
	}

	if len(b) < 2 {
		return fmt.Errorf("параметер b має містит два елемента: (b=%v)", b)
	}

	if len(c) < 2 {
		return fmt.Errorf("параметер c має містит два елемента: (c=%v)", b)
	}

	t.a, t.b, t.c = a, b, c

	return nil

}
