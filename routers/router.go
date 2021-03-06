// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"yk-black-list/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/yk",

		beego.NSNamespace("/blacklist",
			beego.NSRouter("/add.json", &controllers.BlacklistController{}, "get:Add"),
			beego.NSRouter("/batch_add.json", &controllers.BlacklistController{}, "get:BatchAdd"),
			beego.NSRouter("/check.json", &controllers.BlacklistController{}, "get:Check"),
			beego.NSRouter("/batch_check.json", &controllers.BlacklistController{}, "get:BatchCheck"),
			beego.NSRouter("/delete.json", &controllers.BlacklistController{}, "get:Delete"),
			beego.NSRouter("/list.json", &controllers.BlacklistController{}, "get:GetAll"),
			beego.NSInclude(
				&controllers.BlacklistController{},
			),
		),

		beego.NSNamespace("/project",
			beego.NSRouter("/get.json", &controllers.ProjectController{}, "get:GetOne"),
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
