package controllers

import (
	"net/http"

	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
)

type getSettingsRosponse struct {
	MembershipFuncs []models.MembershipFunc `json:"membership_funcs"`
}

// GetSettings ...
func GetSettings(c *gin.Context) {
	mf := []models.MembershipFunc{}

	models.DB.Find(&mf)

	res := getSettingsRosponse{
		MembershipFuncs: mf,
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": res,
		},
	)
}
