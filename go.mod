module gin-xorm

go 1.13

require (
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.55.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/gorilla/sessions v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
	xorm.io/core v0.7.2-0.20190928055935-90aeac8d08eb
)

replace (
	github.com/honorjoey/gin-xorm/config => ./gin-xorm/config
	github.com/honorjoey/gin-xorm/db => ./gin-xorm/db
	github.com/honorjoey/gin-xorm/models => ./gin-xorm/models
	github.com/honorjoey/gin-xorm/router => ./gin-xorm/router
	github.com/honorjoey/gin-xorm/router/middlewares => ./gin-xorm/router/middlewares
	github.com/honorjoey/gin-xorm/utils => ./gin-xorm/utils
)
