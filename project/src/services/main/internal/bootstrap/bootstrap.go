// Package bootstrap
package bootstrap

import (
	"echo-inertia.com/src/internal/configs"
	"echo-inertia.com/src/pkg/utils"

	"runtime"
	"time"

	"echo-inertia.com/src/internal/infra/redis"
	"echo-inertia.com/src/services/main/internal/jobs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/robfig/cron/v3"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Bootstrap struct {
	Job           *jobs.Job
	Cron          *cron.Cron
	Logger        *zap.Logger
	Config        *configs.Config
	Translator    *utils.Translator
	RedisDB       *redis.RedisClient
	RollingLogger *lumberjack.Logger
}

func MustGetBootstrapInstance() *Bootstrap {
	cfg, err := configs.LoadConfig()
	if err != nil {
		utils.FatalResult("failed to load config", err)
	}
	translator := utils.NewTranslator("pt")
	b := Bootstrap{
		Config:     cfg,
		Cron:       cron.New(),
		Translator: translator,
		RedisDB:    redis.NewRedisClient(cfg.RedisURL),
	}
	b.setUpLogger()
	b.setupJobs()

	// REPOSITORIES

	// USECASES

	// HANDLERS

	// set bootstrap handlers
	return &b
}

func (b *Bootstrap) setupJobs() {
	b.Job = jobs.NewJob(
		b.Cron,
		b.Logger,
	)
}

func (b *Bootstrap) setUpLogger() {
	if b.RollingLogger == nil {
		logFile, err := utils.GetDefaultLogsFileName()
		if err != nil {
			utils.FatalResult("unable to set logs file", err)
		}
		b.RollingLogger = &lumberjack.Logger{
			Filename:   logFile, // Will rotate daily
			MaxSize:    10,      // Max size in MB
			MaxBackups: 7,       // Keep 7 old logs
			MaxAge:     30,      // Keep logs for 30 days
			Compress:   false,   // Compress old logs
		}
	}
	if b.Logger == nil {
		logWriter := zapcore.AddSync(b.RollingLogger)
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), logWriter, zap.InfoLevel)
		b.Logger = zap.New(core, zap.AddCaller())
	}
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {
			utils.FatalResult("unable to defer logger sync %v", err)
		}
	}(b.Logger)
}

func (b *Bootstrap) LogSystemUsage() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		// cpuPercent, _ := cpu.Percent(0, false) // Get CPU usage
		b.Logger.Info("SYSTEM_USAGE",
			// zap.Float64("CPU (%)", cpuPercent[0]),
			zap.Uint64("Memory Alloc (MiB)", memStats.Alloc/1024/1024),
			zap.Uint64("Memory Sys (MiB)", memStats.Sys/1024/1024),
			zap.Int("Goroutines", runtime.NumGoroutine()),
		)
	}
}

func (b *Bootstrap) RunJobs() {
	b.Job.RotateLogs()
	b.Cron.Start()
}
