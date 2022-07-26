package log

import (
	"context"
	"errors"
	"fmt"
	ormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"novel/woodlsy"
	"time"
)

type Int ormLogger.Interface

type Config struct {
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	LogLevel                  ormLogger.LogLevel
}

type DbLogger struct {
	LogLevel ormLogger.LogLevel
	Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func NewDbLogger(level ormLogger.LogLevel, config Config) *DbLogger {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = ormLogger.Green + "%s\n" + ormLogger.Reset + ormLogger.Green + "[info] " + ormLogger.Reset
		warnStr = ormLogger.BlueBold + "%s\n" + ormLogger.Reset + ormLogger.Magenta + "[warn] " + ormLogger.Reset
		errStr = ormLogger.Magenta + "%s\n" + ormLogger.Reset + ormLogger.Red + "[error] " + ormLogger.Reset
		traceStr = ormLogger.Green + "%s\n" + ormLogger.Reset + ormLogger.Yellow + "[%.3fms] " + ormLogger.BlueBold + "[rows:%v]" + ormLogger.Reset + " %s"
		traceWarnStr = ormLogger.Green + "%s " + ormLogger.Yellow + "%s\n" + ormLogger.Reset + ormLogger.RedBold + "[%.3fms] " + ormLogger.Yellow + "[rows:%v]" + ormLogger.Magenta + " %s" + ormLogger.Reset
		traceErrStr = ormLogger.RedBold + "%s " + ormLogger.MagentaBold + "%s\n" + ormLogger.Reset + ormLogger.Yellow + "[%.3fms] " + ormLogger.BlueBold + "[rows:%v]" + ormLogger.Reset + " %s"
	}

	return &DbLogger{
		LogLevel:     level,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

// LogMode log mode
func (l *DbLogger) LogMode(level ormLogger.LogLevel) ormLogger.Interface {
	fmt.Println(1, "===")
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l DbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	fmt.Println(2, "===")
	if l.LogLevel >= ormLogger.Info {
		Logger.Info(l.infoStr+msg, utils.FileWithLineNum(), data)
		//l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l DbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	fmt.Println(3, "===")
	if l.LogLevel >= ormLogger.Warn {
		Logger.Warn(l.warnStr+msg, utils.FileWithLineNum(), data)
	}
}

// Error print error messages
func (l DbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	fmt.Println(4, "===")
	Logger.Error(1111)
	if l.LogLevel >= ormLogger.Error {
		Logger.Error(l.errStr+msg, utils.FileWithLineNum(), data)
	}
}

// Trace print sql message
func (l DbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	//if l.LogLevel <= ormLogger.Silent {
	//	return
	//}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= ormLogger.Error && (!errors.Is(err, ormLogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			Logger.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Logger.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= ormLogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			Logger.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Logger.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case woodlsy.Configs.App.PSql:
		sql, rows := fc()
		if rows == -1 {
			Logger.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Logger.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}