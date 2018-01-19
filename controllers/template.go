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

type TemplateController struct {
	BaseController
}

func (self *TemplateController) List() {
	self.Data["pageTitle"] = "模板管理"
	self.display()
}

func (self *TemplateController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	TemplateName := strings.TrimSpace(self.GetString("TemplateName"))

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	if TemplateName != "" {
		filters = append(filters, "template_name__icontains", TemplateName)
	}
	result, count := models.TemplateGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["template_name"] = v.TemplateName
		row["detail"] = v.Detail
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *TemplateController) Add() {
	self.Data["pageTitle"] = "新增模板"

	self.display()
}

func (self *TemplateController) Edit() {
	id, _ := self.GetInt("id", 0)
	detail, _ := models.TemplateGetById(id)
	row := make(map[string]interface{})
	row["id"] = detail.Id
	row["template_name"] = detail.TemplateName
	row["detail"] = detail.Detail
	self.Data["pageTitle"] = "查看 " + detail.TemplateName
	self.Data["Detail"] = row
	self.display()
}

func (self *TemplateController) AjaxSave() {

	Pub_id, _ := self.GetInt("id")
	if Pub_id == 0 {
		Template := new(models.Template)

		Template.TemplateName = strings.TrimSpace(self.GetString("template_name"))
		Template.Detail = strings.TrimSpace(self.GetString("detail"))
		Template.CreateId = self.userId
		Template.UpdateId = self.userId
		Template.CreateTime = time.Now().Unix()
		Template.UpdateTime = time.Now().Unix()
		Template.Status = 1
		_, err := models.TemplateAdd(Template)
		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}
	//修改
	Template, _ := models.TemplateGetById(Pub_id)
	Template.Id, _ = self.GetInt("id")
	Template.TemplateName = strings.TrimSpace(self.GetString("template_name"))
	Template.Detail = strings.TrimSpace(self.GetString("detail"))

	Template.UpdateId = self.userId
	Template.UpdateTime = time.Now().Unix()
	Template.Status = 1
	if err := Template.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *TemplateController) AjaxDel() {

	Pub_id, _ := self.GetInt("id")
	Api, _ := models.TemplateGetById(Pub_id)
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
