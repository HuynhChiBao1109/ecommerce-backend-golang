package initialize

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	enCoder := configEncode()
	sync := writeLogSync()
	core := zapcore.NewCore(enCoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Hello, World!", zap.Int("number", 1))
	logger.Error("Error", zap.String("error", "error"))
}

func configEncode() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// Parse time to dd-mm-yyyy hh:mm:ss
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// Info to INFO with color
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// Caller to full path
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// change time key to time
	encodeConfig.TimeKey = "time"

	return zapcore.NewJSONEncoder(encodeConfig)
}

func writeLogSync() zapcore.WriteSyncer {
	// Open the log file with append mode
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, 0644)

	// Combine file and console outputs
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
