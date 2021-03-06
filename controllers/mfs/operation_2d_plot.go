package mfsController

import (
	"errors"
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

		// resDataItem.X = applyOperationSlice(operation, resDataItem.X, res.Data[i].X)
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
