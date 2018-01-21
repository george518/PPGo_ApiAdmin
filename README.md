PPGo_ApiAdmin
====
什么东西？What?
----
API管理工具 golang开发，基于beego，页面基于layUi,目前版本1.0.0   
demo地址：http://apiadmin.haodaquan.com
后台用户名和密码请加群联系获取（有人删数据，没办法）


有什么价值？
----
1、RBAC权限完善，多角色管理系统    
2、后台界面完整，多标签页面    
3、API相关页面有比较复杂的使用案例    
所以，可以作为一个基础框架使用，快速开发。初学者还可以作为熟悉beego使用。 
4、可用于中小团队API管理使用  

用到了哪些？
----
1、界面框架layUI2.4.5    
2、makedown.md   
3、beego1.8
4、Ztree   


效果展示
----
前端界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/index.png?raw=true "github")
<br/><br/>

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

API列表界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/api_list.png?raw=true "github")

API新增界面<br/>
![github](https://github.com/george518/PPGo_ApiAdmin/blob/master/static/img/demoImage/add_api.png?raw=true "github")

<br/><br/>



安装方法    
----
1、go get github.com/george518/PPGo_ApiAdmin    
2、创建mysql数据库，并将ppgo_api_admin.sql导入    
3、修改config 配置数据库    
4、运行 go build    
5、运行 ./run.sh start|stop


前台访问：http://your_host:8081
后台访问：http://your_host:8081/login
用户名：admin 密码：george518    

联系我
----
qq群号:547564773
欢迎交流，欢迎提交代码。


