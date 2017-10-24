package controller

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type BatteryStatus struct {
	MacID        string `db:"mac_id",json:"macId"`
	BatteryLife  int64  `db:"battery_life",json:"batteryLife"`
	DateReceived int64  `db:"date_received",json:"dateReceived"`
}

func GetBatteryStatus(c *gin.Context) {
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

	var batteryStatus []BatteryStatus
	if err := db.Select(&batteryStatus, "select * from battery_status where mac_id = ? order by date_received desc limit ?,?", macID, (max*page)-max, max); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on retrieve Battery Status"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieve Battery Status from database",
			"error":   err,
		})
		return
	}

	var totalCount int
	if err := db.Get(&totalCount, "select count(*) from battery_status where mac_id = ?", macID); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on getting activities count from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       batteryStatus,
		"totalCount": totalCount,
	})
}
