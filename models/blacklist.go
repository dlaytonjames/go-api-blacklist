package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Blacklist struct {
	Id        int    `orm:"column(id);auto"`
	Appid     int    `orm:"column(appid)"`
	Content   string `orm:"column(content);size(20)"`
	CreatedAt int    `orm:"column(created_at)"`
	UpdatedAt int    `orm:"column(updated_at)"`
}

func (t *Blacklist) TableName() string {
	return "blacklist"
}

func init() {
	orm.RegisterModel(new(Blacklist))
}

//判断黑名单是否属于当前应用
func Check(appid int, content string) (v *Blacklist, err error) {
	o := orm.NewOrm()
	dbRec := &Blacklist{}
	err = o.QueryTable("blacklist").Filter("appid", appid).Filter("content", content).One(dbRec)
	return dbRec, err
}

//批量判断黑名单是否属于当前应用
func BatchCheck(appid int, content []string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	dbRec := []Blacklist{}
	o.QueryTable("blacklist").Filter("appid", appid).Filter("content__in", content).All(&dbRec)
	// trim unused fields
	for _, v := range dbRec {
		if len(v.Content) != 0 {
			ml = append(ml, v.Content)
		}
	}

	return ml, err
}

// AddBlacklist insert a new Blacklist into database and returns
// last inserted Id on success.
func AddBlacklist(m *Blacklist) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//批量添加
func BatchAddBlacklist(m []Blacklist) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.InsertMulti(100, m)
	return
}

//删除指定数据
func DeleteBlacklist(appid int, content string) (err error) {
	o := orm.NewOrm()
	//无指定数据，也会返回正确，这个无需担心
	_, err = o.QueryTable("blacklist").Filter("appid", appid).Filter("content", content).Delete()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

/*****************************************************************/

// GetBlacklistById retrieves Blacklist by Id. Returns error if
// Id doesn't exist
func GetBlacklistById(id int) (v *Blacklist, err error) {
	o := orm.NewOrm()
	v = &Blacklist{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBlacklist retrieves all Blacklist matches certain condition. Returns empty list if
// no records exist
func GetAllBlacklist(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	fields = []string{"Content"}
	qs := o.QueryTable(new(Blacklist))
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

	var l []Blacklist
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				/***************/
				//m := make(map[string]interface{})
				//val := reflect.ValueOf(v)
				//for _, fname := range fields {
				//	m[fname] = val.FieldByName(fname).Interface()
				//}
				//ml = append(ml, m)
				/***************/
				var m interface{}
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m = val.FieldByName(fname).Interface()
					ml = append(ml, m)
				}

			}
		}
		return ml, nil
	}
	return nil, err
}
