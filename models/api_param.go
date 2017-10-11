/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-09 11:48:17
***********************************************/
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ApiParam struct {
	Id         int
	DetailId   int64
	ApiKey     string
	ApiType    string
	ApiValue   string
	ApiDetail  string
	IsNull     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *ApiParam) TableName() string {
	return TableName("api_param")
}

func ApiParamAdd(a *ApiParam) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func ApiParamGetById(id int) ([]*ApiParam, error) {
	list := make([]*ApiParam, 0)
	query := orm.NewOrm().QueryTable(TableName("api_param"))
	query.Filter("detail_id", id).Filter("status", 1).All(&list)
	return list, nil
}

func (a *ApiParam) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *ApiParam) Delete(id int64, update_id int) (int64, error) {
	sql := "UPDATE pp_api_param SET status=0,update_id=?,update_time=? WHERE detail_id=?"
	res, err := orm.NewOrm().Raw(sql, update_id, time.Now().Unix(), id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num, nil
	}
	return 0, err
}

// func ApiParamGetByName(ApiName string) (*ApiParam, error) {
// 	a := new(ApiParam)
// 	err := orm.NewOrm().QueryTable(TableName("api_param")).Filter("source_name", ApiName).One(a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return a, nil
// }

// func ApiParamGetList(page, pageSize int, filters ...interface{}) ([]*ApiParam, int64) {
// 	offset := (page - 1) * pageSize
// 	list := make([]*ApiParam, 0)
// 	query := orm.NewOrm().QueryTable(TableName("api_param"))
// 	if len(filters) > 0 {
// 		l := len(filters)
// 		for k := 0; k < l; k += 2 {
// 			query = query.Filter(filters[k].(string), filters[k+1])
// 		}
// 	}
// 	total, _ := query.Count()
// 	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

// 	return list, total
// }

// func ApiParamGetById(id int) (*ApiParam, error) {
// 	r := new(ApiParam)
// 	err := orm.NewOrm().QueryTable(TableName("api_param")).Filter("id", id).One(r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }

// func (a *ApiParam) Update(fields ...string) error {
// 	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
// 		return err
// 	}
// 	return nil
// }
