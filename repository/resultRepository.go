package repository

import (
	"encoding/json"
	"fmt"

	"../domain/model"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func SearchResult(iN string, s int, t int, q string, c *gin.Context, eC *elastic.Client) (*elastic.SearchResult, error) {

	esQuery := elastic.NewMultiMatchQuery(q, "title", "content").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := eC.Search().
		Index(iN).
		Query(esQuery).
		From(s).Size(t).
		Do(c.Request.Context())
	if err != nil {
		return result, err
	}

	return result, nil
}

func GenerateData(r *elastic.SearchResult) (model.SearchResponse, error) {
	res := model.SearchResponse{}
	docs := make([]model.DataResponse, 0)
	for _, hit := range r.Hits.Hits {
		var d model.DataResponse
		fmt.Println(hit)
		if err := json.Unmarshal(hit.Source, &d); err != nil {
			continue
		}
		docs = append(docs, d)
	}

	res.Time = fmt.Sprintf("%d", r.TookInMillis)
	res.Documents = docs
	return res, nil
}
