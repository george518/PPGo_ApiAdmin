/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-24 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Group struct {
	Id           int
	GroupName    string
	Detail       string
	ApiPublicIds string
	CodeIds      string
	EnvIds       string
	Status       int
	CreateId     int
	UpdateId     int
	CreateTime   int64
	UpdateTime   int64
}

func (a *Group) TableName() string {
	return TableName("set_group")
}

func GroupAdd(a *Group) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func GroupGetByName(groupName string) (*Group, error) {
	a := new(Group)
	err := orm.NewOrm().QueryTable(TableName("set_group")).Filter("group_name", groupName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func GroupGetList(page, pageSize int, filters ...interface{}) ([]*Group, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Group, 0)
	query := orm.NewOrm().QueryTable(TableName("set_group"))
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

func GroupGetById(id int) (*Group, error) {
	r := new(Group)
	err := orm.NewOrm().QueryTable(TableName("set_group")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Group) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
