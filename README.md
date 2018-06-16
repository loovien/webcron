# webcron
------------

一个定时任务管理器基于Go语言和beego框架开发。颗粒度支持到秒级。用于统一管理项目中的定时任务，提供可视化配置界面、执行日志记录、邮件通知等功能，无需依赖*unix下的crontab服务。

项目fork至[lishijie](https://github.com/lisijie)，做了部分改动

1. 任务支持搜索.t
2. 首页dashboard错误任务查看后, 支持删除.

## 截图

![screenshot](https://raw.githubusercontent.com/loovien/webcron/master/screenshot.jpg)

## 功能特点

* 统一管理多种定时任务。
* 秒级定时器，使用crontab的时间表达式。
* 可随时暂停任务。
* 记录每次任务的执行结果。
* 执行结果邮件通知。
* 跨平台支持。

## 安装说明

依赖

0. **Golang**
1. **cron** [cron](https://github.com/robfig/cron) 我自己也fork一份
2. **go
2. **MySQL**


方案1使用`godep`工具(推荐)

```bash
    $ godep restore

```

方案2:

```bash
    $ go get github.com/astaxie/beego
    $ go get github.com/go-sql-driver/mysql
    $ go get github.com/loovien/cron
    $ go get github.com/loovien/webcron
```

编译:

```bash
    $ cd gopath/github.com/loovien/webcron
    $ go build
```

打开配置文件 conf/app.conf，修改相关配置。
	

创建数据库webcron，初始化数据 `install.sql`

	$ mysql -u username -p -D webcron < install.sql

运行
	
	$ ./webcron
	或
	$ nohup ./webcron 2>&1 > error.log &
	设为后台运行

访问： 

[访问:http://localhost:8000](http://localhost:8000)

帐号：admin
密码：admin888