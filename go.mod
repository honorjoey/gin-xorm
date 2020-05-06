module github.com/honorjoey/gin-xorm

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.55.0
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/sessions v1.2.0
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/urfave/cli v1.22.4 // indirect
	golang.org/x/net v0.0.0-20200505041828-1ed23360d12c // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/tools v0.0.0-20200505023115-26f46d2f7ef8 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
	xorm.io/core v0.7.2-0.20190928055935-90aeac8d08eb
)

replace (
	github.com/honorjoey/gin-xorm/ => ./utils
	github.com/honorjoey/gin-xorm/config => ./config
	github.com/honorjoey/gin-xorm/db => ./db
	github.com/honorjoey/gin-xorm/docs => ./docs
	github.com/honorjoey/gin-xorm/log => ./log
	github.com/honorjoey/gin-xorm/models => ./models
	github.com/honorjoey/gin-xorm/router => ./router
	github.com/honorjoey/gin-xorm/router/middlewares => ./router/middlewares
	github.com/honorjoey/gin-xorm/utils => ./utils
)
