/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-09 11:48:17
***********************************************/
package models

type ApiDetail struct {
	Id         int
	SourceId   int
	SourceName string
	Status     int
	CreateId   int
	AuditId    int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	AuditTime  int64
}
