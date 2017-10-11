PPGo_ApiAdmin
====
什么东西？What?
----
API管理工具 golang开发，基于beego，页面基于layUi
练手使用，尚未用于生产环境

有什么价值？
----
1、RBAC权限完善，多角色管理系统
2、后台界面完整，多标签页面
3、API相关页面有比较复杂的使用案例
所以，可以作为一个基础框架使用，快速开发。初学者还可以作为熟悉beego使用。

效果展示
----
登录界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/login.png?raw=true "github")
<br/><br/>
权限界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/auth.png?raw=true "github")
角色界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/role.png?raw=true "github")
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/role_add.png?raw=true "github")
管理员界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/admin.png?raw=true "github")
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/admin_add.png?raw=true "github")

API界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/api.png?raw=true "github")
<br/><br/>

安装方法    
----
1、go get github.com/george518/PPGo_ApiAdmin 
2、创建mysql数据库，并将ppgo_api_admin.sql导入    
3、修改config 配置数据库    
4、运行 go build    
5、运行 ./run.sh start|stop


前台访问：http://your_host:8081
用户名：admin 密码：george518


联系我
----
qq:41352963


