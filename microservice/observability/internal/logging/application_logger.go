package logging

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(lvl string, kafkaBrokers []string) (*zap.Logger, error) {
	level, err := parseLevel(lvl)
	if err != nil {
		return nil, err
	}

	return getLogger(level, kafkaBrokers), nil
}

func getLogger(
	level zapcore.Level,
	kafkaBrokers []string,
) *zap.Logger {
	kafkaURL := url.URL{Scheme: "kafka", Host: strings.Join(kafkaBrokers, ",")}

	ec := zap.NewProductionEncoderConfig()
	ec.EncodeName = zapcore.FullNameEncoder
	ec.EncodeName = zapcore.FullNameEncoder
	ec.EncodeTime = zapcore.RFC3339TimeEncoder
	ec.EncodeDuration = zapcore.MillisDurationEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	ec.EncodeCaller = zapcore.ShortCallerEncoder
	ec.NameKey = "logger"
	ec.MessageKey = "message"
	ec.LevelKey = "level"
	ec.TimeKey = "timestamp"
	ec.CallerKey = "caller"
	ec.StacktraceKey = "trace"
	ec.LineEnding = zapcore.DefaultLineEnding
	ec.ConsoleSeparator = " "

	zapConfig := zap.Config{}
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.Encoding = "json"
	zapConfig.EncoderConfig = ec
	zapConfig.OutputPaths = []string{"stdout", kafkaURL.String()}
	zapConfig.ErrorOutputPaths = []string{"stderr", kafkaURL.String()}

	if kafkaBrokers != nil || len(kafkaBrokers) > 0 {
		err := zap.RegisterSink("kafka", func(_ *url.URL) (zap.Sink, error) {
			config := sarama.NewConfig()
			config.Producer.Return.Successes = true

			producer, err := sarama.NewSyncProducer(kafkaBrokers, config)
			if err != nil {
				log.Fatal("failed to create kafka producer", err)
				return kafkaSink{}, err
			}

			return kafkaSink{
				producer: producer,
				topic:    "application.logs",
			}, nil
		})
		if err != nil {
			log.Fatal("failed to register kafka sink", err)
		}
	}

	zapLogger, err := zapConfig.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		log.Fatal("failed to build zap logger", err)
	}

	return zapLogger
}

func parseLevel(lvl string) (zapcore.Level, error) {
	switch strings.ToLower(lvl) {
	case "debug":
		return zap.DebugLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "warn":
		return zap.WarnLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	}
	return zap.InfoLevel, fmt.Errorf("invalid log level <%v>", lvl)
}
