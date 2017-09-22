/** kit_admin-v1.0.5 MIT License By http://kit/zhengjinfan.cn e-mail:zheng_jinfan@126.com */
;

layui.define(["jquery"],function(i) {
    var s=layui.jquery,e=(s(document),s("body")),a={v: "1.0.0",times:1,_message:function(){var i=s(".kit-message");
        return i.length>0?i: (e.append('<div class="kit-message"></div>'),s(".kit-message"))
    },show:function(i) {
        var e=this,a=e._message(),t=e.times,n=void 0===(i=i||{
        }).skin?"default":i.skin,d=void 0===i.msg?"请输入一些提示信息!":i.msg,m=void 0===i.autoClose||i.autoClose,u=['<div class="kit-message-item layui-anim layui-anim-upbit" data-times="'+t+'">','<div class="kit-message-body kit-skin-'+n+'">',d,"</div>",'<div class="kit-close kit-skin-'+n+'"><i class="fa fa-times" aria-hidden="true"></i></div>',"</div>"];a.append(u.join(""));var o=a.children("div[data-times="+t+"]").find("i.fa-times");o.off("click").on("click",function() {
            var i=s(this).parents("div.kit-message-item").removeClass("layui-anim-upbit").addClass("layui-anim-fadeout");setTimeout(function(){i.remove()
            },1e3)
        }),m&&setTimeout(function() {
            o.click()
        },3e3),e.times++
    }};

    layui.link("/static/css/message.css"),i("message",a)
});