package help

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var  Log *zap.SugaredLogger

func InitZap(){

	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	path := "./runtime/log/gin.log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    10,//MB
		MaxBackups: 30,//保留旧文件最大个数
		MaxAge:     30,//保留旧文件最大天数
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}
