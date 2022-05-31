package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/fbbyqsyea/gin-framework/global"
	"github.com/fbbyqsyea/gin-framework/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 初始化日志
func initLogger() {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.CONFIG.Logger.Director); !ok {
		err := os.Mkdir(global.CONFIG.Logger.Director, os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("fatal error mkdir: [%s]", global.CONFIG.Logger.Director))
		}
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	now := time.Now().Format("2006-01-02")

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/%s/debug.log", global.CONFIG.Logger.Director, now), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/info.log", global.CONFIG.Logger.Director, now), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/warn.log", global.CONFIG.Logger.Director, now), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/error.log", global.CONFIG.Logger.Director, now), errorPriority),
	}
	logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.CONFIG.Logger.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	global.LOGGER = logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.CONFIG.Logger.StackTraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    getEncodeLevel,
		EncodeTime:     getEncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   getEncodeCaller,
	}
	switch {
	case global.CONFIG.Logger.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.CONFIG.Logger.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.CONFIG.Logger.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.CONFIG.Logger.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.CONFIG.Logger.Format == "console" {
		return zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	return zapcore.NewJSONEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := getWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 获取同步写入对象
func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
	// console格式的日志输出到stdout
	if global.CONFIG.Logger.Format == "console" {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncodeLevel 自定义日志级别显示
func getEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
}

// getEncodeTime 自定义时间格式显示
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getEncodeCaller 自定义行号显示
func getEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.FullPath())
}
