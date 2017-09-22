/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:41
***********************************************/
package controllers

type CodeController struct {
	BaseController
}

func (self *CodeController) List() {
	self.Data["pageTitle"] = "状态码"
	self.display()
}
