/** kit_admin-v1.0.5 MIT License By http://kit/zhengjinfan.cn e-mail:zheng_jinfan@126.com */
;
layui.define(["layer", "laytpl", "element"],
    function(i) {
        var e = layui.jquery,
            a = layui.layer,
            t = (e(window), e(document)),
            n = layui.laytpl,
            l = layui.element;
        i("navbar", {
            v: "1.0.2",
            config: {
                data: void 0,
                remote: {
                    url: void 0,
                    type: "GET",
                    jsonp: !1
                },
                cached: !1,
                elem: void 0,
                filter: "kitNavbar"
            },
            set: function(i) {
                var a = this;
                return a.config.data = void 0,
                    e.extend(!0, a.config, i),
                    a
            },
            hasElem: function() {
                var i = this.config;
                return void 0 !== i.elem || 0 !== t.find("ul[kit-navbar]").length || !e(i.elem) || (layui.hint().error("Navbar error:请配置Navbar容器."), !1)
            },
            getElem: function() {
                var i = this.config;
                return void 0 !== i.elem && e(i.elem).length > 0 ? e(i.elem) : t.find("ul[kit-navbar]")
            },
            bind: function(i) {
                var n = this;
                n.config;
                return n.hasElem() ? (n.getElem().find("a[kit-target]").each(function() {
                    var t = e(this),
                        n = void 0;
                    t.hover(function() {
                            n = a.tips(e(this).children("span").text(), this)
                        },
                        function() {
                            n && a.close(n)
                        }),
                        t.off("click").on("click",
                            function() {
                                var e, a = t.data("options");
                                if (void 0 !== a) try {
                                    e = new Function("return " + a)()
                                } catch(i) {
                                    layui.hint().error("Navbar 组件a[data-options]配置项存在语法错误：" + a)
                                } else e = {
                                    icon: t.data("icon"),
                                    id: t.data("id"),
                                    title: t.data("title"),
                                    url: t.data("url")
                                };
                                "function" == typeof i && i(e)
                            })
                }), e(".kit-side-fold").off("click").on("click",
                    function() {
                        var i = t.find("div.kit-side");
                        i.hasClass("kit-sided") ? (i.removeClass("kit-sided"), t.find("div.layui-body").removeClass("kit-body-folded"), t.find("div.layui-footer").removeClass("kit-footer-folded")) : (i.addClass("kit-sided"), t.find("div.layui-body").addClass("kit-body-folded"), t.find("div.layui-footer").addClass("kit-footer-folded"))
                    }), n) : n
            },
            render: function(i) {
                var t = this,
                    d = t.config,
                    o = d.remote,
                    r = ["{{# layui.each(d,function(index, item){ }}", "{{# if(item.spread){ }}", '<li class="layui-nav-item layui-nav-itemed">', "{{# }else{ }}", '<li class="layui-nav-item">', "{{# } }}", "{{# var hasChildren = item.children!==undefined && item.children.length>0; }}", "{{# if(hasChildren){ }}", '<a href="javascript:;">', '{{# if (item.icon.indexOf("fa-") !== -1) { }}', '<i class="fa {{item.icon}}" aria-hidden="true"></i>', "{{# } else { }}", '<i class="layui-icon">{{item.icon}}</i>', "{{# } }}", "<span> {{item.title}}</span>", "</a>", "{{# var children = item.children; }}", '<dl class="layui-nav-child">', "{{# layui.each(children,function(childIndex, child){ }}", "<dd>", "<a href=\"javascript:;\" kit-target data-options=\"{url:'{{child.url}}',icon:'{{child.icon}}',title:'{{child.title}}',id:'{{child.id}}'}\">", '{{# if (child.icon.indexOf("fa-") !== -1) { }}', '<i class="fa {{child.icon}}" aria-hidden="true"></i>', "{{# } else { }}", '<i class="layui-icon">{{child.icon}}</i>', "{{# } }}", "<span> {{child.title}}</span>", "</a>", "</dd>", "{{# }); }}", "</dl>", "{{# }else{ }}", "<a href=\"javascript:;\" kit-target data-options=\"{url:'{{item.url}}',icon:'{{item.icon}}',title:'{{item.title}}',id:'{{item.id}}'}\">", '{{# if (item.icon.indexOf("fa-") !== -1) { }}', '<i class="fa {{item.icon}}" aria-hidden="true"></i>', "{{# } else { }}", '<i class="layui-icon">{{item.icon}}</i>', "{{# } }}", "<span> {{item.title}}</span>", "</a>", "{{# } }}", "</li>", "{{# }); }}"],
                    c = [],
                    s = a.load(2);
                if (!t.hasElem()) return t;
                var f = t.getElem();
                if (void 0 !== d.data && d.data.length > 0) c = d.data;
                else {
                    o.jsonp;
                    var u = {
                        url: o.url,
                        type: o.type,
                        error: function(i, e, a) {
                            layui.hint().error("Navbar error:AJAX请求出错." + a)
                        },
                        success: function(i) {
                            c = i
                        }
                    };
                    e.extend(!0, u, o.jsonp ? {
                        dataType: "jsonp",
                        jsonp: "callback",
                        jsonpCallback: "jsonpCallback"
                    }: {
                        dataType: "json"
                    }),
                        e.support.cors = !0,
                        e.ajax(u)
                }
                var h = setInterval(function() {
                        c.length > 0 && (clearInterval(h), n(r.join("")).render(c,
                            function(e) {
                                f.html(e),
                                    l.init(),
                                    t.bind(function(e) {
                                        "function" == typeof i && i(e)
                                    }),
                                s && a.close(s)
                            }))
                    },
                    50);
                return t
            }
        })
    });