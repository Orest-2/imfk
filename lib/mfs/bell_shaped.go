package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// BellShaped ...
type BellShaped struct {
	a, b, c float64
}

// NewBellShaped ...
func NewBellShaped(params []float64) (*BellShaped, error) {

	t := &BellShaped{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t BellShaped) GetParams() []float64 {

	return []float64{t.a, t.b, t.c}

}

// SetParams ...
func (t BellShaped) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t BellShaped) Eval(x float64) float64 {

	a, b, c := t.a, t.b, t.c

	e1 := math.Abs((x - c) / a)
	e2 := math.Pow(e1, 2*b)

	return 1 / (1 + e2)

}

// validateAndSetParams ...
func (t *BellShaped) validateAndSetParams(params []float64) error {

	var a, b, c float64

	mp := 3

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b, c)", mp)
	}

	lib.UnpackFloat64(params, &a, &b, &c)

	if a <= 0 {
		return fmt.Errorf("має виконуватися співвідношенням a > 0: (a=%v)", a)
	}

	if !(a < c || b < c) {
		return fmt.Errorf("має виконуватися співвідношенням a < c, b < c: (a=%v, b=%v, c=%v)", a, b, c)
	}

	t.a, t.b, t.c = a, b, c

	return nil

}
