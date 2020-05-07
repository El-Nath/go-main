package usecase

import (
	"database/sql"

	"../repository"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context, db *sql.DB) {
	q := `SELECT * FROM public.user ORDER BY id ASC`
	v := []interface{}{}
	d, err := repository.SelectSql(db, q, v)
	if err != nil {
		repository.BadResponse(c, "Error while getting data")
		return
	}
	repository.SuccessGettingData(c, d)
}

func InsertUser(c *gin.Context, db *sql.DB) {
	nN := c.PostForm("name")
	if nN == "" {
		repository.BadResponse(c, "no new data")
		return
	}

	q := "INSERT INTO public.user (name) VALUES ($1) RETURNING id"
	v := []interface{}{}
	v = append(v, nN)

	d, err := repository.InsertSqlReturnData(db, q, v)
	if err != nil {
		repository.BadResponse(c, "Error while inserting data")
		return
	}

	d.Name = nN

	repository.SuccessResponse(c, d)
}
