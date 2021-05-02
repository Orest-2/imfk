package mfsController

import (
	"errors"
	"math"
	"net/http"

	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type operation2DPlotRequest struct {
	Data []make2DPlotRequest `json:"data"`
}

type operation2DPlotResponse struct {
	Data []make2DPlotResponse `json:"data"`
}

const (
	IntersectionOperation          = "intersection"
	AssociationOperation           = "association"
	DifferenceOperation            = "difference"
	SymmetricalDifferenceOperation = "symmetrical_difference"
	DisjunctiveSumOperation        = "disjunctive_sum"
)

// Operation2DPlot ...
func Operation2DPlot(c *gin.Context) {
	json := operation2DPlotRequest{}

	operation := c.Param("operation")

	if operation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "operation is mandatory"})
		return
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(json.Data) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "minimun 2 functions"})
		return
	}

	res := operation2DPlotResponse{}

	for _, data := range json.Data {

		mf := models.MembershipFunc{}

		err := models.DB.First(&mf, data.MembershipFuncID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resDataItem := make2DPlotResponse{
			make2DPlotRequest: make2DPlotRequest{
				MembershipFuncID: data.MembershipFuncID,
				PlotParams:       data.PlotParams,
			},
		}

		resDataItem.X,
			resDataItem.Y,
			resDataItem.make2DPlotRequest.FuncParams,
			err =
			get2DPlotData(
				mf,
				data.FuncParams,
				nil,
				data.PlotParams,
			)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res.Data = append(res.Data, resDataItem)

	}

	resDataItem := make2DPlotResponse{
		make2DPlotRequest: make2DPlotRequest{
			MembershipFuncID: 0,
			FuncParams:       []float64{},
			PlotParams:       []float64{},
		},
		X: res.Data[0].X,
		Y: res.Data[0].Y,
	}

	for i := 1; i < len(res.Data); i++ {

		resDataItem.X = applyOperationSlice(operation, resDataItem.X, res.Data[i].X)
		resDataItem.Y = applyOperationSlice(operation, resDataItem.Y, res.Data[i].Y)

	}

	res.Data = append(res.Data, resDataItem)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": res.Data,
		},
	)

}

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
