/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-09 12:53:05
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-24 18:50:54
***********************************************/
package controllers

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

type GroupController struct {
	BaseController
}

func (self *GroupController) List() {
	self.Data["pageTitle"] = "分组设置"
	self.display()
}

func (self *GroupController) Add() {
	self.Data["pageTitle"] = "新增分组"
	envlists := envLists()
	self.Data["envlists"] = envlists

	codelists := codeLists()
	self.Data["codelists"] = codelists

	apiPublicLists := apiPublicLists()
	self.Data["apiPublicLists"] = apiPublicLists

	self.display()
}

func (self *GroupController) Edit() {
	self.Data["pageTitle"] = "编辑分组"

	id, _ := self.GetInt("id", 0)
	group, _ := models.GroupGetById(id)
	row := make(map[string]interface{})
	row["id"] = group.Id
	row["group_name"] = group.GroupName
	row["detail"] = group.Detail
	row["env_ids"] = group.EnvIds
	row["code_ids"] = group.CodeIds
	row["api_public_ids"] = group.ApiPublicIds
	self.Data["group"] = row

	envlists := envLists()
	self.Data["envlists"] = envlists

	codelists := codeLists()
	self.Data["codelists"] = codelists

	apiPublicLists := apiPublicLists()
	self.Data["apiPublicLists"] = apiPublicLists

	self.display()
}

func (self *GroupController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	groupName := strings.TrimSpace(self.GetString("groupName"))

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if groupName != "" {
		filters = append(filters, "group_name__icontains", groupName)
	}
	result, count := models.GroupGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["group_name"] = v.GroupName
		row["detail"] = v.Detail
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *GroupController) AjaxSave() {
	Group_id, _ := self.GetInt("id")
	if Group_id == 0 {
		Group := new(models.Group)

		Group.GroupName = strings.TrimSpace(self.GetString("group_name"))
		Group.Detail = strings.TrimSpace(self.GetString("detail"))
		Group.CodeIds = strings.TrimSpace(self.GetString("code_ids"))
		Group.EnvIds = strings.TrimSpace(self.GetString("env_ids"))
		Group.ApiPublicIds = strings.TrimSpace(self.GetString("api_public_ids"))
		Group.CreateId = self.userId
		Group.UpdateId = self.userId
		Group.CreateTime = time.Now().Unix()
		Group.UpdateTime = time.Now().Unix()
		Group.Status = 1

		// 检查登录名是否已经存在
		_, err := models.GroupGetByName(Group.GroupName)

		if err == nil {
			self.ajaxMsg("分组名已经存在", MSG_ERR)
		}

		if _, err := models.GroupAdd(Group); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	GroupUpdate, _ := models.GroupGetById(Group_id)
	// 修改
	GroupUpdate.GroupName = strings.TrimSpace(self.GetString("group_name"))
	GroupUpdate.Detail = strings.TrimSpace(self.GetString("detail"))
	GroupUpdate.CodeIds = strings.TrimSpace(self.GetString("code_ids"))
	GroupUpdate.EnvIds = strings.TrimSpace(self.GetString("env_ids"))
	GroupUpdate.ApiPublicIds = strings.TrimSpace(self.GetString("api_public_ids"))
	GroupUpdate.UpdateId = self.userId
	GroupUpdate.UpdateTime = time.Now().Unix()
	GroupUpdate.Status = 1

	if err := GroupUpdate.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *GroupController) AjaxDel() {

	Group_id, _ := self.GetInt("id")
	Group, _ := models.GroupGetById(Group_id)
	Group.UpdateTime = time.Now().Unix()
	Group.UpdateId = self.userId
	Group.Status = 0
	Group.Id = Group_id

	//TODO 判断是否暂用分组

	if err := Group.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
