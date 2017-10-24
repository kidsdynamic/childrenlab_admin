package controller

import (
	"net/http"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Kid struct {
	ID              int64     `db:"id",json:"id"`
	Name            string    `db:"name",json:"name"`
	DateCreated     time.Time `db:"date_created",json:"dateCreated"`
	MacID           string    `db:"mac_id",json:"macId"`
	FirmwareVersion *string   `db:"firmware_version",json:"firmwareVersion"`
	Profile         *string   `db:"profile",json:"profile"`
	ParentEmail     string    `db:"parent_email",json:"parent,omitempty"`
}

func GetAllKidList(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	max, page, err := GetMaxAndPage(*c)
	if err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on parse string to int"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
	}

	var kids []Kid
	if err := db.Select(&kids, "select k.id, k.name, k.date_created, k.mac_id, k.firmware_version, k.profile, u.email as parent_email from"+
		" kids k left join user u on k.parent_id=u.id order by date_created desc limit ?, ?", (max*page)-max, max); err != nil {

		fmt.Printf("%+v", errors.Wrapf(err, "Error on retriving kid list: %#v", kids))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retriving kid list",
			"error":   err,
		})
		return
	}

	var totalKidCount int
	if err := db.Get(&totalKidCount, "select count(*) from kids"); err != nil {
		fmt.Printf("%+v", errors.Wrapf(err, "Error on retriving kid list: %#v", kids))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retriving kid list",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":          kids,
		"totalKidCount": totalKidCount,
		"page":          page,
	})
}
