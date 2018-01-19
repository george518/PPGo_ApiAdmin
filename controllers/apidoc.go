/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:41
***********************************************/
package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

type ApiDocController struct {
	BaseController
}

func (self *ApiDocController) Index() {
	self.Data["pageTitle"] = "API文档"
	self.Data["ts"] = time.Now()

	grouplists := groupLists()
	self.Data["grouplists"] = grouplists

	groupId, _ := self.GetInt("id", 1)
	groupInfo, err := models.GroupGetById(groupId)

	if err != nil {
		fmt.Println("数据不存在")
	}

	//公共文档
	apiPublic, err := models.ApiPublicGetByIds(groupInfo.ApiPublicIds)
	self.Data["apiPublic"] = apiPublic

	//环境
	// env, err := models.EnvGetByIds(groupInfo.EnvIds)
	// self.Data["env"] = env

	// //状态码
	// code, err := models.CodeGetByIds(groupInfo.CodeIds)
	// self.Data["code"] = code

	//接口
	apiMenu, _ := models.ApiTreeData(groupId)
	self.Data["apiMenu"] = apiMenu
	self.Data["groupId"] = groupId

	self.TplName = "apidoc/index.html"
}

func (self *ApiDocController) Public() {
	apiPublicId, _ := self.GetInt("id", 1)
	apiPublic, _ := models.ApiPublicGetById(apiPublicId)
	self.Data["apiPublic"] = apiPublic
	self.TplName = "apidoc/apipublic.html"
}

func (self *ApiDocController) Env() {
	groupId, _ := self.GetInt("id", 0)
	groupInfo, _ := models.GroupGetById(groupId)
	env, _ := models.EnvGetByIds(groupInfo.EnvIds)
	self.Data["env"] = env
	self.TplName = "apidoc/env.html"
}

func (self *ApiDocController) Code() {
	groupId, _ := self.GetInt("id", 0)
	groupInfo, _ := models.GroupGetById(groupId)
	code, _ := models.CodeGetByIds(groupInfo.CodeIds)
	self.Data["code"] = code
	self.TplName = "apidoc/code.html"
}

func (self *ApiDocController) ApiDetail() {
	id, _ := self.GetInt("id", 0)
	detail, _ := models.ApiFullDetailById(id)
	row := make(map[string]interface{})
	row["id"] = detail.Id
	row["source_id"] = detail.SourceId
	row["api_url"] = detail.ApiUrl
	row["api_name"] = detail.ApiName
	row["detail"] = detail.Detail
	row["status"] = detail.Status
	row["create_name"] = detail.CreateName
	row["update_name"] = detail.UpdateName
	row["audit_name"] = detail.AuditName
	row["audit_status"] = AUDIT_STATUS[detail.Status]
	row["method"] = REQUEST_METHOD[detail.Method]
	row["audit_time"] = beego.Date(time.Unix(detail.AuditTime, 0), "Y-m-d H:i:s")
	row["update_time"] = beego.Date(time.Unix(detail.UpdateTime, 0), "Y-m-d H:i:s")

	self.Data["pageTitle"] = "查看 " + detail.ApiName
	self.Data["Detail"] = row
	self.TplName = "apidoc/apidetail.html"
}
