/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:41
***********************************************/
package controllers

type EnvController struct {
	BaseController
}

func (self *EnvController) List() {
	self.Data["pageTitle"] = "环境配置"
	self.display()
}
