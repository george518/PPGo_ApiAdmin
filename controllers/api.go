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
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

var (
	AUDIT_STATUS = [5]string{"<span class='layui-badge layui-bg-black'>暂停使用</span>",
		"<span class='layui-badge layui-bg-orange'>正在开发</span>",
		"<span class='layui-badge layui-bg-blue'>正在审核</span>",
		"<span class='layui-badge layui-bg-green'>审核通过</span>",
		"<span class='layui-badge'>未通过</span>"}
	AUDIT_STATUS_TEXT = [5]string{"暂停使用",
		"正在开发",
		"正在审核",
		"审核通过",
		"未通过"}
	REQUEST_METHOD = [6]string{"未知", "GET", "POST", "PUT", "PATCH", "DELETE"}
)

type ApiController struct {
	BaseController
}

func (self *ApiController) List() {
	self.Data["pageTitle"] = "API接口"
	self.Data["ApiCss"] = true
	self.Data["auditStatus"] = AUDIT_STATUS_TEXT
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

	apiName := strings.TrimSpace(self.GetString("apiName"))
	status, _ := self.GetInt("status", -1)

	//获取分组
	sourceList := sourceLists()

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	if status != -1 {
		filters = append(filters, "status", status)
	} else {
		filters = append(filters, "status__in", []int{0, 1, 2, 3, 4})
	}

	if apiName != "" {
		filters = append(filters, "api_name__icontains", apiName)
	}
	result, count := models.ApiDetailGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["api_name"] = v.ApiName
		row["api_url"] = v.ApiUrl
		row["status_text"] = AUDIT_STATUS[v.Status]
		row["status"] = v.Status
		row["method"] = REQUEST_METHOD[v.Method]
		sourceInfo := getSourceInfo(sourceList, v.SourceId)
		row["source_name"] = sourceInfo.SourceName
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

//查看详情
func (self *ApiController) Detail() {
	self.Data["ApiCss"] = true
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
	self.display()
}

//审核
func (self *ApiController) Audit() {
	self.Data["ApiCss"] = true
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

	self.Data["pageTitle"] = "审核 " + detail.ApiName
	self.Data["Detail"] = row
	self.display()
}

func (self *ApiController) AjaxDel() {

	Api_id, _ := self.GetInt("id")
	Api, _ := models.ApiDetailGetById(Api_id)
	Api.UpdateTime = time.Now().Unix()
	Api.UpdateId = self.userId

	if Api.Status == 0 {
		Api.Status = 1
	} else {
		Api.Status = 0
	}

	Api.Id = Api_id

	//TODO 判断是否暂用API

	if err := Api.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *ApiController) AjaxChangeStatus() {
	Api_ids := strings.TrimSpace(self.GetString("ids"))
	status, _ := self.GetInt("status")

	if status == 1 {
		status = 2
	}
	_, err := models.ApiChangeStatus(Api_ids, status)

	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

//新增接口实例
func (self *ApiController) Add() {
	self.Data["pageTitle"] = "添加接口"
	// //分组
	// groupList := groupLists()
	// self.Data["groupList"] = groupList
	//资源
	sourceList := sourceLists()
	self.Data["sourceList"] = sourceList
	fmt.Println(sourceList)
	tmplates := templateLists()
	self.Data["templates"] = tmplates

	self.display()
}

//修改接口实例
func (self *ApiController) Edit() {
	self.Data["ApiCss"] = true
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
	row["method"] = detail.Method
	row["audit_time"] = beego.Date(time.Unix(detail.AuditTime, 0), "Y-m-d H:i:s")
	row["update_time"] = beego.Date(time.Unix(detail.UpdateTime, 0), "Y-m-d H:i:s")

	self.Data["pageTitle"] = "查看 " + detail.ApiName
	self.Data["Detail"] = row

	sourceList := sourceLists()
	self.Data["sourceList"] = sourceList

	tmplates := templateLists()
	self.Data["templates"] = tmplates

	self.display()
}

func (self *ApiController) AjaxSave() {

	Api_id, _ := self.GetInt("id")
	if Api_id == 0 {
		ApiDetail := new(models.ApiDetail)
		ApiDetail.SourceId, _ = self.GetInt("source_id")
		ApiDetail.Method, _ = self.GetInt("method")
		ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
		ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
		ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))
		ApiDetail.CreateId = self.userId
		ApiDetail.UpdateId = self.userId
		ApiDetail.CreateTime = time.Now().Unix()
		ApiDetail.UpdateTime = time.Now().Unix()
		ApiDetail.Status = 1
		_, err := models.ApiDetailAdd(ApiDetail)
		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}
	//修改
	ApiDetail, _ := models.ApiDetailGetById(Api_id)
	ApiDetail.SourceId, _ = self.GetInt("source_id")
	ApiDetail.Id, _ = self.GetInt("id")
	ApiDetail.Method, _ = self.GetInt("method")
	ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
	ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
	ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))

	ApiDetail.UpdateId = self.userId
	ApiDetail.UpdateTime = time.Now().Unix()
	ApiDetail.Status, _ = self.GetInt("status")

	ApiDetail.Status = 1
	if err := ApiDetail.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
