/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:41
***********************************************/
package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

var (
	AUDIT_STATUS   = [4]string{"暂停使用", "正在开发", "正在审核", "审核通过"}
	PROTOCOL_TYPE  = [3]string{"HTTP/HTTPS", "HTTP", "HTTPS"}
	REQUEST_METHOD = [6]string{"未知", "GET", "POST", "PUT", "PATCH", "DELETE"}
)

type ApiController struct {
	BaseController
}

func (self *ApiController) List() {
	self.Data["pageTitle"] = "API接口"
	self.Data["ApiCss"] = true
	group_id, _ := self.GetInt("gid", 0)
	// 分组
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	result, _ := models.GroupGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = int(v.Id)
		row["group_name"] = v.GroupName
		list[k] = row
	}
	self.Data["Groups"] = list

	//所有接口
	filters_source := make([]interface{}, 0)
	filters_source = append(filters_source, "status__in", []int{1, 2, 3})
	if group_id != 0 {
		filters_source = append(filters_source, "group_id", group_id)
	}
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
	self.Data["Gid"] = group_id
	self.display()
}

//显示所有的接口详情
func (self *ApiController) Show() {
	self.Data["ApiCss"] = true
	id, _ := self.GetInt("id", 0)
	detail, _ := models.ApiDetailsGetById(id)
	list := make([]map[string]interface{}, len(detail))

	for k, v := range detail {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["source_id"] = v.SourceId
		row["api_url"] = v.ApiUrl
		row["api_name"] = v.ApiName
		row["example"] = v.Example
		row["detail"] = v.Detail
		row["result"] = v.Result
		row["status"] = v.Status
		row["create_name"] = v.CreateName
		row["update_name"] = v.UpdateName
		row["audit_name"] = v.AuditName
		row["protocol_type"] = PROTOCOL_TYPE[v.ProtocolType]
		row["audit_status"] = AUDIT_STATUS[v.Status]
		row["method"] = REQUEST_METHOD[v.Method]
		row["audit_time"] = beego.Date(time.Unix(v.AuditTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		//参数
		row["Params"], _ = models.ApiParamGetById(v.Id)
		list[k] = row
	}

	self.Data["Detail"] = list

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

// func (self *ApiController) Table() {
// 	//列表
// 	page, err := self.GetInt("page")
// 	if err != nil {
// 		page = 1
// 	}
// 	limit, err := self.GetInt("limit")
// 	if err != nil {
// 		limit = 30
// 	}

// 	self.pageSize = limit
// 	//查询条件
// 	filters := make([]interface{}, 0)
// 	filters = append(filters, "status", 1)
// 	result, count := models.ApiGetList(page, self.pageSize, filters...)
// 	list := make([]map[string]interface{}, len(result))
// 	for k, v := range result {
// 		row := make(map[string]interface{})
// 		row["id"] = v.Id
// 		row["source_name"] = v.SourceName
// 		row["group_id"] = v.GroupId
// 		list[k] = row
// 	}
// 	self.ajaxList("成功", MSG_OK, count, list)
// }

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

//修改接口实例
func (self *ApiController) EditApi() {
	id, _ := self.GetInt("id", 0)
	detail, _ := models.ApiDetailGetById(id)
	params, _ := models.ApiParamGetById(detail.Id)
	self.Data["Detail"] = detail
	self.Data["Params"] = params
	self.Data["ParamsCount"] = len(params)
	self.display()
}

func (self *ApiController) AjaxApiSave() {

	Api_id, _ := self.GetInt("id")
	if Api_id == 0 {
		ApiDetail := new(models.ApiDetail)
		ApiDetail.SourceId, _ = self.GetInt("source_id")
		ApiDetail.ProtocolType, _ = self.GetInt("protocol_type")
		ApiDetail.Method, _ = self.GetInt("method")
		ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
		ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
		ApiDetail.Result = strings.TrimSpace(self.GetString("result"))
		ApiDetail.Example = strings.TrimSpace(self.GetString("example"))
		ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))
		ApiDetail.CreateId = self.userId
		ApiDetail.UpdateId = self.userId
		ApiDetail.CreateTime = time.Now().Unix()
		ApiDetail.UpdateTime = time.Now().Unix()
		ApiDetail.Status = 1
		detail_id, err := models.ApiDetailAdd(ApiDetail)
		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}

		params := make(map[int]map[string]string)
		for k, v := range self.Ctx.Request.Form {
			if strings.Contains(k, "attr_") == true {
				ks := strings.Split(k, "_")
				i, _ := strconv.Atoi(ks[1])
				if _, ok := params[i]; ok {
					params[i][ks[2]] = v[0]
				} else {
					param := make(map[string]string)
					param[ks[2]] = v[0]
					params[i] = param
				}
			}
		}
		for _, vv := range params {
			apiParam := new(models.ApiParam)
			if vv["key"] == "" {
				break
			}
			apiParam.ApiKey = vv["key"]
			apiParam.ApiType = vv["type"]
			apiParam.ApiValue = vv["value"]
			apiParam.ApiDetail = vv["detail"]
			apiParam.IsNull = vv["isnull"]
			apiParam.DetailId = detail_id
			apiParam.CreateId = self.userId
			apiParam.UpdateId = self.userId
			apiParam.CreateTime = time.Now().Unix()
			apiParam.UpdateTime = time.Now().Unix()
			apiParam.Status = 1

			if _, err := models.ApiParamAdd(apiParam); err != nil {
				self.ajaxMsg(err.Error(), MSG_ERR)
			}
		}
		self.ajaxMsg("", MSG_OK)
	}
	//修改
	ApiDetail, _ := models.ApiDetailGetById(Api_id)
	ApiDetail.SourceId, _ = self.GetInt("source_id")
	ApiDetail.Id, _ = self.GetInt("id")
	ApiDetail.ProtocolType, _ = self.GetInt("protocol_type")
	ApiDetail.Method, _ = self.GetInt("method")
	ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
	ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
	ApiDetail.Result = strings.TrimSpace(self.GetString("result"))
	ApiDetail.Example = strings.TrimSpace(self.GetString("example"))
	ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))

	ApiDetail.UpdateId = self.userId
	ApiDetail.UpdateTime = time.Now().Unix()
	ApiDetail.Status, _ = self.GetInt("status")

	if err := ApiDetail.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	//先删除后新增参数
	apiParamDel := new(models.ApiParam)

	if _, err := apiParamDel.Delete(int64(Api_id), self.userId); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	params := make(map[int]map[string]string)
	for k, v := range self.Ctx.Request.Form {
		if strings.Contains(k, "attr_") == true {
			ks := strings.Split(k, "_")
			i, _ := strconv.Atoi(ks[1])
			if _, ok := params[i]; ok {
				params[i][ks[2]] = v[0]
			} else {
				param := make(map[string]string)
				param[ks[2]] = v[0]
				params[i] = param
			}
		}
	}
	for _, vv := range params {
		apiParam := new(models.ApiParam)
		apiParam.ApiKey = vv["key"]
		apiParam.ApiType = vv["type"]
		apiParam.ApiValue = vv["value"]
		apiParam.ApiDetail = vv["detail"]
		apiParam.IsNull = vv["isnull"]
		apiParam.DetailId = int64(Api_id)
		apiParam.CreateId = self.userId
		apiParam.UpdateId = self.userId
		apiParam.CreateTime = time.Now().Unix()
		apiParam.UpdateTime = time.Now().Unix()
		apiParam.Status = 1

		if _, err := models.ApiParamAdd(apiParam); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)

	// ApiUpdate, _ := models.ApiGetById(Api_id)
	// // 修改
	// ApiUpdate.SourceName = strings.TrimSpace(self.GetString("source_name"))
	// ApiUpdate.GroupId, _ = self.GetInt("group_id")
	// ApiUpdate.UpdateId = self.userId
	// ApiUpdate.UpdateTime = time.Now().Unix()
	// ApiUpdate.Status = 2

	// if err := ApiUpdate.Update(); err != nil {
	// 	self.ajaxMsg(err.Error(), MSG_ERR)
	// }
	self.ajaxMsg("", MSG_OK)
}
