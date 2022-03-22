package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *zap.Logger
)

// Init 初始化日志
func Init() {
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
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	//atomicLevel.SetLevel(getLogLevel())

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),               // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到文件
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 构造日志
	Logger = zap.New(core, caller)
}

// 获取日志的级别
func getLogLevel() (l zapcore.Level) {
	level := viper.GetString("log.level")
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	default: // 默认info基本
		return zapcore.InfoLevel
	}
}
