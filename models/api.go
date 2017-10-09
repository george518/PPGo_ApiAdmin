/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-26 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Api struct {
	Id         int
	GroupId    int
	SourceName string
	Status     int
	CreateId   int
	AuditId    int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	AuditTime  int64
}

func (a *Api) TableName() string {
	return TableName("api_source")
}

func ApiAdd(a *Api) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func ApiGetByName(ApiName string) (*Api, error) {
	a := new(Api)
	err := orm.NewOrm().QueryTable(TableName("api_source")).Filter("source_name", ApiName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func ApiGetList(page, pageSize int, filters ...interface{}) ([]*Api, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Api, 0)
	query := orm.NewOrm().QueryTable(TableName("api_source"))
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

func ApiGetById(id int) (*Api, error) {
	r := new(Api)
	err := orm.NewOrm().QueryTable(TableName("api_source")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Api) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
