package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"],
		beego.ControllerComments{
			"Add",
			`/add.json`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"],
		beego.ControllerComments{
			"Check",
			`/check.json`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"],
		beego.ControllerComments{
			"BatchCheck",
			`/batch_check.json`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:BlacklistController"],
		beego.ControllerComments{
			"GetAll",
			`/list.json`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:ProjectController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:ProjectController"],
		beego.ControllerComments{
			"GetOne",
			`/get.json`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:ProjectController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:ProjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:UserController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:UserController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:UserController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:UserController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["yk-black-list/controllers:UserController"] = append(beego.GlobalControllerRouter["yk-black-list/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
