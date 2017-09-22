/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-09 12:53:05
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:54
***********************************************/
package controllers

type GroupController struct {
	BaseController
}

func (self *GroupController) List() {
	self.Data["pageTitle"] = "分组设置"
	self.display()
}
