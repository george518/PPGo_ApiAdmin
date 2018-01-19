/**********************************************
** @Des: 公共文档设置
** @Author: haodaquan
** @Date:   2018-01-16 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-16 17:48:30
***********************************************/
package controllers

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

var ()

type ApiPublicController struct {
	BaseController
}

func (self *ApiPublicController) List() {
	self.Data["pageTitle"] = "公共文档"
	self.display()
}

func (self *ApiPublicController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	ApiPublicName := strings.TrimSpace(self.GetString("ApiPublicName"))

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	if ApiPublicName != "" {
		filters = append(filters, "api_public_name__icontains", ApiPublicName)
	}
	result, count := models.ApiPublicGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["api_public_name"] = v.ApiPublicName
		row["detail"] = v.Detail
		row["sort"] = v.Sort
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *ApiPublicController) Add() {
	self.Data["pageTitle"] = "新增公共文档"
	tmplates := templateLists()
	self.Data["templates"] = tmplates
	self.display()
}

func (self *ApiPublicController) Edit() {
	id, _ := self.GetInt("id", 0)
	detail, _ := models.ApiPublicGetById(id)
	row := make(map[string]interface{})
	row["id"] = detail.Id
	row["api_public_name"] = detail.ApiPublicName
	row["detail"] = detail.Detail
	row["sort"] = detail.Sort

	self.Data["pageTitle"] = "查看 " + detail.ApiPublicName
	self.Data["Detail"] = row

	tmplates := templateLists()
	self.Data["templates"] = tmplates
	self.display()
}

func (self *ApiPublicController) AjaxSave() {

	Pub_id, _ := self.GetInt("id")
	if Pub_id == 0 {
		ApiPublic := new(models.ApiPublic)

		ApiPublic.ApiPublicName = strings.TrimSpace(self.GetString("api_public_name"))
		ApiPublic.Detail = strings.TrimSpace(self.GetString("detail"))
		ApiPublic.Sort, _ = self.GetInt("sort", 99)
		ApiPublic.CreateId = self.userId
		ApiPublic.UpdateId = self.userId
		ApiPublic.CreateTime = time.Now().Unix()
		ApiPublic.UpdateTime = time.Now().Unix()
		ApiPublic.Status = 1
		_, err := models.ApiPublicAdd(ApiPublic)
		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}
	//修改
	ApiPublic, _ := models.ApiPublicGetById(Pub_id)
	ApiPublic.Id, _ = self.GetInt("id")
	ApiPublic.ApiPublicName = strings.TrimSpace(self.GetString("api_public_name"))
	ApiPublic.Detail = strings.TrimSpace(self.GetString("detail"))
	ApiPublic.Sort, _ = self.GetInt("sort", 99)
	ApiPublic.UpdateId = self.userId
	ApiPublic.UpdateTime = time.Now().Unix()
	ApiPublic.Status = 1
	if err := ApiPublic.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *ApiPublicController) AjaxDel() {

	Pub_id, _ := self.GetInt("id")
	Api, _ := models.ApiPublicGetById(Pub_id)
	Api.UpdateTime = time.Now().Unix()
	Api.UpdateId = self.userId
	Api.Status = 0
	Api.Id = Pub_id

	//TODO 判断是否被使用

	if err := Api.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
