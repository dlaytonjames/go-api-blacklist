package controllers

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"yk-black-list/models"
	"yk-black-list/util"

	"github.com/astaxie/beego"
)

// oprations for Project
type ProjectController struct {
	beego.Controller
}

func (c *ProjectController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Get
// @Description 通过appid获取
// @Param	appid		query	string	true		"The key for staticblock"
// @Success 200 {object} models.Project
// @Failure 403 :appid is empty
// @router /get.json [get]
func (c *ProjectController) GetOne() {
	startTime := time.Now()
	idStr := c.GetString("appid")
	appid, _ := strconv.Atoi(idStr)
	//原：Checkratelimit(ip，应用id，动作方法)
	//新：Checkratelimit(应用id，动作方法，动作方法)
	rateLimitRes := util.CheckRateLimit(idStr, "GetOne", "GetOne")
	if rateLimitRes == true {
		c.Data["json"] = JsonFormat(0, "访问过快，请稍后访问!", "", startTime)
		c.ServeJSON()
		return
	}
	v, err := models.GetProjectById(appid)
	if err != nil {
		c.Data["json"] = JsonFormat(0, "fail", err.Error(), startTime)
	} else {
		c.Data["json"] = JsonFormat(1, "success", v, startTime)
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Project
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Project
// @Failure 403
// @router / [get]
func (c *ProjectController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0
	startTime := time.Now()

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllProject(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = JsonFormat(1, "success", l, startTime)
	}
	c.ServeJSON()
}
