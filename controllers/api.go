/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:41
***********************************************/
package controllers

import (
	"PPGo_ApiAdmin/models"
	"fmt"
	"strings"
	"time"
)

type ApiController struct {
	BaseController
}

func (self *ApiController) List() {
	self.Data["pageTitle"] = "API接口"
	self.Data["ApiCss"] = true
	// 分组
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.GroupGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["group_name"] = v.GroupName
		list[k] = row
	}
	self.Data["Groups"] = list

	//所有接口
	filters_source := make([]interface{}, 0)
	filters_source = append(filters_source, "status__in", []int{1, 2, 3})
	result_source, _ := models.ApiGetList(1, 1000, filters_source...)
	list_source := make([]map[string]interface{}, len(result_source))
	for k, v := range result_source {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["source_name"] = v.SourceName
		row["group_id"] = v.GroupId
		row["status"] = v.Status
		list_source[k] = row
	}
	self.Data["Source"] = list_source

	self.display()
}

//显示
func (self *ApiController) Show() {
	id, _ := self.GetInt("id", 0)
	fmt.Println(id)
	self.TplName = "api/info.html"
}

func (self *ApiController) Add() {
	self.Data["pageTitle"] = "新增资源"
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.GroupGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["group_name"] = v.GroupName
		list[k] = row
	}
	self.Data["Groups"] = list
	self.display()
}

func (self *ApiController) Edit() {
	self.Data["pageTitle"] = "编辑API"

	id, _ := self.GetInt("id", 0)
	Api, err := models.ApiGetById(id)
	if err != nil {
		self.Ctx.WriteString("数据不存在")
		return
	}
	row := make(map[string]interface{})
	row["id"] = Api.Id
	row["source_name"] = Api.SourceName
	row["group_id"] = int(Api.GroupId)
	self.Data["Source"] = row

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.GroupGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["group_name"] = v.GroupName
		list[k] = row
	}
	self.Data["Groups"] = list
	self.display()
}

func (self *ApiController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.ApiGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["source_name"] = v.SourceName
		row["group_id"] = v.GroupId
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

//存储资源
func (self *ApiController) AjaxSave() {
	Api_id, _ := self.GetInt("id")
	if Api_id == 0 {
		Api := new(models.Api)

		Api.SourceName = strings.TrimSpace(self.GetString("source_name"))
		Api.GroupId, _ = self.GetInt("group_id")
		Api.CreateId = self.userId
		Api.UpdateId = self.userId
		Api.CreateTime = time.Now().Unix()
		Api.UpdateTime = time.Now().Unix()
		Api.Status = 2

		// 检查登录名是否已经存在
		_, err := models.ApiGetByName(Api.SourceName)

		if err == nil {
			self.ajaxMsg("资源名已经存在", MSG_ERR)
		}

		if _, err := models.ApiAdd(Api); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	ApiUpdate, _ := models.ApiGetById(Api_id)
	// 修改
	ApiUpdate.SourceName = strings.TrimSpace(self.GetString("source_name"))
	ApiUpdate.GroupId, _ = self.GetInt("group_id")
	ApiUpdate.UpdateId = self.userId
	ApiUpdate.UpdateTime = time.Now().Unix()
	ApiUpdate.Status = 2

	if err := ApiUpdate.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *ApiController) AjaxDel() {

	Api_id, _ := self.GetInt("id")
	Api, _ := models.ApiGetById(Api_id)
	Api.UpdateTime = time.Now().Unix()
	Api.UpdateId = self.userId
	Api.Status = 0
	Api.Id = Api_id

	//TODO 判断是否暂用API

	if err := Api.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

//新增接口实例
func (self *ApiController) AddApi() {
	self.Data["pageTitle"] = ""
	source_id, _ := self.GetInt("sid")
	self.Data["Sid"] = source_id
	//查询条件
	self.display()
}

func (self *ApiController) AjaxApiSave() {

	Api_id, _ := self.GetInt("id")
	if Api_id == 0 {
		Api := new(models.Api)

		Api.SourceName = strings.TrimSpace(self.GetString("source_name"))
		Api.GroupId, _ = self.GetInt("group_id")
		Api.CreateId = self.userId
		Api.UpdateId = self.userId
		Api.CreateTime = time.Now().Unix()
		Api.UpdateTime = time.Now().Unix()
		Api.Status = 2

		// 检查登录名是否已经存在
		_, err := models.ApiGetByName(Api.SourceName)

		if err == nil {
			self.ajaxMsg("资源名已经存在", MSG_ERR)
		}

		if _, err := models.ApiAdd(Api); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	ApiUpdate, _ := models.ApiGetById(Api_id)
	// 修改
	ApiUpdate.SourceName = strings.TrimSpace(self.GetString("source_name"))
	ApiUpdate.GroupId, _ = self.GetInt("group_id")
	ApiUpdate.UpdateId = self.userId
	ApiUpdate.UpdateTime = time.Now().Unix()
	ApiUpdate.Status = 2

	if err := ApiUpdate.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
