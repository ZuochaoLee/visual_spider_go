# visual_spider_go
##a visual spider by go
![](https://github.com/ZuochaoLee/visual_spider_go/blob/master/static/img/logo.png)
#简介
visual_spider_go是使用golang编写的可视化爬虫框架，目标是尽可能降低爬虫开发周期，对垂直领域爬虫进行了高度的抽象，用户只需要简单配置不需要编写代码就能完成数据爬取，所以与其说是爬虫框架不如说是一个完整的数据爬取软件产品，特别适合没有编程能力，但掌握基本的计算机操作知识的用户使用。正如本项目logo所示，“小孩子”都可以爬数据了
本项目基于go_spider，结合了ip代理等全部反反爬模块，拥有可视化编辑维护爬取规则功能，系统自带golang原生的任务定时调度模块，可以自己管理定时执行爬虫任务，同时支持动态修改header，用户登录等功能。
#安装
##依赖
 ###go get github.com/ZuochaoLee/go-candyjs
 ###go get github.com/ZuochaoLee/go_spider
 ###go get github.com/ZuochaoLee/cron
 ###go get github.com/astaxie/beego

##go get github.com/ZuochaoLee/visual_spider_go
