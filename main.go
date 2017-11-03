package main

import (
	"net/http"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kidsdynamic/childrenlab_admin/config"
	"github.com/kidsdynamic/childrenlab_admin/controller"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Childrenlab Admin"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "Debug",
			Value:  "false",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_USER",
			Name:   "database_user",
			Usage:  "Database user name",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_PASSWORD",
			Name:   "database_password",
			Usage:  "Database password",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_IP",
			Name:   "database_IP",
			Usage:  "Database IP address with port number",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_NAME",
			Name:   "database_name",
			Usage:  "Database name",
			Value:  "swing_test_record",
		},
		cli.StringFlag{
			EnvVar: "SUPER_ADMIN_TOKEN",
			Name:   "super_admin_token",
			Value:  "1",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "BASE_URL",
			Name:   "base_url",
			Value:  "http://localhost:8110",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "S3_BASE_URL",
			Name:   "s3_base_url",
			Value:  "https://childrenlabqa.s3.amazonaws.com",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "PORT",
			Name:   "port",
			Value:  "8114",
			Usage:  "",
		},
	}

	app.Action = func(c *cli.Context) error {
		controller.SetupDatabase(controller.Database{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		})

		config.SetConfig(c.String("s3_base_url"), c.String("base_url"))

		controller.SuperAdminToken = c.String("super_admin_token")

		r := gin.New()
		if c.Bool("debug") {
			r.Use(gin.Logger())
		}

		r.Use(gin.Recovery())

		r.Static("/dist", "view/dist")

		r.LoadHTMLGlob("view/dist/*.html")

		//Super Admin API
		superAdminAuthAPI := r.Group("/superAdmin/api")
		superAdminAuthAPI.Use(controller.SuperAdminAuth)
		superAdminAuthAPI.POST("/createAdminUser", controller.CreateAdminUser)

		//Admin API
		r.POST("/admin/api/login", controller.Login)
		adminAuthAPI := r.Group("/admin/api")
		adminAuthAPI.Use(controller.AdminAuth)
		adminAuthAPI.GET("/checkAuth", controller.CheckAuth)
		adminAuthAPI.GET("/dashboard", controller.Dashboard)
		adminAuthAPI.GET("/user", controller.GetAllUser)
		adminAuthAPI.GET("/kid", controller.GetAllKidList)
		adminAuthAPI.GET("/kid/rawActivity/:macID", controller.GetActivityRawForAdmin)
		adminAuthAPI.GET("/kid/activity/:macID", controller.GetActivityListForAdmin)
		adminAuthAPI.GET("/kid/battery/:macID", controller.GetBatteryStatus)
		adminAuthAPI.DELETE("/kid/deleteMacID", controller.DeleteMacID)
		adminAuthAPI.GET("/fw", controller.GetFwFileList)
		adminAuthAPI.PUT("/fw/updateActivation", controller.UpdateActivation)

		//Page
		r.GET("/", indexPage)
		r.GET("/login", indexPage)
		r.GET("/dashboard", indexPage)
		r.GET("/user", indexPage)
		r.GET("/kid", indexPage)
		r.GET("/kid/rawActivity/:macID", indexPage)
		r.GET("/kid/activity/:macID", indexPage)
		r.GET("/fw", indexPage)

		return r.Run(fmt.Sprintf(":%s", c.String("port")))

	}
	app.Run(os.Args)

}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"S3BaseURL": config.ConfigSettings.S3BasedURL,
		"BaseURL":   config.ConfigSettings.BaseURL,
	})
}
