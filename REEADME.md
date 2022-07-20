# beego_paper
前端模板使用的是B站wzomg大佬的模板，后端使用beego框架并添加了md5的验证(初学练手项目)

### 项目简介

golang + redis + mysql + html5 + css3 + jquery + beego + md5

### 快速上手

1、安装golang开发环境；git，bee工具：
```
`go get -u github.com/beego/bee`
```

2、修改2个配置文件中的redis和mysql登录信息：<br>
.../controllers/utils/redisPool.go、<br>
.../conf/app/conf

3、需要安装的依赖包
```
go get github.com/gomodule/redigo/redis
go get github.com/astaxie/beego/orm
go get -u github.com/go-sql-driver/mysql
```
