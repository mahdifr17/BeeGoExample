package main

import (
	"strconv"

	"github.com/mahdifr17/BeeGoExample/modules/db"
	_ "github.com/mahdifr17/BeeGoExample/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Have to import driver "github.com/lib/pq", imported on db.go
	dbDriver := beego.AppConfig.String("dbdriver")
	dbString := beego.AppConfig.String("dbstring")
	// If true => Force drop table and re-create. Value at app.conf
	dbForce, _ := strconv.ParseBool(beego.AppConfig.String("dbForce"))
	// If true => Print log. Value at app.conf
	dbVerbose, _ := strconv.ParseBool(beego.AppConfig.String("dbVerbose"))
	db.StartDB("default", dbDriver, dbString, dbForce, dbVerbose)

	beego.Run()
}
