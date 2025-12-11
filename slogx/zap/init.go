package zap

import (
	"fmt"
	"log/slog"
	"math"
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

var (
	globalZapLevel  = zap.NewAtomicLevelAt(zap.InfoLevel)
	globalZapLogger atomic.Pointer[zap.Logger]
	globalSLogger   *slog.Logger
)

func init() {
	logConfig := zap.NewProductionConfig()

	logConfig.Level = globalZapLevel
	logConfig.EncoderConfig.TimeKey = slog.TimeKey
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger := zap.Must(logConfig.Build())

	globalZapLogger.Store(zapLogger)

	baseHandler := zapslog.NewHandler(zapLogger.Core(),
		zapslog.AddStacktraceAt(math.MaxInt),
	)

	globalSLogger = slog.New(baseHandler)
}

// Init initializes the zap logger and sets it as the default slog logger
func Init() {
	slog.SetDefault(globalSLogger)
}

// Sync flushes any buffered log entries
func Sync() error {
	zapLogger := globalZapLogger.Load()

	if err := zapLogger.Sync(); err != nil {
		return fmt.Errorf("zapLogger.Sync: %w", err)
	}

	return nil
}

// SetLevel sets the logging level
func SetLevel(level slog.Level) {
	switch level {
	case slog.LevelDebug:
		globalZapLevel.SetLevel(zapcore.DebugLevel)
	case slog.LevelInfo:
		globalZapLevel.SetLevel(zapcore.InfoLevel)
	case slog.LevelWarn:
		globalZapLevel.SetLevel(zapcore.WarnLevel)
	case slog.LevelError:
		globalZapLevel.SetLevel(zapcore.ErrorLevel)
	default:
		globalZapLevel.SetLevel(zapcore.InfoLevel)
	}
}
