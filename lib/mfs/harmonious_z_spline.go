package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// HarmoniousZSpline ...
type HarmoniousZSpline struct {
	a, b float64
}

// NewHarmoniousZSpline ...
func NewHarmoniousZSpline(params []float64) (*HarmoniousZSpline, error) {

	t := &HarmoniousZSpline{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t HarmoniousZSpline) GetParams() []float64 {

	return []float64{t.a, t.b}

}

// SetParams ...
func (t HarmoniousZSpline) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t HarmoniousZSpline) Eval(x float64) float64 {

	a, b := t.a, t.b

	if x < a {
		return 1
	}

	if a <= x && x <= b {
		c := 0.5
		f := x - a
		s := b - a

		return c + c*math.Cos((f/s)*math.Pi)
	}

	return 0

}

// validateAndSetParams ...
func (t *HarmoniousZSpline) validateAndSetParams(params []float64) error {

	var a, b float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("Мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64(params, &a, &b)

	if !(a < b) {
		return fmt.Errorf("має виконуватися співвідношенням a < b: (a=%v, b=%v)", a, b)
	}

	t.a, t.b = a, b

	return nil

}
