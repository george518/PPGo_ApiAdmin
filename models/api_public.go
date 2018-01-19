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

type ApiPublic struct {
	Id            int
	ApiPublicName string
	Detail        string
	Sort          int
	Status        int
	CreateId      int
	UpdateId      int
	CreateTime    int64
	UpdateTime    int64
}

func (a *ApiPublic) TableName() string {
	return TableName("api_public")
}

func ApiPublicGetList(page, pageSize int, filters ...interface{}) ([]*ApiPublic, int64) {
	offset := (page - 1) * pageSize
	list := make([]*ApiPublic, 0)
	query := orm.NewOrm().QueryTable(TableName("api_public"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("sort").Limit(pageSize, offset).All(&list)

	return list, total
}

func ApiPublicAdd(a *ApiPublic) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func ApiPublicGetById(id int) (ApiPublic, error) {
	var list ApiPublic
	query := orm.NewOrm().QueryTable(TableName("api_public"))
	query.Filter("id", id).Filter("status", 1).One(&list)
	return list, nil
}

func ApiPublicGetByIds(ids string) ([]*ApiPublic, error) {
	list := make([]*ApiPublic, 0)
	sql := "SELECT * FROM pp_api_public WHERE id in(" + ids + ")"
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list, nil
}

func (a *ApiPublic) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *ApiPublic) Delete(id int64, update_id int) (int64, error) {
	sql := "UPDATE pp_api_public SET status=0,update_id=?,update_time=? WHERE id=?"
	res, err := orm.NewOrm().Raw(sql, update_id, time.Now().Unix(), id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num, nil
	}
	return 0, err
}
