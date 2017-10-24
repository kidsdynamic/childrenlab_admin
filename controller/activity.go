package controller

import (
	"net/http"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ActivityRawData struct {
	ID             int64     `db:"id",json:"id"`
	Indoor         string    `db:"indoor",json:"indoor"`
	Outdoor        string    `db:"outdoor",json:"outdoor"`
	Time           int64     `db:"time",json:"time"`
	MacID          string    `db:"mac_id",json:"macId"`
	TimeZoneOffset int64     `db:"time_zone_offset",json:"timeZoneOffset"`
	DateCreated    time.Time `db:"date_created",json:"dateCreated"`
}
type Activity struct {
	ID           int64     `db:"id",json:"id"`
	MacID        string    `db:"mac_id",json:"macId"`
	Type         string    `db:"type",json:"type"`
	Steps        int64     `db:"steps",json:"steps"`
	ReceivedDate time.Time `db:"received_date",json:"receivedDate"`
}

func GetActivityRawForAdmin(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	macID := c.Param("macID")

	var activityRaw []ActivityRawData
	if err := db.Select(&activityRaw, "select id, indoor, outdoor, time, mac_id, time_zone_offset, date_created from activity_raw where mac_id = ? order by date_created desc limit 100", macID); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on getting activities from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}
	var totalCount int
	if err := db.Get(&totalCount, "select count(*) from activity_raw where mac_id = ?", macID); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on getting activities count from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       activityRaw,
		"totalCount": totalCount,
	})
}

func GetActivityListForAdmin(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	macID := c.Param("macID")
	max, page, err := GetMaxAndPage(*c)
	if err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on parse string to int"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
	}

	var activity []Activity
	if err := db.Select(&activity, "select id, mac_id, type, steps, received_date from activity where mac_id = ? order by received_date desc limit ?,?", macID, (max*page)-max, max); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on getting activities from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	var totalCount int
	if err := db.Get(&totalCount, "select count(*) from activity where mac_id = ?", macID); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on getting activities count from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       activity,
		"totalCount": totalCount,
	})
}
