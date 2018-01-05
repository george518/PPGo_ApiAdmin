/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-24 11:48:17
***********************************************/
package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Group struct {
	Id         int
	GroupName  string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
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
	query := orm.NewOrm()

	groupName := ""
	status := 1
	sql := ""
	var total int64
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			if filters[k].(string) == "groupName" {
				groupName = filters[k+1].(string)
			}

			if filters[k].(string) == "status" {
				status = filters[k+1].(int)
			}
		}
	}
	if groupName == "" {
		sql = "SELECT * FROM pp_set_group WHERE status=? ORDER BY id DESC LIMIT ?,?"
		total, _ = query.Raw(sql, status, strconv.Itoa(offset), strconv.Itoa(pageSize)).QueryRows(&list)
	} else {
		sql = "SELECT * FROM pp_set_group WHERE status=? and group_name like ?  ORDER BY id DESC LIMIT ?,?"
		total, _ = query.Raw(sql, status, "%"+groupName+"%", strconv.Itoa(offset), strconv.Itoa(pageSize)).QueryRows(&list)
	}
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
