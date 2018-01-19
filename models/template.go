/**********************************************
** @Des: markdown模板
** @Author: haodaquan
** @Date:   2018-01-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-16 11:48:17
***********************************************/
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Template struct {
	Id           int
	TemplateName string
	Detail       string
	Status       int
	CreateId     int
	UpdateId     int
	CreateTime   int64
	UpdateTime   int64
}

func (a *Template) TableName() string {
	return TableName("set_template")
}

func TemplateGetList(page, pageSize int, filters ...interface{}) ([]*Template, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Template, 0)
	query := orm.NewOrm().QueryTable(TableName("set_template"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func TemplateAdd(a *Template) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func TemplateGetById(id int) (Template, error) {
	var list Template
	query := orm.NewOrm().QueryTable(TableName("set_template"))
	query.Filter("id", id).Filter("status", 1).One(&list)
	return list, nil
}

func (a *Template) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Template) Delete(id int64, update_id int) (int64, error) {
	sql := "UPDATE pp_set_template SET status=0,update_id=?,update_time=? WHERE id=?"
	res, err := orm.NewOrm().Raw(sql, update_id, time.Now().Unix(), id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num, nil
	}
	return 0, err
}
