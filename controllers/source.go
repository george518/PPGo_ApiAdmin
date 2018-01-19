/**********************************************
** @Des: 资源设置
** @Author: haodaquan
** @Date:   2017-09-08 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:41
***********************************************/
package controllers

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

type ApiSourceController struct {
	BaseController
}

func (self *ApiSourceController) List() {
	self.Data["pageTitle"] = "API资源"
	self.display()
}

func (self *ApiSourceController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	sourceName := strings.TrimSpace(self.GetString("sourceName"))

	//获取分组
	groupList := groupLists()

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if sourceName != "" {
		filters = append(filters, "source_name__icontains", sourceName)
	}
	result, count := models.ApiSourceGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["source_name"] = v.SourceName
		row["group_id"] = v.GroupId
		groupInfo := getGroupInfo(groupList, v.GroupId)
		row["group_name"] = groupInfo.GroupName
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

// //显示所有的接口详情
// func (self *ApiController) Show() {
// 	self.Data["ApiCss"] = true
// 	id, _ := self.GetInt("id", 0)
// 	sourceName := strings.TrimSpace(self.GetString("sourceName"))
// 	self.Data["pageTitle"] = sourceName

// 	detail, _ := models.ApiDetailsGetById(id)
// 	list := make([]map[string]interface{}, len(detail))

// 	for k, v := range detail {
// 		row := make(map[string]interface{})
// 		row["id"] = v.Id
// 		row["source_id"] = v.SourceId
// 		row["api_url"] = v.ApiUrl
// 		row["api_name"] = v.ApiName
// 		row["detail"] = v.Detail
// 		row["status"] = v.Status
// 		row["create_name"] = v.CreateName
// 		row["update_name"] = v.UpdateName
// 		row["audit_name"] = v.AuditName
// 		row["audit_status"] = AUDIT_STATUS[v.Status]
// 		row["method"] = REQUEST_METHOD[v.Method]
// 		row["audit_time"] = beego.Date(time.Unix(v.AuditTime, 0), "Y-m-d H:i:s")
// 		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
// 		//参数
// 		row["Params"], _ = models.ApiParamGetById(v.Id)
// 		list[k] = row
// 	}

// 	self.Data["Detail"] = list
// 	self.Data["sid"] = id
// 	self.display("api/info")
// }

func (self *ApiSourceController) Add() {
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

func (self *ApiSourceController) Edit() {
	self.Data["pageTitle"] = "编辑API"

	id, _ := self.GetInt("id", 0)
	Api, err := models.ApiSourceGetById(id)
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

//存储资源
func (self *ApiSourceController) AjaxSave() {
	Api_id, _ := self.GetInt("id")
	if Api_id == 0 {
		Api := new(models.ApiSource)

		Api.SourceName = strings.TrimSpace(self.GetString("source_name"))
		Api.GroupId, _ = self.GetInt("group_id")
		Api.CreateId = self.userId
		Api.UpdateId = self.userId
		Api.CreateTime = time.Now().Unix()
		Api.UpdateTime = time.Now().Unix()
		Api.Status = 1

		// 检查登录名是否已经存在
		_, err := models.ApiSourceGetByName(Api.SourceName)

		if err == nil {
			self.ajaxMsg("资源名已经存在", MSG_ERR)
		}

		if _, err := models.ApiSourceAdd(Api); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	ApiUpdate, _ := models.ApiSourceGetById(Api_id)
	// 修改
	ApiUpdate.SourceName = strings.TrimSpace(self.GetString("source_name"))
	ApiUpdate.GroupId, _ = self.GetInt("group_id")
	ApiUpdate.UpdateId = self.userId
	ApiUpdate.UpdateTime = time.Now().Unix()
	ApiUpdate.Status = 1

	if err := ApiUpdate.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *ApiSourceController) AjaxDel() {
	Api_id, _ := self.GetInt("id")
	Api, _ := models.ApiSourceGetById(Api_id)
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
// func (self *ApiController) AddApi() {
// 	self.Data["pageTitle"] = "添加接口"
// 	source_id, _ := self.GetInt("sid")
// 	self.Data["Sid"] = source_id

// 	sourceName := strings.TrimSpace(self.GetString("sourceName"))
// 	self.Data["pageTitle"] = sourceName + " > " + "添加接口"
// 	self.Data["sourceName"] = sourceName

// 	//查询条件
// 	self.display()
// }

// //修改接口实例
// func (self *ApiController) EditApi() {
// 	id, _ := self.GetInt("id", 0)
// 	detail, _ := models.ApiDetailGetById(id)
// 	params, _ := models.ApiParamGetById(detail.Id)
// 	self.Data["Detail"] = detail
// 	self.Data["Params"] = params
// 	self.Data["ParamsCount"] = len(params)
// 	self.display()
// }

// func (self *ApiController) AjaxApiSave() {

// 	Api_id, _ := self.GetInt("id")
// 	if Api_id == 0 {
// 		ApiDetail := new(models.ApiDetail)
// 		ApiDetail.SourceId, _ = self.GetInt("source_id")
// 		ApiDetail.Method, _ = self.GetInt("method")
// 		ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
// 		ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
// 		ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))
// 		ApiDetail.CreateId = self.userId
// 		ApiDetail.UpdateId = self.userId
// 		ApiDetail.CreateTime = time.Now().Unix()
// 		ApiDetail.UpdateTime = time.Now().Unix()
// 		ApiDetail.Status = 1
// 		_, err := models.ApiDetailAdd(ApiDetail)
// 		if err != nil {
// 			self.ajaxMsg(err.Error(), MSG_ERR)
// 		}
// 		self.ajaxMsg("", MSG_OK)
// 	}
// 	//修改
// 	ApiDetail, _ := models.ApiDetailGetById(Api_id)
// 	ApiDetail.SourceId, _ = self.GetInt("source_id")
// 	ApiDetail.Id, _ = self.GetInt("id")
// 	ApiDetail.Method, _ = self.GetInt("method")
// 	ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
// 	ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
// 	ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))

// 	ApiDetail.UpdateId = self.userId
// 	ApiDetail.UpdateTime = time.Now().Unix()
// 	ApiDetail.Status, _ = self.GetInt("status")

// 	if err := ApiDetail.Update(); err != nil {
// 		self.ajaxMsg(err.Error(), MSG_ERR)
// 	}
// 	self.ajaxMsg("", MSG_OK)
// }
