### 介绍
用`gin`和`xorm`等技术栈写的一个web项目示例，包括连接`mysql`数据库，`redis`，`kafka`等，还有`swagger`接口文档。


### 项目运行使用
```shell script
$ go run main.go
```
或
```shell script
$ go build
$ ./gin-xorm
```

### xorm工具的使用
使用`xorm`因为要写对应的库表结构的`struct`，有点麻烦，但是`xorm`提供了工具来自动生成对应的库表结构。
#### 安装
安装过程见中文文档的`XORM 工具`部分
[http://www.admpub.com:8080/xorm-manual-zh-CN/](http://www.admpub.com:8080/xorm-manual-zh-CN/)
#### 生成

首先介绍`templates/goxorm`，是`github.com/go-xorm/cmd/xorm/templates/goxorm`
文件结构
```shell script
templates/goxorm
├── config
└── struct.go.tpl
```
定义了生成库表结构的一些参数，具体可参阅官方文档。

在项目根目录下
```shell script
$ xorm reverse mysql "root:root@tcp(192.168.199.236:3306)/{dbname}?charset=utf8" templates/goxorm
```
`{dbname}`即要生成库表结构的数据库名称。`templates/goxorm`前面已介绍，这里可以自己在项目根目录创建，然后把`go-xorm`对应文件复制过来，或者直接在命令里指定到你电脑里面的对应目录， 最后还有一个参数，是指定生成库表结构文件的目录的，不写的话默认在当前目录下创建`models`文件夹。

命令执行完成会在当前目录下创建`models`文件夹，并在文件夹下创建库表结构`go`文件。
```shell script
models
└── user.go
...
```

其他数据库的使用情参阅文档
[http://www.admpub.com:8080/xorm-manual-zh-CN/](http://www.admpub.com:8080/xorm-manual-zh-CN/)

### swagger的使用
#### 安装swag
```shell script
$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
```
若 `$GOROOT/bin` 没有加入`$PATH`中，你需要执行将其可执行文件移动到`$GOBIN`下
```shell script
$ mv $GOPATH/bin/swag /usr/local/go/bin
```
验证是否安装成功
检查 `$GOBIN` 下是否有 `swag` 文件，如下：
```shell script
$ swag -v
swag version v1.6.5
```
安装 `gin-swagger`
```shell script
$ go get -u github.com/swaggo/gin-swagger@v1.2.0 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template
```
注：若无科学上网，请务必配置 Go modules proxy。
```shell script
$ go env GOPROXY=https://goproxy.cn,direct
```
#### 初始化
`controller`里面的在`router`方法上加上类似以下的注释
```golang
// @Summary List users
// @Description list user
// @Produce  json
// @Param phone query string true "phone"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /api/user [get]
func (c UserController) List(ctx *gin.Context) {
	phone := ctx.Query("phone")
	user := *c.UserDao.QueryUser(phone)

	ctx.SecureJSON(http.StatusOK, user[0])
}
```
`gin-swagger` 给出的范例：
```golang
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
```

#### 生成

然后在项目根目录下运行
```shell script
$ swag init
2020/05/05 22:49:31 Generate swagger docs....
2020/05/05 22:49:31 Generate general API Info, search dir:./
2020/05/05 22:49:31 Generating app.Response
2020/05/05 22:49:31 create docs.go at  docs/docs.go
2020/05/05 22:49:31 create swagger.json at  docs/swagger.json
2020/05/05 22:49:31 create swagger.yaml at  docs/swagger.yaml
```

然后会在项目根目录生成`docs`文件夹
```shell script
docs
├── docs.go
├── swagger.json
└── swagger.yaml
```

`router/router.go`文件中
```golang
import (
    _ "github.com/honorjoey/gin-xorm/docs"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

func  Init() *gin.Engine {
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

#### 验证
在浏览器输入`http://localhost:8080/swagger/index.html`即可看到`api`文档。


### log的使用
此项目自定义了一个`log`包，
#### 介绍
```golang
Debug(v ...interface{})
Debugf(format string, v ...interface{})
Error(v ...interface{})
Errorf(format string, v ...interface{})
Info(v ...interface{})
Infof(format string, v ...interface{})
Warn(v ...interface{})
Warnf(format string, v ...interface{})
```

#### 实现
```golang
import "github.com/honorjoey/gin-xorm/log"

type TLog struct {
	Log
}

var tLog TLog

func init() {
	tLog.Tag = "Test"

	tLog.Info("test...")
	// [Info] [Test] 2020/05/05 19:40:49.549948 test...
}
```

即可在日志中加上日期和对应的标签。