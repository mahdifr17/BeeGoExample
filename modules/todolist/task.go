package todolist

import (
	"github.com/astaxie/beego/logs"
	"github.com/mahdifr17/BeeGoExample/modules/db"
)

// Task represent table to store to do list
type Task struct {
	ID int `json:"i_d" orm:auto;pk;column(i_d)`
	db.Model
	Description string `json:"description" orm:"size(256)"`
	IsDone      bool   `json:"is_done" orm:"default(false)"`
}

func init() {
	db.RegisterModel(new(Task))
}

func getAllTask(orm db.ORM) ([]*Task, error) {
	var tasks []*Task
	_, err := orm.QueryTable(Task{}).
		Exclude("is_deleted", true).
		All(&tasks)
	if err != nil {
		logs.Error("Error db: ", err)
	}
	return tasks, err
}

func getTask(orm db.ORM, taskID int) (*Task, error) {
	var task Task
	err := orm.QueryTable(Task{}).
		Filter("i_d", taskID).
		Exclude("is_deleted", true).One(&task)
	if err != nil {
		logs.Error("Error db: ", err)
	}
	return &task, err
}

func insertTask(orm db.ORM, task *Task) (int64, error) {
	logs.Info("Insert task: ", task)
	return orm.Insert(task)
}

func updateTask(orm db.ORM, task *Task) (int64, error) {
	logs.Info("Update task: ", task)
	// Update column, expect created_at
	return orm.Update(task, "updated_at", "description", "is_done", "is_deleted")
}

func setDone(orm db.ORM, taskID int) (int64, error) {
	task, err := getTask(orm, taskID)
	if err != nil {
		logs.Error(err, "Not found")
		return 0, err
	}
	task.IsDone = true
	logs.Info("Task has done: ", task)
	return orm.Update(task, "updated_at", "is_done")
}

func setDeleted(orm db.ORM, taskID int) (int64, error) {
	task, err := getTask(orm, taskID)
	if err != nil {
		logs.Error(err, "Not found")
		return 0, err
	}
	task.IsDeleted = true
	logs.Info("Delete task: ", task)
	return orm.Update(task, "updated_at", "is_deleted")
}
