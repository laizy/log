package ext

import (
	"github.com/laizy/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

// create rolling file handler
func RollingFileHandler(fileName string) log.Handler {
	return log.StreamHandler(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    200, // megabytes
		MaxAge:     28,  //days
		MaxBackups: 100,
		LocalTime:  true,
		Compress:   false,
	}, log.LogfmtWithGIDFormat())
}
