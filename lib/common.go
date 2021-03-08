package lib

// UnpackInt ...
func UnpackInt(s []int, vars ...*int) {
	for i := range vars {
		*vars[i] = s[i]
	}
}

// UnpackFloat64 ...
func UnpackFloat64(s []float64, vars ...*float64) {
	for i := range vars {
		*vars[i] = s[i]
	}
}

// UnpackFloat64Slice ...
func UnpackFloat64Slice(s [][]float64, vars ...*[]float64) {
	for i := range vars {
		*vars[i] = s[i]
	}
}

// Vector2Matrix ...
func Vector2Matrix(in []float64, sliceLen int) [][]float64 {

	out := [][]float64{}

	if len(in) < sliceLen*sliceLen {
		return out
	}

	for i := 0; i <= sliceLen; i++ {

		t := in[i*sliceLen : (i+1)*sliceLen]
		out = append(out, t)

	}

	return out

}
