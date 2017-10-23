package controller

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_admin/model"
	"github.com/pkg/errors"
)

func Dashboard(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var dashboard model.Dashboard

	var signupCounts []model.SignupCountByDate
	if err := db.Select(&signupCounts, "select count(*) as signup, DATE_FORMAT(DATE(date_created), '%Y/%m/%d') as date from user u JOIN role r ON u.role_id = r.id where date_created != 0000-00-00"+
		" and r.`authority` = 'ROLE_USER' AND date_created >= '2017-04-24' group by date order by date desc LIMIT 14"); err != nil {

		fmt.Printf("%+v", errors.Wrap(err, "Error on retrieve signup dashboard from Admin"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dashboard.Signup = signupCounts

	var activityCount []model.ActivityCountByDate
	if err := db.Select(&activityCount, "select count(*) as count, DATE_FORMAT(DATE(date_created), '%Y/%m/%d') as date, count(DISTINCT(user_id)) as userCount, COALESCE(sum(indoor_steps), 0) as indoorSteps"+
		", COALESCE(sum(outdoor_steps), 0) as outdoorSteps from activity_raw group by date order by date desc LIMIT 14"); err != nil {

		fmt.Printf("%+v", errors.Wrap(err, "Error on retrieve activityCountByDate dashboard from Admin"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	dashboard.Activity = activityCount

	if err := db.Get(&dashboard.TotalUserCount, "select count(*) from user u JOIN role r ON u.role_id = r.id WHERE date_created != 0000-00-00 and authority = 'ROLE_USER' AND email not like '%kidsdynamic.com'"); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on retrieve TotalUserCount dashboard from Admin"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := db.Get(&dashboard.TotalActivityCount, "select count(*) from activity_raw"); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on retrieve TotalActivityCount dashboard from Admin"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	/*	// Activity count depend on activity (event) date - https://app.asana.com/0/33043844747220/349867637183239
		var activityCountOnEventDate []model.ActivityCountByDate
		if rows, err := db.Raw("select count(*), DATE_FORMAT(FROM_UNIXTIME(time), '%Y/%m/%d') as date, count(DISTINCT(user_id)) as userCount, sum(indoor_steps), sum(outdoor_steps) " +
			"from activity_raw group by date order by date desc LIMIT 20").Rows(); err != nil {
			logError(errors.Wrap(err, "Error on retrieve activity dashboard from Admin"))
		} else {
			defer rows.Close()
			for rows.Next() {
				var activity model.ActivityCountByDate
				rows.Scan(&activity.ActivityCount, &activity.Date, &activity.UserCount, &activity.IndoorSteps, &activity.OutdoorSteps)
				activityCountOnEventDate = append(activityCountOnEventDate, activity)
			}
		}
		dashboard.ActivityByEventDate = activityCountOnEventDate*/

	c.JSON(http.StatusOK, dashboard)
}
