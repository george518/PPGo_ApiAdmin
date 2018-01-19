/**********************************************
** @Des: 接口资源
** @Author: haodaquan
** @Date:   2018-01-14 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-14 15:42:43
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type ApiSource struct {
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

func (a *ApiSource) TableName() string {
	return TableName("api_source")
}

func ApiSourceAdd(a *ApiSource) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func ApiSourceGetByName(ApiSourceName string) (*ApiSource, error) {
	a := new(ApiSource)
	err := orm.NewOrm().QueryTable(TableName("api_source")).Filter("source_name", ApiSourceName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func ApiSourceGetList(page, pageSize int, filters ...interface{}) ([]*ApiSource, int64) {
	offset := (page - 1) * pageSize
	list := make([]*ApiSource, 0)
	query := orm.NewOrm().QueryTable(TableName("api_source"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	// total := int64(12)
	return list, total
}

func ApiSourceGetById(id int) (*ApiSource, error) {
	r := new(ApiSource)
	err := orm.NewOrm().QueryTable(TableName("api_source")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *ApiSource) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
