package log

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

// 初始化日志
func Init() {
	Logger = zap.NewExample()
	defer Logger.Sync()

	// todo 配置日志的基本

	// 测试打印
	//url := "http://example.org/api"
	//Logger.Info("failed to fetch URL",
	//	zap.String("url", url),
	//	zap.Int("attempt", 3),
	//	zap.Duration("backoff", time.Second),
	//)

}
