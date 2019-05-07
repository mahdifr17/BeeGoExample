package controllers

import "github.com/mahdifr17/BeeGoExample/modules/db"

// BasicController holds basic controller with data orm
type BasicController struct {
	BaseController
	Orm         db.ORM
	RequestData []byte
}
