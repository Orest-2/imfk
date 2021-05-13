package mfsController

import (
	"math"

	"github.com/Orest-2/imfk/lib"
	"github.com/Orest-2/imfk/lib/mfs"
	"github.com/Orest-2/imfk/models"
)

func get2DPlotData(
	mf models.MembershipFunc, params, inData, plotParams []float64,
) (
	[]float64, []float64, []float64, error,
) {

	vx := []float64{}
	vy := []float64{}
	mfParams := []float64{}

	getVectorX := func(p []float64, i int) []float64 {

		if inData != nil {
			return inData
		}

		x := []float64{}
		fromX := 0.0
		toX := 100.0
		step := 1.0

		if len(plotParams) >= 2 {

			lib.UnpackFloat64(plotParams, &toX, &step)
			if step <= 0 {
				step = toX / 100
			}

		} else {

			fp := p[i]
			fpabs := math.Abs(fp)
			lp := p[len(p)-1]

			fromX = fp - fpabs
			toX = lp + fpabs

		}

		for i := fromX; i <= toX/step; i++ {
			x = append(x, i*step)
		}

		return x

	}

	switch mf.Code {
	case mfs.TriangularMembershipFuncCode:

		tmf, err := mfs.NewTriangular(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.TrapezoidalMembershipFuncCode:

		tmf, err := mfs.NewTrapezoidal(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.SquareZSplineMembershipFuncCode:

		tmf, err := mfs.NewSquareZSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.HarmoniousZSplineMembershipFuncCode:

		tmf, err := mfs.NewHarmoniousZSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case
		mfs.SSigmoidalMembershipFuncCode,
		mfs.ZSigmoidalMembershipFuncCode:

		tmf, err := mfs.NewSigmoidal(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 1)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.ZLinearMembershipFuncCode:

		tmf, err := mfs.NewZLinear(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.SquareSSplineMembershipFuncCode:

		tmf, err := mfs.NewSquareSSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.HarmoniousSSplineMembershipFuncCode:

		tmf, err := mfs.NewHarmoniousSSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.SLinearMembershipFuncCode:

		tmf, err := mfs.NewSLinear(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.BellShapedMembershipFuncCode:

		tmf, err := mfs.NewBellShaped(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 2)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.GaussMembershipFuncCode:

		tmf, err := mfs.NewGauss(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 1)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	}

	return vx, vy, mfParams, nil

}

func get3DPlotData(
	mf models.MembershipFunc, funcParams, inData [][]float64, plotParams []float64,
) (
	[]float64, []float64, []float64, [][]float64, [][]float64, error,
) {
	vx := []float64{}
	vy := []float64{}
	vz := []float64{}
	mz := [][]float64{}
	mfParams := [][]float64{}

	p := 0.1
	m := 10.0
	s := 1

	if len(plotParams) >= 2 {
		lib.UnpackFloat64(plotParams, &m, &p)

		if p <= 0 {
			p = m / 100
		}

		s = int(m / p)

		for i := 0; i <= s; i++ {

			k := float64(i)
			vx = append(vx, k*p)
			vy = append(vy, k*p)

		}
	}

	getInData := func() [][]float64 {

		if inData != nil {
			return inData
		}

		res := [][]float64{}

		x := 0
		y := 0

		for i := 0; i <= s*s; i++ {
			x = i / s

			if i%s == 0 {
				y = 0
			}

			res = append(res, []float64{float64(y) * p, float64(x) * p})

			y++
		}

		return res

	}

	switch mf.Code {
	case mfs.ConeMembershipFuncCode:

		tmf, err := mfs.NewCone(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.PyramidalMembershipFuncCode:

		tmf, err := mfs.NewPyramidal(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.TrapezoidalPyramidalMembershipFuncCode:

		tmf, err := mfs.NewTrapezoidalPyramidal(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.GeneralizedSigmoidMembershipFuncCode:

		tmf, err := mfs.NewGeneralizedSigmoid(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.Bell3DShapedMembershipFuncCode:

		tmf, err := mfs.NewBellShaped3D(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.Gauss3DMembershipFuncCode:

		tmf, err := mfs.NewGauss3D(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.HyperboloidMembershipFuncCode:

		tmf, err := mfs.NewHyperboloid(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.Ellipsoid3DMembershipFuncCode:

		tmf, err := mfs.NewEllipsoid(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	}

	return vx, vy, vz, mz, mfParams, nil
}

const (
	IntersectionOperation          = "intersection"
	AssociationOperation           = "association"
	DifferenceOperation            = "difference"
	SymmetricalDifferenceOperation = "symmetrical_difference"
	DisjunctiveSumOperation        = "disjunctive_sum"
)

func applyOperationSlice(operation string, in, data []float64) []float64 {

	res := []float64{}

	switch operation {

	case IntersectionOperation:

		for i, v := range in {
			v2 := data[i]
			res = append(res, math.Min(v, v2))
		}

	case AssociationOperation:

		for i, v := range in {
			v2 := data[i]
			res = append(res, math.Max(v, v2))
		}

	case DifferenceOperation:

		for i, v := range in {
			v2 := data[i]
			res = append(res, math.Min(v, 1-v2))
		}

	case SymmetricalDifferenceOperation:

		for i, v := range in {
			v2 := data[i]
			res = append(res, math.Abs(v-v2))
		}

	case DisjunctiveSumOperation:

		for i, v := range in {
			v2 := data[i]
			res = append(res, math.Max(math.Min(v, 1-v2), math.Min(1-v, v2)))
		}

	}

	return res

}
