/**********************************************
** @Des: base controller
** @Author: haodaquan
** @Date:   2017-09-07 16:54:40
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-18 10:28:01
***********************************************/
package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/libs"
	"github.com/george518/PPGo_ApiAdmin/models"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.Admin
	userId         int
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
}

//前期准备
func (self *BaseController) Prepare() {
	self.pageSize = 20
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["version"] = beego.AppConfig.String("version")
	self.Data["siteName"] = beego.AppConfig.String("site.name")
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName
	// noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start"
	// isNoAuth := strings.Contains(noAuth, self.actionName)
	fmt.Println(self.controllerName)
	if (strings.Compare(self.controllerName, "apidoc")) != 0 {
		self.auth()
	}

	self.Data["loginUserId"] = self.userId
	self.Data["loginUserName"] = self.userName
}

//登录权限验证
func (self *BaseController) auth() {

	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
	self.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.AdminGetById(userId)
			if err == nil && password == libs.Md5([]byte(self.getClientIp()+"|"+user.Password+user.Salt)) {
				self.userId = user.Id
				self.loginName = user.LoginName
				self.userName = user.RealName
				self.user = user
				self.AdminAuth()
			}

			isHasAuth := strings.Contains(self.allowUrl, self.controllerName+"/"+self.actionName)
			noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start/show/ajaxapisave/index/group/public/env/code/apidetail"
			isNoAuth := strings.Contains(noAuth, self.actionName)
			if isHasAuth == false && isNoAuth == false {
				self.Ctx.WriteString("没有权限")
				self.ajaxMsg("没有权限", MSG_ERR)
				return
			}
		}
	}

	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "loginin") {
		self.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

func (self *BaseController) AdminAuth() {
	// 左侧导航栏
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if self.userId != 1 {
		//普通管理员
		adminAuthIds, _ := models.RoleAuthGetByIds(self.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	list2 := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _, v := range result {
		if v.AuthUrl != " " || v.AuthUrl != "/" {
			allow_url += v.AuthUrl
		}
		row := make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2[j] = row
			j++
		}
	}

	self.Data["SideMenu1"] = list[:i]  //一级菜单
	self.Data["SideMenu2"] = list2[:j] //二级菜单

	self.allowUrl = allow_url + "/home/index"
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := self.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//加载模板
func (self *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}

//ajax返回
func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//ajax返回 列表
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//分组公共方法
type groupList struct {
	Id        int
	GroupName string
}

func groupLists() (gl []groupList) {
	groupFilters := make([]interface{}, 0)
	groupFilters = append(groupFilters, "status", 1)
	groupResult, _ := models.GroupGetList(1, 1000, groupFilters...)
	for _, gv := range groupResult {
		groupRow := groupList{}
		groupRow.Id = int(gv.Id)
		groupRow.GroupName = gv.GroupName
		gl = append(gl, groupRow)
	}
	return gl
}

//获取单个分组信息
func getGroupInfo(gl []groupList, groupId int) (groupInfo groupList) {
	for _, v := range gl {
		if v.Id == groupId {
			groupInfo = v
		}
	}
	return
}

type sourceList struct {
	Id         int
	SourceName string
	GroupId    int
	GroupName  string
}

func sourceLists() (sl []sourceList) {

	grouplists := groupLists()
	var groupinfo groupList
	sourceFilters := make([]interface{}, 0)
	sourceFilters = append(sourceFilters, "status", 1)
	sourceResult, _ := models.ApiSourceGetList(1, 1000, sourceFilters...)
	for _, sv := range sourceResult {
		sourceRow := sourceList{}
		sourceRow.Id = int(sv.Id)
		sourceRow.GroupId = sv.GroupId
		groupinfo = getGroupInfo(grouplists, sv.GroupId)
		sourceRow.GroupName = groupinfo.GroupName
		sourceRow.SourceName = sv.SourceName
		sl = append(sl, sourceRow)
	}
	return sl
}

func getSourceInfo(gl []sourceList, sourceId int) (sourceInfo sourceList) {
	for _, v := range gl {
		if v.Id == sourceId {
			sourceInfo = v
		}
	}
	return
}

type envList struct {
	Id      int
	EnvName string
	EnvHost string
}

func envLists() (sl []envList) {
	envFilters := make([]interface{}, 0)
	envFilters = append(envFilters, "status__in", 1)
	envResult, _ := models.EnvGetList(1, 1000, envFilters...)
	for _, sv := range envResult {
		envRow := envList{}
		envRow.Id = int(sv.Id)
		envRow.EnvName = sv.EnvName
		envRow.EnvHost = sv.EnvHost
		sl = append(sl, envRow)
	}
	return sl
}

type templateList struct {
	Id           int
	TemplateName string
	Detail       string
}

func templateLists() (sl []templateList) {
	templateFilters := make([]interface{}, 0)
	templateFilters = append(templateFilters, "status", 1)
	templateResult, _ := models.TemplateGetList(1, 1000, templateFilters...)
	for _, sv := range templateResult {
		templateRow := templateList{}
		templateRow.Id = int(sv.Id)
		templateRow.TemplateName = sv.TemplateName
		templateRow.Detail = sv.Detail
		sl = append(sl, templateRow)
	}
	return sl
}

type codeList struct {
	Id     int
	Code   string
	Desc   string
	Detail string
}

func codeLists() (sl []codeList) {
	codeFilters := make([]interface{}, 0)
	codeFilters = append(codeFilters, "status", 1)
	codeResult, _ := models.CodeGetList(1, 1000, codeFilters...)
	for _, sv := range codeResult {
		codeRow := codeList{}
		codeRow.Id = int(sv.Id)
		codeRow.Code = sv.Code
		codeRow.Desc = sv.Desc
		codeRow.Detail = sv.Detail
		sl = append(sl, codeRow)
	}
	return sl
}

type apiPublicList struct {
	Id            int
	ApiPublicName string
	Sort          int
}

func apiPublicLists() (sl []apiPublicList) {
	apiPublicFilters := make([]interface{}, 0)
	apiPublicFilters = append(apiPublicFilters, "status", 1)
	apiPublicResult, _ := models.ApiPublicGetList(1, 1000, apiPublicFilters...)
	for _, sv := range apiPublicResult {
		apiPublicRow := apiPublicList{}
		apiPublicRow.Id = int(sv.Id)
		apiPublicRow.ApiPublicName = sv.ApiPublicName
		apiPublicRow.Sort = sv.Sort
		sl = append(sl, apiPublicRow)
	}
	return sl
}
