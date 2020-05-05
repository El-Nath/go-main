package repository

import (
	"log"
	"time"

	"../domain/model"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"github.com/teris-io/shortid"
)

func BindJsonRequest(c *gin.Context) ([]model.DataRequest, error) {
	var d []model.DataRequest
	if err := c.BindJSON(&d); err != nil {
		return d, err
	}
	return d, nil
}

func CreateNewData(d []model.DataRequest, i string, t string, eC *elastic.Client) *elastic.BulkService {
	bD := eC.
		Bulk().
		Index(i).
		Type(t)
	for _, d := range d {
		doc := model.Data{
			ID:        shortid.MustGenerate(),
			Title:     d.Title,
			CreatedAt: time.Now().UTC(),
			Content:   d.Content,
		}
		bD.Add(elastic.NewBulkIndexRequest().Id(doc.ID).Doc(doc))
	}
	return bD
}

func ValidateData(c *gin.Context, b *elastic.BulkService) error {
	if _, err := b.Do(c.Request.Context()); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
