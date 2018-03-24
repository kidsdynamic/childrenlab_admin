package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type MacIDNoKid struct {
	MacID         string    `json:"mac_id" db:"mac_id"`
	LastUpload    time.Time `json:"last_upload" db:"last_updated"`
	ActivityCount int       `json:"activity_count" db:"activity_count"`
}

func ErrorData(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var macIdNoKid []MacIDNoKid
	if err := db.Select(&macIdNoKid, "select mac_id, max(last_updated) as last_updated, count(*) as activity_count from activity where (mac_id not in (select mac_id from kids)) group by mac_id"); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on get mac id no kid list"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieve mac id no kid list from database",
			"error":   err,
		})
		return
	}

	fmt.Println(len(macIdNoKid))

	c.JSON(http.StatusOK, macIdNoKid)

}
