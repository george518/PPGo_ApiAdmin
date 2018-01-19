/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-25 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Env struct {
	Id         int
	EnvName    string
	EnvHost    string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Env) TableName() string {
	return TableName("set_env")
}

func EnvAdd(a *Env) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func EnvGetByName(EnvName string) (*Env, error) {
	a := new(Env)
	err := orm.NewOrm().QueryTable(TableName("set_env")).Filter("env_name", EnvName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func EnvGetList(page, pageSize int, filters ...interface{}) ([]*Env, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Env, 0)
	query := orm.NewOrm().QueryTable(TableName("set_env"))
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

func EnvGetByIds(ids string) ([]*Env, error) {
	list := make([]*Env, 0)
	sql := "SELECT * FROM pp_set_env WHERE id in(" + ids + ")"
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list, nil
}

func EnvGetById(id int) (*Env, error) {
	r := new(Env)
	err := orm.NewOrm().QueryTable(TableName("set_env")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Env) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
