package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// HarmoniousSSpline ...
type HarmoniousSSpline struct {
	a, b float64
}

// NewHarmoniousSSpline ...
func NewHarmoniousSSpline(params []float64) (*HarmoniousSSpline, error) {

	t := &HarmoniousSSpline{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t HarmoniousSSpline) GetParams() []float64 {

	return []float64{t.a, t.b}

}

// SetParams ...
func (t HarmoniousSSpline) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t HarmoniousSSpline) Eval(x float64) float64 {

	a, b := t.a, t.b

	if x < a {
		return 0
	}

	if a <= x && x <= b {
		c := 0.5
		f := x - b
		s := b - a

		return c + c*math.Cos((f/s)*math.Pi)
	}

	return 1

}

// validateAndSetParams ...
func (t *HarmoniousSSpline) validateAndSetParams(params []float64) error {

	var a, b float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64(params, &a, &b)

	if !(a < b) {
		return fmt.Errorf("має виконуватися співвідношенням a < b: (a=%v, b=%v)", a, b)
	}

	t.a, t.b = a, b

	return nil

}
