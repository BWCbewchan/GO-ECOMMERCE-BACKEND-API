package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//2.
	// logger := zap.NewExample()
	// logger.Info("Hello Example!")

	// //DEvenlopment
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development!")

	// //Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production!")

	//3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

}

// format log
func getEncoderLog() zapcore.Encoder {
	//1721744687.602859 -> 2024-07-23T21:24:47.602+0700
	encodeCofig := zap.NewProductionEncoderConfig()
	encodeCofig.EncodeTime = zapcore.ISO8601TimeEncoder

	//ts -> time
	encodeCofig.TimeKey = "time"

	//from info INFO
	encodeCofig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller":"cli/main.log.go:16"
	encodeCofig.EncodeCaller = zapcore.ShortCallerEncoder //

	return zapcore.NewConsoleEncoder(encodeCofig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
