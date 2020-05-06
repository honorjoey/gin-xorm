package log

import "testing"

// import "github.com/honorjoey/gin-xorm/log"
type TLog struct {
	Log
}

var tLog TLog

func TestLog_Info(t *testing.T) {
	tLog.Tag = "Test"

	tLog.Info("test...")
	// [Info] [Test] 2020/05/05 19:40:49.549948 test...
}
