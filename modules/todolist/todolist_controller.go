package todolist

import (
	"encoding/json"

	"github.com/mahdifr17/BeeGoExample/controllers"
	"github.com/mahdifr17/BeeGoExample/modules/db"
)

// TaskController is controller for crud tasks
type TaskController struct {
	controllers.BasicController
}

// Prepare runs after Init before request function execution
// Used to initialize Orm instance
func (c *TaskController) Prepare() {
	// Override method Prepare on BaseController -> call BaseController.Prepare() inside this method
	c.BaseController.Prepare()
	c.Orm = db.NewOrm()
}

// ViewAllTask method
// @Title ViewAllTask
// @Description method to get all task in db
// @Success 200 {[]Task} Task
// @Error 500 body is empty
// @router / [get]
func (c *TaskController) ViewAllTask() {
	if taskList, err := getAllTask(c.Orm); err != nil {
		c.Error(err, "Not found")
	} else {
		c.SetResponse("tasks", taskList)
		c.Success("OK")
	}
}

// ViewTask method
// @Title ViewTask
// @Description method to get specific task in db
// @Param task_id	query	int	true	"task_id"
// @Success 200 {Task} Task
// @Error 500 body is empty
// @router /:task_id:int [get]
func (c *TaskController) ViewTask() {
	taskID, _ := c.GetInt(":task_id") // Error parsing handle by routing
	if task, errNotFound := getTask(c.Orm, taskID); errNotFound != nil {
		c.Error(errNotFound, "Not found")
	} else {
		c.SetResponse("task", task)
		c.Success("OK")
	}
}

// InsertTask method
// @Title InsertTask
// @Description method to insert task into db
// @Success 200 {Task} Task
// @Error 500 body is empty
// @router /insert [post]
func (c *TaskController) InsertTask() {
	var request *Task
	if errRequest := json.Unmarshal(c.Ctx.Input.RequestBody, &request); errRequest != nil {
		c.Error(errRequest, "Invalid Request")
	} else {
		if _, errInsert := insertTask(c.Orm, request); errInsert != nil {
			c.Error(errInsert, "Error insert")
		} else {
			c.Success("Insert success")
		}
	}
}

// UpdateTask method
// @Title UpdateTask
// @Description method to update task
// @Param task_id	query	int	true	"task_id"
// @Success 200 {string} string
// @Error 500 body is empty
// @router /:task_id:int/update [put]
func (c *TaskController) UpdateTask() {
	taskID, _ := c.GetInt(":task_id") // Error parsing handle by routing

	var request *Task
	if errRequest := json.Unmarshal(c.Ctx.Input.RequestBody, &request); errRequest != nil {
		c.Error(errRequest, "Invalid Request")
	} else {
		request.ID = taskID
		if _, errUpdate := updateTask(c.Orm, request); errUpdate != nil {
			c.Error(errUpdate, "Error update")
		} else {
			c.Success("Update success")
		}
	}
}

// DeleteTask method
// @Title DeleteTask
// @Description method to soft delete task
// @Param task_id	query	int	true	"task_id"
// @Success 200 {string} string
// @Error 500 body is empty
// @router /:task_id:int/delete [delete]
func (c *TaskController) DeleteTask() {
	taskID, _ := c.GetInt(":task_id") // Error parsing handle by routing

	if _, errDelete := setDeleted(c.Orm, taskID); errDelete != nil {
		c.Error(errDelete, "Error delete")
	} else {
		c.Success("Update success")
	}
}

// SetDoneTask method
// @Title SetDoneTask
// @Description method to set done task
// @Param task_id	query	int	true	"task_id"
// @Success 200 {string} string
// @Error 500 body is empty
// @router /:task_id:int/done [put]
func (c *TaskController) SetDoneTask() {
	taskID, _ := c.GetInt(":task_id") // Error parsing handle by routing

	if _, errSetDone := setDone(c.Orm, taskID); errSetDone != nil {
		c.Error(errSetDone, "Error set done")
	} else {
		c.Success("Update success")
	}
}
