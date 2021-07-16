package utils

import (
	"github.com/utahta/go-cronowriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var ZLogger *zap.SugaredLogger

func InitLogger(path string)  {
	w1 := cronowriter.MustNew(path + "application.log.%Y%m%d")
	w2 := cronowriter.MustNew(path + "application_error.log.%Y%m%d")
	l := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(w1),zapcore.AddSync(os.Stdout)),
			zapcore.InfoLevel,
		),
		zap.ErrorOutput(zapcore.AddSync(w2)),
		zap.AddCaller(),
	)
	ZLogger = l.Sugar()
}