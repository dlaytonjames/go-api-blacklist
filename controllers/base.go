package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type BaseController struct {
	beego.Controller
}

// 1 success
// 0 fail
func JsonFormat(retcode int, retmsg string, retdata interface{}, stime time.Time) (json map[string]interface{}) {
	cost := time.Now().Sub(stime).Seconds()
	if retcode == 1 {
		json = map[string]interface{}{
			"retcode": retcode,
			"retdata": retdata,
			"retmsg":  retmsg,
			"cost":    cost,
		}
	} else {
		json = map[string]interface{}{
			"retcode": retcode,
			"retmsg":  retmsg,
			"cost":    cost,
		}
	}

	return json
}
