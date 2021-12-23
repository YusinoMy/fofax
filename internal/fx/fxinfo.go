package fx

var defalutPlugins = []Plugin{

	{
		Id:          "fx-2021-1001",
		Query:       "google-reverse",
		RuleName:    "Google反代服务器",
		RuleEnglish: "Google Reverse proxy",
		Description: "不用挂代理就可以访问的Google搜索，但搜索记录可能会被记录。",
		Author:      "fofa",
		FofaQuery:   `body="var c = Array.prototype.slice.call(arguments, 1);return function() {var d=c.slice();"`,
		Tag:         []string{"google"},
		Type:        TypeInline,
		Source:      "https://tp.wjx.top/m/67114073.aspx",
	},
	{
		Id:          "fx-2021-1002",
		Query:       "python-simplehttp",
		RuleName:    "Python SimpleHTTP",
		RuleEnglish: "Python SimpleHTTP Server",
		Description: "Python SimpleHTTP临时服务器",
		Author:      "fofa",
		FofaQuery:   `(server="SimpleHTTP/0.6 Python/3.6.3" || server="SimpleHTTP/0.6 Python/3.6.8" || server="SimpleHTTP/0.6 Python/3.7.0" || server="SimpleHTTP/0.6 Python/3.7.1" || server="SimpleHTTP/0.6 Python/3.7.2" || server="SimpleHTTP/0.6 Python/3.7.3" || server="SimpleHTTP/0.6 Python/3.7.4" || server="SimpleHTTP/0.6 Python/3.7.5" || server="SimpleHTTP/0.6 Python/3.7.6" || server="SimpleHTTP/0.6 Python/3.7.7" || server="SimpleHTTP/0.6 Python/3.7.8" || server="SimpleHTTP/0.6 Python/3.7.9" || server="SimpleHTTP/0.6 Python/3.8.1" || server="SimpleHTTP/0.6 Python/3.8.2" || server="SimpleHTTP/0.6 Python/3.8.3" || server="SimpleHTTP/0.6 Python/3.8.4" || server="SimpleHTTP/0.6 Python/3.8.5" || server="SimpleHTTP/0.6 Python/3.8.6" || server="SimpleHTTP/0.6 Python/3.8.7" || server="SimpleHTTP/0.6 Python/3.8.8" || server="SimpleHTTP/0.6 Python/3.8.9" || server="SimpleHTTP/0.6 Python/3.9.1" || server="SimpleHTTP/0.6 Python/3.9.2" || server="SimpleHTTP/0.6 Python/3.9.3" || server="SimpleHTTP/0.6 Python/3.9.4" || server="SimpleHTTP/0.6 Python/3.9.5" || server="SimpleHTTP/0.6 Python/3.9.6" || server="SimpleHTTP/0.6 Python/3.9.7" || server="SimpleHTTP/0.6 Python/3.9.8" || server="SimpleHTTP/0.6 Python/3.9.9") && title="Directory listing for"`,
		Tag:         []string{"python"},
		Type:        TypeInline,
		Source:      "https://tp.wjx.top/m/67114073.aspx",
	},
	{
		Id:          "fx-2021-1003",
		Query:       "data-leak",
		RuleName:    "社工库",
		RuleEnglish: "data leak",
		Description: "社工库",
		Author:      "fofa",
		FofaQuery:   `title="社工库" || ((title="社工库" && title="系统") || (title="社工库查询" ))`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1004",
		Query:       "hfs-rce",
		RuleName:    "存在命令执行的HFS服务",
		RuleEnglish: "Presence of HFS services for command execution",
		Description: "这个语法可搜索存在命令执行的HFS服务，使用者多数为抓鸡黑客，可以直接上线能捡到不少有趣的东西~",
		Author:      "fofa",
		FofaQuery:   `body="HttpFileServer v2.3 beta 287"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1005",
		Query:       "satellite-ftp",
		RuleName:    "一键日卫星FTP？",
		RuleEnglish: "Satellites FTP",
		Description: "一键日卫星FTP？",
		Author:      "fofa",
		FofaQuery:   `banner="Cobham SATCOM"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1006",
		Query:       "mk-mining",
		RuleName:    "mk路由器全球挖矿感染",
		RuleEnglish: "mk router global mining infection",
		Description: "mk路由器全球挖矿感染",
		Author:      "fofa",
		FofaQuery:   `app="Mikrotik-HttpProxy"&&(body="CoinHive.Anonymous" || body="CRLT.Anonymous" || body=" WMP.Anonymous(")`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1007",
		Query:       "ss-manager-login",
		RuleName:    "ss-Manager 登录",
		RuleEnglish: "ss-Manager Login",
		Description: "Shadowsocks-Manager",
		Author:      "fofa",
		FofaQuery:   `body="indeterminate" && body="MainController" && header="X-Powered-By: Express"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1008",
		Query:       "heating-monitor",
		RuleName:    "供暖监控系统",
		RuleEnglish: "Heating monitoring systems",
		Description: "供暖监控系统",
		Author:      "fofa",
		FofaQuery:   `body="s1v13.htm"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1009",
		Query:       "free-proxy",
		RuleName:    "免费代理池",
		RuleEnglish: "Free Proxy Pool",
		Description: "获取免费代理池",
		Author:      "fofa",
		FofaQuery:   `body="get all proxy from proxy pool"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1010",
		Query:       "honeypot",
		RuleName:    "蜜罐",
		RuleEnglish: "honeypot",
		Description: "蜜罐",
		Author:      "fofa",
		FofaQuery:   `(header="uc-httpd 1.0.0" && server="JBoss-5.0") || server="Apache,Tomcat,Jboss,weblogic,phpstudy,struts"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1011",
		Query:       "hacked-website",
		RuleName:    "被挂黑的站点",
		RuleEnglish: "hacked by xxx",
		Description: "被挂黑的站点",
		Author:      "fofa",
		FofaQuery:   `body="hacked by"`,
		Tag:         []string{"fun"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-1012",
		Query:       "jupyter-unauth",
		RuleName:    "Jupyter 未授权",
		RuleEnglish: "Jupyter unauthorized",
		Description: "Jupyter 未授权访问",
		Author:      "xiecat",
		FofaQuery:   `(body="ipython-main-app" && title="Home Page - Select or create a notebook")"`,
		Tag:         []string{"unauth"},
		Type:        TypeInline,
		Source:      "",
	},
	//log4j2
	{
		Id:          "fx-2021-11001",
		Query:       "APACHE-ActiveMQ",
		RuleName:    "APACHE ActiveMQ",
		RuleEnglish: "APACHE ActiveMQ",
		Description: ">Apache ActiveMQ® 是最流行的开源、多协议、基于 Java 的消息代理。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-ActiveMQ"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11002",
		Query:       "Apache_OFBiz",
		RuleName:    "Apache OFBiz",
		RuleEnglish: "Apache OFBiz",
		Description: "Apache OFBiz是一个开源的企业资源计划系统。它提供了一套集成和自动化许多企业的业务流程的企业应用程序。",
		Author:      "fofa",
		FofaQuery:   `app="Apache_OFBiz"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11003",
		Query:       "Jenkins",
		RuleName:    "Jenkins",
		RuleEnglish: "Jenkins",
		Description: "Jenkins是一款由Java编写的开源的持续集成工具。",
		Author:      "fofa",
		FofaQuery:   `app="Jenkins"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11004",
		Query:       "RabbitMQ",
		RuleName:    "RabbitMQ",
		RuleEnglish: "RabbitMQ",
		Description: "RabbitMQ 是最流行的开源消息代理之一。RabbitMQ 是轻量级的，易于在本地和云中部署。",
		Author:      "fofa",
		FofaQuery:   `app="RabbitMQ"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11005",
		Query:       "Apache-log4j2-Web",
		RuleName:    "Apache log4j2 Web",
		RuleEnglish: "Apache log4j2 Web",
		Description: "Apache log4j2 是对 Web servlet  containers的支持。",
		Author:      "fofa",
		FofaQuery:   `app="Apache-log4j2-Web"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11006",
		Query:       "Jedis",
		RuleName:    "Jedis",
		RuleEnglish: "Jedis",
		Description: "Jedis是Redis官方推荐的Java连接开发工具。",
		Author:      "fofa",
		FofaQuery:   `app="Jedis"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11007",
		Query:       "APACHE-tika",
		RuleName:    "APACHE tika",
		RuleEnglish: "APACHE tika",
		Description: "Apache Tika是一个用java编写的内容检测和分析框架，是Apache的Lucene项目的子项目。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-tika"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11009",
		Query:       "致远互联-FE",
		RuleName:    "致远互联 FE",
		RuleEnglish: "致远互联 FE",
		Description: "智远互联推出的办公自动化（Office Automation，简称OA），是将计算机、通信等现代化技术运用到传统办公方式，进而形成的一种新型办公方式。",
		Author:      "fofa",
		FofaQuery:   `app="致远互联-FE"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11010",
		Query:       "泛微-EMobile",
		RuleName:    "泛微 EMobile",
		RuleEnglish: "泛微 EMobile",
		Description: "e-mobile一款由上海泛微网络科技股份有限公司开发的移动办公产品，专门为手机、平板电脑等移动终端用户精心打造的移动办公产品。",
		Author:      "fofa",
		FofaQuery:   `app="泛微-EMobile"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11011",
		Query:       "泛微-E-Weaver",
		RuleName:    "泛微 E Weaver",
		RuleEnglish: "泛微 E Weaver",
		Description: "泛微协同管理平台e-weaver继承e-cology八大功能模块应用，并可进一步打通企业更深层的个性管理需求，基于“协同”思想打造全面整合企业管理资源的环境。",
		Author:      "fofa",
		FofaQuery:   `app="泛微-E-Weaver"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11012",
		Query:       "泛微-协同办公OA",
		RuleName:    "泛微 协同办公OA",
		RuleEnglish: "泛微 协同办公OA",
		Description: "泛微e-cology依托全新的设计理念，全新的管理思想，为中大型组织创建全新的高效协同办公环境。",
		Author:      "fofa",
		FofaQuery:   `app="泛微-协同办公OA"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11013",
		Query:       "致远互联-OA",
		RuleName:    "致远互联 OA",
		RuleEnglish: "致远互联 OA",
		Description: "致远互联推出的办公自动化（Office Automation，简称OA），是将计算机、通信等现代化技术运用到传统办公方式，进而形成的一种新型办公方式。",
		Author:      "fofa",
		FofaQuery:   `app="致远互联-OA"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11014",
		Query:       "致远A6",
		RuleName:    "致远A6",
		RuleEnglish: "致远A6",
		Description: "用友致远A6协同管理系统：面向广大的企事业组织应用设计，是一个基于互联网的高效协同工作平台和优秀的协同管理系统。",
		Author:      "fofa",
		FofaQuery:   `app="致远A6"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11015",
		Query:       "致远A8",
		RuleName:    "致远A8",
		RuleEnglish: "致远A8",
		Description: "A8协同管理软件是面向中型、大型、集团型组织机构、涉外工作组织及组织群，针对其异地管理、跨区域分支机构、跨地域审批等大范围协作应用设计的集团管控和信息资源管控的综合平台。",
		Author:      "fofa",
		FofaQuery:   `app="致远A8"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11016",
		Query:       "用友-ERP-NC",
		RuleName:    "用友 ERP NC",
		RuleEnglish: "用友 ERP NC",
		Description: "用友NC服务大型企业互联网化，互联网创造了新的经济形态和运行模式，正在成为经济发展新引擎，企业互联网化是大型企业信息化的新进阶。",
		Author:      "fofa",
		FofaQuery:   `app="用友-ERP-NC"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11017",
		Query:       "用友-GRP-U8",
		RuleName:    "用友 GRP U8",
		RuleEnglish: "用友 GRP U8",
		Description: "用友GRP-U8是由用友开发的行政事业单位内控的管理软件。",
		Author:      "fofa",
		FofaQuery:   `app="用友-GRP-U8"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11018",
		Query:       "用友-UFIDA-NC",
		RuleName:    "用友-UFIDA-NC",
		RuleEnglish: "用友 UFIDA NC",
		Description: "UFIDA NC是高端管理软件产品，提供了包括财务会计、管理会计、资金管理、资产管理、预算管理、人力资源管理、供应链、分销、多工厂制造、分析决策、合并报表等完整的解决方案。",
		Author:      "fofa",
		FofaQuery:   `app="用友-UFIDA-NC"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11019",
		Query:       "jeewms",
		RuleName:    "jeewms",
		RuleEnglish: "jeewms",
		Description: "jeewms是一个基于JAVA的仓库管理系统,是由灵鹿谷科技主导的开源项目,WMS在经过多家公司上线运行后,为了降低物流仓储企业的信息化成本,决定全面开源此产品。",
		Author:      "fofa",
		FofaQuery:   `app="jeewms"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11020",
		Query:       "APACHE-Skywalking",
		RuleName:    "APACHE Skywalking",
		RuleEnglish: "APACHE Skywalking",
		Description: "SkyWalking 是一个可观察性分析平台和应用程序性能管理系统。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-Skywalking"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11021",
		Query:       "Struts2",
		RuleName:    "Struts2",
		RuleEnglish: "Struts2",
		Description: "Apache Struts 2是一个用于开发Java EE网络应用程序的开放源代码网页应用程序架构。",
		Author:      "fofa",
		FofaQuery:   `app="Struts2"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11022",
		Query:       "APACHE-Shiro",
		RuleName:    "APACHE Shiro",
		RuleEnglish: "APACHE Shiro",
		Description: "Apache Shiro是一个开源安全框架，提供身份验证、授权、密码学和会话管理。Shiro框架直观、易用，同时也能提供健壮的安全性。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-Shiro"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11023",
		Query:       "JEECMS",
		RuleName:    "JEECMS",
		RuleEnglish: "JEECMS",
		Description: "jeecms基于java、vue、springboot等技术自主研发的开源内容管理系统(cms, java cms,jsp cms),可高效快捷新闻采编、发布、模板设计制作,具有性能稳定,安全,易扩展等特点",
		Author:      "fofa",
		FofaQuery:   `app="JEECMS"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11024",
		Query:       "JeeSite",
		RuleName:    "JeeSite",
		RuleEnglish: "JeeSite",
		Description: "JeeSite 是 Spring Boot 目前最好的快速开发平台, Java 开源框架, 使用经典技术组合：SpringBoot、SpringCloud、MyBatis、Shiro、Beetl、Bootstrap、AdminLTE 等。",
		Author:      "fofa",
		FofaQuery:   `app="JeeSite"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11025",
		Query:       "APACHE-dubbo",
		RuleName:    "APACHE dubbo",
		RuleEnglish: "APACHE dubbo",
		Description: "Apache Dubbo 是一款高性能、轻量级的开源服务框架。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-dubbo"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11026",
		Query:       "OPENCms",
		RuleName:    "OPENCms",
		RuleEnglish: "OPENCms",
		Description: "OpenCms 是一个专业的、易于使用的网站内容管理系统。",
		Author:      "fofa",
		FofaQuery:   `app="OPENCms"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11027",
		Query:       "JEECG",
		RuleName:    "JEECG",
		RuleEnglish: "JEECG",
		Description: "JeecgBoot是一款基于BPM的低代码平台！前后端分离架构 SpringBoot 2.x，SpringCloud，Ant Design&amp;Vue，Mybatis-plus，Shiro，JWT，支持微服务。",
		Author:      "fofa",
		FofaQuery:   `app="JEECG"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11028",
		Query:       "Jeeplus",
		RuleName:    "Jeeplus",
		RuleEnglish: "Jeeplus",
		Description: "Jeeplus，一个集成了代码生成器的java快速开发框架。",
		Author:      "fofa",
		FofaQuery:   `app="Jeeplus"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11029",
		Query:       "MyBatis",
		RuleName:    "MyBatis",
		RuleEnglish: "MyBatis",
		Description: "MyBatis是一个Java持久化框架，它通过XML描述符或注解把对象与存储过程或SQL语句关联起来，映射成数据库内对应的纪录。",
		Author:      "fofa",
		FofaQuery:   `app="MyBatis"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11030",
		Query:       "vmware-SpringBoot-Framework",
		RuleName:    "vmware SpringBoot Framework",
		RuleEnglish: "vmware SpringBoot Framework",
		Description: "Spring Boot是由Pivotal团队提供的全新框架，其设计目的是用来简化新Spring应用的初始搭建以及开发过程。",
		Author:      "fofa",
		FofaQuery:   `app="vmware-SpringBoot-Framework"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11031",
		Query:       "APACHE-Solr",
		RuleName:    "APACHE-Solr",
		RuleEnglish: "APACHE-Solr",
		Description: "Solr是Apache Lucene项目的开源企业搜索平台。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-Solr"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11032",
		Query:       "Map/Reduce",
		RuleName:    "Map/Reduce",
		RuleEnglish: "Map/Reduce",
		Description: "Apache Hadoop是一款支持数据密集型分布式应用程序并以Apache 2.0许可协议发布的开源软件框架。",
		Author:      "fofa",
		FofaQuery:   `app="Map/Reduce"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11033",
		Query:       "CLOUDERA-Hadoop-Hue",
		RuleName:    "CLOUDERA Hadoop Hue",
		RuleEnglish: "CLOUDERA Hadoop Hue",
		Description: "Apache Hadoop是一款支持数据密集型分布式应用程序并以Apache 2.0许可协议发布的开源软件框架。",
		Author:      "fofa",
		FofaQuery:   `app="CLOUDERA-Hadoop-Hue"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11034",
		Query:       "APACHE-hadoop-YARN",
		RuleName:    "APACHE hadoop YARN",
		RuleEnglish: "APACHE hadoop YARN",
		Description: "Apache Hadoop是一款支持数据密集型分布式应用程序并以Apache 2.0许可协议发布的开源软件框架。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-hadoop-YARN"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11035",
		Query:       "APACHE-hadoop-HttpFS",
		RuleName:    "APACHE hadoop HttpFS",
		RuleEnglish: "APACHE hadoop HttpFS",
		Description: "Apache Hadoop是一款支持数据密集型分布式应用程序并以Apache 2.0许可协议发布的开源软件框架。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-hadoop-HttpFS"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11036",
		Query:       "Oracle-E-Business-Suite",
		RuleName:    "Oracle E Business-Suite",
		RuleEnglish: "Oracle E Business-Suite",
		Description: "Oracle 电子商务套件支持当今不断发展的业务模型、提高生产力并满足现代移动用户的需求。",
		Author:      "fofa",
		FofaQuery:   `app="Oracle-E-Business-Suite"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11037",
		Query:       "splunk-日志分析",
		RuleName:    "splunk 日志分析",
		RuleEnglish: "splunk 日志分析",
		Description: "Splunk 是用于机器生成数据、非结构化/结构化和复杂多行数据的集中日志分析工具。",
		Author:      "fofa",
		FofaQuery:   `app="splunk-日志分析"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11038",
		Query:       "vmware-vCenter",
		RuleName:    "vmware vCenter",
		RuleEnglish: "vmware vCenter",
		Description: "VMware vCenter Server 提供了一个可伸缩、可扩展的平台，为虚拟化管理奠定了基础。",
		Author:      "fofa",
		FofaQuery:   `app="vmware-vCenter"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11039",
		Query:       "APACHE-Storm",
		RuleName:    "APACHE Storm",
		RuleEnglish: "APACHE Storm",
		Description: "Storm 是一个分布式实时大数据处理系统，可以帮助我们方便地处理海量数据，具有高可靠、高容错、高扩展的特点。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-Storm"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11040",
		Query:       "用友-NC-Cloud",
		RuleName:    "用友 NC Cloud",
		RuleEnglish: "用友 NC Cloud",
		Description: "NC Cloud，帮助企业打造数字化商业创新平台，实现业务创新、管理变革、金融嵌入的全面数字化。",
		Author:      "fofa",
		FofaQuery:   `app="用友-NC-Cloud"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11041",
		Query:       "APACHE-Druid",
		RuleName:    "APACHE Druid",
		RuleEnglish: "APACHE Druid",
		Description: "Druid是用Java编写的面向列的，开放源代码的分布式数据存储。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-Druid"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11042",
		Query:       "vmware-Spring-Framework",
		RuleName:    "vmware Spring-Framework",
		RuleEnglish: "vmware Spring-Framework",
		Description: "Spring使Java编程对每个人都更快、更容易、更安全。",
		Author:      "fofa",
		FofaQuery:   `app="vmware-Spring-Framework"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11043",
		Query:       "JavaMelody",
		RuleName:    "JavaMelody",
		RuleEnglish: "JavaMelody",
		Description: "JavaMelody 的目标是在 QA 和生产环境中监控 Java 或 Java EE 应用程序。",
		Author:      "fofa",
		FofaQuery:   `app="JavaMelody"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11044",
		Query:       "openNMS-产品",
		RuleName:    "openNMS 产品",
		RuleEnglish: "openNMS 产品",
		Description: "OpenNMS是一个免费的开源企业级网络监视和网络管理平台。",
		Author:      "fofa",
		FofaQuery:   `app="openNMS-产品"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11045",
		Query:       "APACHE-Unomi",
		RuleName:    "APACHE Unomi",
		RuleEnglish: "APACHE Unomi",
		Description: "Apache Unomi是一个 Java 开源客户数据平台，是一个 Java 服务器。",
		Author:      "fofa",
		FofaQuery:   `app="APACHE-Unomi"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11046",
		Query:       "ECLIPSE-jetty",
		RuleName:    "ECLIPSE jetty",
		RuleEnglish: "ECLIPSE jetty",
		Description: "Jetty是一个纯粹的基于Java的网页服务器和Java Servlet容器。",
		Author:      "fofa",
		FofaQuery:   `app="ECLIPSE-jetty"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11047",
		Query:       "elastic-Elasticsearch",
		RuleName:    "elastic Elasticsearch",
		RuleEnglish: "elastic Elasticsearch",
		Description: "Elasticsearch是一个基于Lucene库的搜索引擎。",
		Author:      "fofa",
		FofaQuery:   `app="elastic-Elasticsearch"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11048",
		Query:       "RedHat-Jboss",
		RuleName:    "RedHat Jboss",
		RuleEnglish: "RedHat Jboss",
		Description: "一种开源的Java EE 8 兼容应用服务器，让您只需构建Java 应用一次便可部署到所有位置。",
		Author:      "fofa",
		FofaQuery:   `app="RedHat-Jboss"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11049",
		Query:       "Openfire",
		RuleName:    "Openfire",
		RuleEnglish: "Openfire",
		Description: "Openfire 是根据开源 Apache 许可证获得许可的实时协作 (RTC) 服务器。",
		Author:      "fofa",
		FofaQuery:   `app="Openfire"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11050",
		Query:       "Oracle-BEA-WebLogic-Server",
		RuleName:    "Oracle BEA WebLogic Server",
		RuleEnglish: "Oracle BEA WebLogic Server",
		Description: "BEA WebLogic Server是一个可扩展的, Java Two Enterprise Edition (J2EE) 应用服务器。",
		Author:      "fofa",
		FofaQuery:   `app="Oracle-BEA-WebLogic-Server"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11051",
		Query:       "Oracle-Weblogic_interface_7001",
		RuleName:    "Oracle Weblogic interface 7001",
		RuleEnglish: "Oracle Weblogic interface 7001",
		Description: "WebLogic是美商Oracle的主要产品之一，系购并得来。是商业市场上主要的Java应用服务器软件之一，是世界上第一个成功商业化的J2EE应用服务器。",
		Author:      "fofa",
		FofaQuery:   `app="Oracle-Weblogic_interface_7001"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11052",
		Query:       "Oracle-BI-Publisher-Enterprise",
		RuleName:    "Oracle BI-Publisher Enterprise",
		RuleEnglish: "Oracle BI Publisher Enterprise",
		Description: "Oracle XML Publisher是Oracle Corporation的最新报告技术",
		Author:      "fofa",
		FofaQuery:   `app="Oracle-BI-Publisher-Enterprise"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11053",
		Query:       "vmware-vSphere-Web-Client",
		RuleName:    "vmware vSphere Web-Client",
		RuleEnglish: "vmware vSphere Web-Client",
		Description: "vSphere Web Client是为忙碌的管理员提供的一款通用的、基于浏览器的VMware管理工具，能够监控并管理VMware基础设施。",
		Author:      "fofa",
		FofaQuery:   `app="vmware-vSphere-Web-Client"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11054",
		Query:       "RedHat-JBoss-AS",
		RuleName:    "RedHat JBoss AS",
		RuleEnglish: "RedHat JBoss AS",
		Description: "一种开源的Java EE 8 兼容应用服务器，让您只需构建Java 应用一次便可部署到所有位置。",
		Author:      "fofa",
		FofaQuery:   `app="RedHat-JBoss-AS"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
	{
		Id:          "fx-2021-11055",
		Query:       "用友-ism",
		RuleName:    "用友 ism",
		RuleEnglish: "yongyou ism",
		Description: "iSM智能服务管理器。",
		Author:      "fofa",
		FofaQuery:   `app="用友-ism"`,
		Tag:         []string{"log4j2", "fofa"},
		Type:        TypeInline,
		Source:      "https://fofa.so/static_pages/log4j2",
	},
}
