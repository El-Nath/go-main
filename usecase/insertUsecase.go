package usecase

import (
	"../repository"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func NewData(c *gin.Context, eC *elastic.Client) {
	iN := c.Params.ByName("fakedata")
	tN := c.Params.ByName("fakejson")

	s, err := repository.BindJsonRequest(c)
	if err != nil {
		repository.BadResponse(c, "Error request data")
		return
	}
	d := repository.CreateNewData(s, iN, tN, eC)

	if err := repository.ValidateData(c, d); err != nil {
		repository.BadResponse(c, "Failed to inserting data")
		return
	}
	repository.SuccessResponse(c, "Success inserting data")
}
