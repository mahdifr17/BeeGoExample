package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	responseStatusSuccess = "success"
	responseStatusError   = "error"
)

// BaseController holds basic response data
type BaseController struct {
	beego.Controller
	status  string
	message string
	data    map[string]interface{}
}

// Prepare runs after Init before request function execution
// Used to initialize data & set content-type header
func (c *BaseController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.status = responseStatusError
	c.message = "-"
	c.data = make(map[string]interface{}, 0)
}

// SetResponse insert <key, value> to response body at field 'data'
func (c *BaseController) SetResponse(key string, value interface{}) {
	if key == "status" || key == "message" || key == "error" {
		logs.Error("Reserved key")
	} else {
		logs.Info(c.data)
		c.data[key] = value
	}
}

func (c *BaseController) setMessage(message []string) {
	c.message = strings.Join(message, " ")
}

// Success set response header code & response message
func (c *BaseController) Success(message ...string) {
	c.status = responseStatusSuccess
	c.setMessage(message)
	c.serveResponse()
}

// Success set response header code & response message
func (c *BaseController) Error(err error, message ...string) {
	c.status = responseStatusError
	if err != nil {
		c.data["error"] = err
	}
	c.setMessage(message)
	c.serveResponse()
}

func (c *BaseController) serveResponse() {
	if c.status == responseStatusSuccess {
		c.Ctx.ResponseWriter.WriteHeader(200)
	} else if c.status == responseStatusError {
		c.Ctx.ResponseWriter.WriteHeader(500)
	}

	// Move content 'data' to 'Data' to be exported
	m := c.data
	m["status"] = c.status
	m["message"] = c.message

	c.Data["json"] = m

	c.ServeJSON()
}
