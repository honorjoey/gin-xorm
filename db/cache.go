package db

import (
	"github.com/gomodule/redigo/redis"
)

type NameDao struct {}

func (n NameDao) GetName(name string) map[string]string {
	res := make(map[string]string)
	str, err := redis.String(r.Do("get", name))
	if err != nil {
		rLog.Error("Get error", err)
		return nil
	}
	res[name]=str
	rLog.Infof("%v\n", res)
	return res
}
