package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"novel/woodlsy"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

func init() {

	file := openLogFile()

	encoder := getEncoder()

	consoleEncoder := getConsoleEncoder()
	newCore := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			file,
			zapcore.DebugLevel,
		), // 写入文件
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.ErrorLevel), // 写入控制台
	)

	//core := zapcore.NewCore(
	//	zapcore.NewConsoleEncoder(encoder),
	//	file,
	//	zapcore.InfoLevel,
	//)

	l := zap.New(newCore, zap.AddCaller())
	zap.ReplaceGlobals(l)
	Logger = l.Sugar()
}

func openLogFile() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   woodlsy.Configs.Log.FilePath,
		MaxSize:    woodlsy.Configs.Log.MaxSize,
		MaxBackups: woodlsy.Configs.Log.MaxBackups,
		MaxAge:     woodlsy.Configs.Log.MaxAge,
		Compress:   woodlsy.Configs.Log.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GetConsoleEncoder 输出日志到控制台
func getConsoleEncoder() zapcore.Encoder {

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//
// getEncoder
// @Description: 日志格式编码
// @return zapcore.EncoderConfig
//
func getEncoder() zapcore.EncoderConfig {
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller_line",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeDuration: zapcore.SecondsDurationEncoder,

		//MessageKey:     "message",
		//TimeKey:        "time",
		//CallerKey:      "caller_line",
		LineEnding:   zapcore.DefaultLineEnding,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeTime:   customTimeEncoder,
		EncodeCaller: cEncodeCaller,
	}
}

// cEncodeCaller 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

//
// GinLogger
// @Description: gin的日志注入
// @return gin.HandlerFunc
//
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Debug(fmt.Sprintf("[url] [%s] %s", cost, path),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			//zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			//zap.Duration("cost", cost),
		)
	}
}

func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
