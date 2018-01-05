/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-09 12:53:05
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-25 18:50:54
***********************************************/
package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

type CodeController struct {
	BaseController
}

func (self *CodeController) List() {
	self.Data["pageTitle"] = "状态码设置"
	self.display()
}

func (self *CodeController) Add() {
	self.Data["pageTitle"] = "新增状态码"
	self.display()
}

func (self *CodeController) Edit() {
	self.Data["pageTitle"] = "编辑状态码"

	id, _ := self.GetInt("id", 0)
	Code, _ := models.CodeGetById(id)
	row := make(map[string]interface{})
	row["id"] = Code.Id
	row["code"] = Code.Code
	row["desc"] = Code.Desc
	row["detail"] = Code.Detail
	self.Data["code"] = row
	self.display()
}

func (self *CodeController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}
	code := strings.TrimSpace(self.GetString("code"))

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if code != "" {
		filters = append(filters, "code", code)
	}
	result, count := models.CodeGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["code"] = v.Code
		row["detail"] = v.Detail
		row["desc"] = v.Desc
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *CodeController) AjaxSave() {
	Code_id, _ := self.GetInt("id")
	if Code_id == 0 {
		Code := new(models.Code)

		Code.Code = strings.TrimSpace(self.GetString("code"))
		Code.Desc = strings.TrimSpace(self.GetString("desc"))
		Code.Detail = strings.TrimSpace(self.GetString("detail"))
		Code.CreateId = self.userId
		Code.UpdateId = self.userId
		Code.CreateTime = time.Now().Unix()
		Code.UpdateTime = time.Now().Unix()
		Code.Status = 1

		res, err := models.CodeGetByName(Code.Code)
		fmt.Println(res)
		if err == nil {
			self.ajaxMsg("状态码已经存在", MSG_ERR)
		}

		if _, err := models.CodeAdd(Code); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	CodeUpdate, _ := models.CodeGetById(Code_id)
	// 修改
	CodeUpdate.Code = strings.TrimSpace(self.GetString("code"))
	CodeUpdate.Desc = strings.TrimSpace(self.GetString("desc"))
	CodeUpdate.Detail = strings.TrimSpace(self.GetString("detail"))
	CodeUpdate.UpdateId = self.userId
	CodeUpdate.UpdateTime = time.Now().Unix()
	CodeUpdate.Status = 1

	if err := CodeUpdate.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *CodeController) AjaxDel() {

	Code_id, _ := self.GetInt("id")
	Code, _ := models.CodeGetById(Code_id)
	Code.UpdateTime = time.Now().Unix()
	Code.UpdateId = self.userId
	Code.Status = 0
	Code.Id = Code_id

	if err := Code.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
