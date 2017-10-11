/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-09 11:48:17
***********************************************/
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type ApiDetail struct {
	Id           int
	SourceId     int
	Method       int
	ApiName      string
	ApiUrl       string
	ProtocolType int
	Result       string
	Example      string
	Detail       string
	Status       int
	CreateId     int
	AuditId      int
	UpdateId     int
	CreateTime   int64
	UpdateTime   int64
	AuditTime    int64
}
type ApiDetails struct {
	Id           int
	SourceId     int
	Method       int
	ApiName      string
	ApiUrl       string
	ProtocolType int
	Result       string
	Example      string
	Detail       string
	Status       int
	CreateId     int
	AuditId      int
	UpdateId     int
	CreateTime   int64
	UpdateTime   int64
	AuditTime    int64
	CreateName   string
	UpdateName   string
	AuditName    string
}

func (a *ApiDetail) TableName() string {
	return TableName("api_detail")
}

func ApiDetailAdd(a *ApiDetail) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func ApiDetailGetById(id int) (*ApiDetail, error) {
	r := new(ApiDetail)
	err := orm.NewOrm().QueryTable(TableName("api_detail")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func ApiDetailsGetById(id int) ([]*ApiDetails, error) {
	list := make([]*ApiDetails, 0)
	sql := "SELECT pp_api_detail.*,a.real_name as create_name,b.real_name as update_name,c.real_name as audit_name FROM pp_api_detail LEFT JOIN pp_uc_admin as a ON pp_api_detail.create_id=a.id LEFT JOIN pp_uc_admin as b ON pp_api_detail.update_id=b.id LEFT JOIN pp_uc_admin as c ON pp_api_detail.audit_id=c.id WHERE pp_api_detail.source_id=?"
	orm.NewOrm().Raw(sql, id).QueryRows(&list)
	fmt.Println(list)
	return list, nil
}

// func ApiDetailGetByName(ApiName string) (*ApiDetail, error) {
// 	a := new(ApiDetail)
// 	err := orm.NewOrm().QueryTable(TableName("api_detail")).Filter("source_name", ApiName).One(a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return a, nil
// }

// func ApiDetailGetList(page, pageSize int, filters ...interface{}) ([]*ApiDetail, int64) {
// 	offset := (page - 1) * pageSize
// 	list := make([]*ApiDetail, 0)
// 	query := orm.NewOrm().QueryTable(TableName("api_detail"))
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

func (a *ApiDetail) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
