package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
	"time"
)

var (
	// 日志全局对象
	Logger *zap.Logger

	// 日志级别map
	levelType = map[string]zapcore.Level{
		"debug":  zap.DebugLevel,
		"info":   zap.InfoLevel,
		"warn":   zap.WarnLevel,
		"error":  zap.ErrorLevel,
		"dpanic": zap.DPanicLevel,
		"panic":  zap.InfoLevel,
	}

	// 执行一次
	once sync.Once
)

// Init 初始化日志
func Init() {
	once.Do(func() {
		fileName := viper.GetString("log.path")
		maxSize := int(viper.GetInt64("log.logsize"))
		hook := lumberjack.Logger{
			Filename: fileName, // 日志文件路径
			MaxSize:  maxSize,  // 每个日志文件保存的最大尺寸 单位：M
			//MaxBackups: 30,                       // 日志文件最多保存多少个备份
			//MaxAge:     7,                        // 文件最多保存多少天
			//Compress:   true,                     // 是否压缩
		}

		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",  // 时间
			LevelKey:       "level", // 日志级别
			CallerKey:      "line",  // 行数
			NameKey:        "logger",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder, //
			EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
			EncodeName:     zapcore.FullNameEncoder,
		}

		// 设置日志级别
		level := viper.GetString("log.level")
		var l zapcore.Level
		if v, ok := levelType[level]; ok {
			l = v
		} else {
			l = zap.InfoLevel
		}
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(l)

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),               // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到文件
			atomicLevel, // 日志级别
		)
		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()
		// 构造日志
		Logger = zap.New(core, caller)
	})
}

// timeEncoder 日志时间格式化
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
