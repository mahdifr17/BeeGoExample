package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"] = append(beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"],
        beego.ControllerComments{
            Method: "ViewAllTask",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"] = append(beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"],
        beego.ControllerComments{
            Method: "ViewTask",
            Router: `/:task_id:int`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"] = append(beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"],
        beego.ControllerComments{
            Method: "DeleteTask",
            Router: `/:task_id:int/delete`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"] = append(beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"],
        beego.ControllerComments{
            Method: "SetDoneTask",
            Router: `/:task_id:int/done`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"] = append(beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"],
        beego.ControllerComments{
            Method: "UpdateTask",
            Router: `/:task_id:int/update`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"] = append(beego.GlobalControllerRouter["github.com/mahdifr17/BeeGoExample/modules/todolist:TaskController"],
        beego.ControllerComments{
            Method: "InsertTask",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
