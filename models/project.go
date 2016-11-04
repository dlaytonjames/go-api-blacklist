package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id        int    `orm:"column(id);auto"`
	Name      string `orm:"column(name);size(64)"`
	Appid     int    `orm:"column(appid)"`
	Appkey    string `orm:"column(appkey);size(64)"`
	Status    int16  `orm:"column(status)"`
	Desc      string `orm:"column(desc);size(255)"`
	BelongUid int    `orm:"column(belong_uid)"`
	CreatedAt int    `orm:"column(created_at)"`
	UpdatedAt int    `orm:"column(updated_at)"`
}

func (t *Project) TableName() string {
	return "project"
}

func init() {
	orm.RegisterModel(new(Project))
}

// AddProject insert a new Project into database and returns
// last inserted Id on success.
func AddProject(m *Project) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProjectById retrieves Project by Id. Returns error if
// Id doesn't exist
func GetProjectById(appid int) (v *Project, err error) {
	o := orm.NewOrm()
	dbRec := &Project{}
	err = o.QueryTable("project").Filter("appid", appid).One(dbRec)
	return dbRec, err
}

// GetAllProject retrieves all Project matches certain condition. Returns empty list if
// no records exist
func GetAllProject(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Project))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Project
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateProject updates Project by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectById(m *Project) (err error) {
	o := orm.NewOrm()
	v := Project{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProject deletes Project by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProject(id int) (err error) {
	o := orm.NewOrm()
	v := Project{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Project{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
