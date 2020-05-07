package repository

import (
	"net/http"

	"../domain/model"
	"github.com/gin-gonic/gin"
)

var (
	rN model.ResponseNoData
	sS model.ResponseSearchData
	sN model.ResponseInsertData
)

func BadResponse(c *gin.Context, m string) {
	rN.Status = http.StatusBadRequest
	rN.Message = m

	c.JSON(http.StatusBadRequest, rN)
}

func SuccessResponse(c *gin.Context, d model.User) {
	sN.Status = http.StatusOK
	sN.Message = "success inserting data"
	sN.Data = d

	c.JSON(http.StatusOK, sN)
}

func SuccessGettingData(c *gin.Context, d []model.User) {
	sS.Status = http.StatusOK
	sS.Message = "success getting data"
	sS.Data = d

	c.JSON(http.StatusOK, sS)
}
