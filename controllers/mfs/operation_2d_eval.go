package mfsController

import (
	"errors"
	"net/http"

	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type operation2DEvalResponseData struct {
	eval2DRequest
	Result []float64 `json:"result"`
}

type operation2DEvalRequest struct {
	Data []eval2DRequest `json:"data" binding:"required"`
}

type operation2DEvalResponse struct {
	Data []operation2DEvalResponseData `json:"data"`
}

// Operation2DEval ...
func Operation2DEval(c *gin.Context) {
	json := operation2DEvalRequest{}

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

	res := operation2DEvalResponse{}

	for _, data := range json.Data {

		mf := models.MembershipFunc{}

		err := models.DB.First(&mf, data.MembershipFuncID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resDataItem := operation2DEvalResponseData{
			eval2DRequest: eval2DRequest{
				MembershipFuncID: data.MembershipFuncID,
			},
		}

		resDataItem.InData,
			resDataItem.Result,
			resDataItem.FuncParams,
			err =
			get2DPlotData(
				mf,
				data.FuncParams,
				data.InData,
				nil,
			)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res.Data = append(res.Data, resDataItem)

	}

	resDataItem := operation2DEvalResponseData{
		eval2DRequest: eval2DRequest{
			MembershipFuncID: 0,
			FuncParams:       []float64{},
			InData:           res.Data[0].InData,
		},
		Result: res.Data[0].Result,
	}

	for i := 1; i < len(res.Data); i++ {

		resDataItem.Result = applyOperationSlice(operation, resDataItem.Result, res.Data[i].Result)

	}

	res.Data = append(res.Data, resDataItem)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": res.Data,
		},
	)

}
