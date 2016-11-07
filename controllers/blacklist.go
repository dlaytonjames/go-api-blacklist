package controllers

import (
	"strconv"
	"strings"
	"time"
	"yk-black-list/models"

	"github.com/astaxie/beego"
)

// oprations for Blacklist
type BlacklistController struct {
	beego.Controller
}

func (c *BlacklistController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Add
// @Description 添加黑名单
// @Param	appid	query	string	false	"应用id"
// @Param	content	query	string	false	"黑名单数据"
// @Success 200 {object} models.Blacklist
// @Failure 403 body is empty
// @router /add.json [get]
func (c *BlacklistController) Add() {
	startTime := time.Now()
	timestamp := int(time.Now().Unix())
	appid, _ := strconv.Atoi(c.GetString("appid"))
	content := c.GetString("content")
	if len(content) == 0 || appid == 0 {
		c.Data["json"] = JsonFormat(0, "miss params", "", startTime)
	} else {
		record := &models.Blacklist{
			Id:        0,
			Appid:     appid,
			Content:   content,
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		}
		if _, err := models.AddBlacklist(record); err == nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = JsonFormat(1, "success", record.Id, startTime)
		} else {
			c.Data["json"] = JsonFormat(0, err.Error(), "", startTime)
		}
	}
	c.ServeJSON()
}

//批量添加黑名单
func (c *BlacklistController) BatchAdd() {
	startTime := time.Now()
	timestamp := int(time.Now().Unix())
	appid, _ := strconv.Atoi(c.GetString("appid"))
	content := c.GetString("content")
	if len(content) == 0 || appid == 0 {
		c.Data["json"] = JsonFormat(0, "miss params", "", startTime)
	} else {
		contentsArray := strings.Split(content, ",")
		multiRecord := []models.Blacklist{}
		for _, v := range contentsArray {
			record := models.Blacklist{
				Id:        0,
				Appid:     appid,
				Content:   v,
				CreatedAt: timestamp,
				UpdatedAt: timestamp,
			}
			multiRecord = append(multiRecord, record)
		}
		if successNumber, err := models.BatchAddBlacklist(multiRecord); err == nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = JsonFormat(1, "success", successNumber, startTime)
		} else {
			c.Data["json"] = JsonFormat(0, err.Error(), "", startTime)
		}
	}
	c.ServeJSON()
}

// @Title 判断黑名单是否属于当前应用
// @Description 判断黑名单是否属于当前应用
// @Param	appid		query	string	true		"The key for staticblock"
// @Param	content		query	string	true		"The key for staticblock"
// @Success 200 {object} models.Blacklist
// @Failure 403 :appid is empty
// @Failure 403 :content is empty
// @router /check.json [get]
func (c *BlacklistController) Check() {
	startTime := time.Now()
	idStr := c.GetString("appid")
	contentStr := c.GetString("content")
	appid, _ := strconv.Atoi(idStr)
	v, err := models.Check(appid, contentStr)
	if err != nil {
		c.Data["json"] = JsonFormat(0, err.Error(), "", startTime)
	} else {
		c.Data["json"] = JsonFormat(1, "success", v, startTime)
	}
	c.ServeJSON()
}

// @Title 批量判断黑名单是否属于当前应用
// @Description 批量判断黑名单是否属于当前应用
// @Param	appid		query	string	true		"The key for staticblock"
// @Param	content		query	string	true		"The key for staticblock"
// @Success 200 {object} models.Blacklist
// @Failure 403 :appid is empty
// @Failure 403 :content is empty
// @router /batch_check.json [get]
func (c *BlacklistController) BatchCheck() {
	startTime := time.Now()
	idStr := c.GetString("appid")
	contentStr := c.GetString("content")
	contentArray := strings.Split(contentStr, ",")
	appid, _ := strconv.Atoi(idStr)
	beego.Notice(contentArray)
	v, err := models.BatchCheck(appid, contentArray)
	if err != nil {
		c.Data["json"] = JsonFormat(0, err.Error(), "", startTime)
	} else {
		c.Data["json"] = JsonFormat(1, "success", v, startTime)
	}
	c.ServeJSON()
}

//删除指定数据
func (c *BlacklistController) Delete() {
	startTime := time.Now()
	idStr := c.GetString("appid")
	contentStr := c.GetString("content")
	appid, _ := strconv.Atoi(idStr)
	if err := models.DeleteBlacklist(appid, contentStr); err == nil {
		c.Data["json"] = JsonFormat(1, "success", "", startTime)
	} else {
		c.Data["json"] = JsonFormat(0, err.Error(), "", startTime)
	}
	c.ServeJSON()
}

/**************************************************************************************/

// @Title Get
// @Description get Blacklist by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Blacklist
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BlacklistController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetBlacklistById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title 获取黑名单列表
// @Description 获取黑名单列表
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Blacklist
// @Failure 403
// @router /list.json [get]
func (c *BlacklistController) GetAll() {
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
				c.Data["json"] = JsonFormat(0, "Error: invalid query key/value pair", "", startTime)
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllBlacklist(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = JsonFormat(0, err.Error(), "", startTime)
	} else {
		c.Data["json"] = JsonFormat(1, "success", l, startTime)
	}
	c.ServeJSON()
}
