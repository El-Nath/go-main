package usecase

import (
	"strconv"

	"../repository"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func ResultSearch(c *gin.Context, eC *elastic.Client) {
	iN := c.Params.ByName("index")
	if iN == "" {
		repository.BadResponse(c, "No Index to be searched")
		return
	}

	q := c.Query("query")
	if q == "" {
		repository.BadResponse(c, "No query")
		return
	}
	s := 0
	t := 10
	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		s = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		t = i
	}

	sT, err := repository.SearchResult(iN, s, t, q, c, eC)
	if err != nil {
		repository.BadResponse(c, err.Error())
		return
	}

	rS, err := repository.GenerateData(sT)
	if err != nil {
		repository.BadResponse(c, err.Error())
	}

	repository.SearchSuccess(c, rS)
}
