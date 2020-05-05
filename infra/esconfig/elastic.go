package esconfig

import (
	"log"

	"github.com/olivere/elastic"
)

type Env struct {
	ES *elastic.Client
}

func ConnectElastic(u string, s bool) (*elastic.Client, error) {
	elasticClient, err := elastic.NewClient(
		elastic.SetURL(u),
		elastic.SetSniff(s),
	)

	if err != nil {
		log.Println(err)
	}

	return elasticClient, nil
}
