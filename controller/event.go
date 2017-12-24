package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Event struct {
	ID             int64     `db:"id" json:"id"`
	Email          string    `db:"email" json:"email"`
	Name           string    `db:"name" json:"name"`
	Start          time.Time `db:"start" json:"start"`
	End            time.Time `db:"end" json:"end"`
	Color          string    `db:"color" json:"color"`
	Description    string    `db:"description" json:"description"`
	Alert          int64     `db:"alert" json:"alert"`
	Repeat         string    `db:"repeat" json:"repeat"`
	TimezoneOffset int64     `db:"timezone_offset" json:"timezoneOffset"`
	DateCreated    time.Time `db:"date_created" json:"dateCreated"`
}

func GetAllEvents(c *gin.Context) {
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

	var events []Event
	if err := db.Select(&events, "SELECT e.id, u.email, e.name, start, end, color, description, alert, `repeat`, e.timezone_offset, e.date_created FROM"+
		" event e JOIN user u ON e.user_id = u.id ORDER BY date_created desc limit ?,?", (max*page)-max, max); err != nil {

		fmt.Printf("%+v", errors.Wrap(err, "Error on getting Events"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting events",
			"error":   err,
		})
		return
	}

	var totalCount int
	if err := db.Get(&totalCount, "select count(*) from event"); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on getting event count from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting event",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       events,
		"totalCount": totalCount,
		"page":       page,
	})
}
