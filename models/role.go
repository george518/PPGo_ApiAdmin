/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-14 15:24:51
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:48:52
***********************************************/
package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id         int
	RoleName   string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Role) TableName() string {
	return TableName("uc_role")
}

func RoleGetList(page, pageSize int, filters ...interface{}) ([]*Role, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Role, 0)
	query := orm.NewOrm()

	roleName := ""
	sql := ""
	var total int64
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			if filters[k].(string) == "roleName" {
				roleName = filters[k+1].(string)
			}
		}
	}
	if roleName == "" {
		sql = "SELECT * FROM pp_uc_role ORDER BY id DESC LIMIT ?,?"
		total, _ = query.Raw(sql, strconv.Itoa(offset), strconv.Itoa(pageSize)).QueryRows(&list)
	} else {
		sql = "SELECT * FROM pp_uc_role WHERE role_name like ?  ORDER BY id DESC LIMIT ?,?"
		total, _ = query.Raw(sql, "%"+roleName+"%", strconv.Itoa(offset), strconv.Itoa(pageSize)).QueryRows(&list)
	}
	return list, total
}

func RoleAdd(role *Role) (int64, error) {
	id, err := orm.NewOrm().Insert(role)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func RoleGetById(id int) (*Role, error) {
	r := new(Role)
	err := orm.NewOrm().QueryTable(TableName("uc_role")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
