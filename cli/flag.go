package cli

import (
	"github.com/urfave/cli/v2"
)

var (
	PrometheusFlag = []cli.Flag{
		&cli.IntFlag{
			Name:    "prometheus-port",
			Value:   7070,
			EnvVars: []string{"PROMETHEUS_PORT", "CONFIG_PROMETHEUS_PORT"},
			Usage:   "The port for exposing prometheus metrics endpoints",
		},
		&cli.BoolFlag{
			Name:    "disable-prometheus",
			EnvVars: []string{"DISABLE_PROMETHEUS", "CONFIG_PROMETHEUS_DISABLE"},
			Usage:   "disable-prometheus",
		},
	}
	FeatureToggleFlag = []cli.Flag{
		&cli.BoolFlag{
			Name:    "feature-toggle",
			EnvVars: []string{"FEATURE_TOGGLE", "CONFIG_FEATURE_TOGGLE_DISABLE"},
			Usage:   "Enable feature toggle",
		},
		&cli.StringFlag{
			Name:    "unleash-api-url",
			EnvVars: []string{"UNLEASH_API_URL", "CONFIG_FEATURE_TOGGLE_API_URL"},
			Usage:   "Unleash api url",
		},
		&cli.StringFlag{
			Name:    "unleash-token",
			EnvVars: []string{"UNLEASH_TOKEN", "CONFIG_FEATURE_TOGGLE_TOKEN"},
			Usage:   "Unleash token",
		},
		&cli.StringFlag{
			Name:    "unleash-app-name",
			EnvVars: []string{"UNLEASH_APP_NAME", "CONFIG_FEATURE_TOGGLE_NAME"},
			Usage:   "Unleash app name",
		},
	}

	CommonServerFlag = []cli.Flag{
		&cli.StringFlag{
			Name:    "runtime-version",
			EnvVars: []string{"RUNTIME_VERSION"},
			Value:   "v1.0.0",
		},
		&cli.IntFlag{
			Name:    "grpc-port",
			Value:   50051,
			EnvVars: []string{"GRPC_PORT", "CONFIG_GRPC_PORT"},
			Usage:   "The port for exposing the gRPC endpoints for accessing",
		},
		&cli.IntFlag{
			Name:    "http-port",
			Value:   80,
			EnvVars: []string{"HTTP_PORT", "PORT", "CONFIG_HTTP_PORT"},
			Usage:   "The port for exposing the api endpoints for accessing",
		},
		&cli.IntFlag{
			Name:    "pprof-port",
			Value:   6060,
			EnvVars: []string{"PPROF_PORT", "CONFIG_PPROF_PORT"},
			Usage:   "The port for exposing pprof endpoints",
		},

		&cli.BoolFlag{
			Name:    "disable-tracing",
			EnvVars: []string{"DISABLE_TRACING", "CONFIG_DISABLE_TRACING"},
			Usage:   "disable-tracing",
		},

		&cli.BoolFlag{
			Name:    "disable-profiler",
			EnvVars: []string{"DISABLE_PROFILER", "CONFIG_PROFILER_DISABLE"},
			Usage:   "disable-profiler",
		},
		&cli.BoolFlag{
			Name:    "disable-stats",
			EnvVars: []string{"DISABLE_STATS", "CONFIG_STATS_DISABLE"},
			Usage:   "disable-stats",
		},
		&cli.BoolFlag{
			Name:    "allow-kill",
			EnvVars: []string{"ALLOW_KILL", "CONFIG_ALLOW_KILL"},
			Usage:   "allow remote request to kill server",
		},
	}
	GormFlag = []cli.Flag{
		&cli.StringFlag{
			Name:     "db-uri",
			Required: true,
			EnvVars:  []string{"DB_URI", "DATABASE_URL", "CONFIG_GORM_DB_URI"},
			Usage:    "The URI for connecting to database (supported URIs: in-memory://, postgresql://auth@host:26257/linkgraph?sslmode=disable)",
		},
	}
	LoggerFlag = []cli.Flag{
		&cli.StringFlag{
			Name:    "log-level",
			EnvVars: []string{"LOG_LEVEL", "CONFIG_LOG_LEVEL"},
			Usage:   "Log level: (panic|fatal|error|warn|warning|info|debug|trace)",
			Value:   "info",
		},
		&cli.StringFlag{
			Name:    "log-format",
			EnvVars: []string{"LOG_FORMAT", "CONFIG_LOG_FORMAT"},
			Usage:   "Log format: (plain|json)",
			Value:   "plain",
		},
		&cli.StringFlag{
			Name:    "log-file-path",
			EnvVars: []string{"LOG_FILE_PATH", "CONFIG_LOG_PATH"},
			Usage:   "Log file path",
		},
	}
	RedisFlag = []cli.Flag{
		&cli.StringFlag{
			Name:     "redis-address",
			Required: true,
			EnvVars:  []string{"REDIS_ADDRESS", "CONFIG_REDIS_ADDRESS"},
			Usage:    "Redis address",
		},
		&cli.StringFlag{
			Name:    "redis-pass",
			EnvVars: []string{"REDIS_PASS", "CONFIG_REDIS_PASS"},
			Usage:   "Redis pass",
		},
		&cli.StringFlag{
			Name:    "redis-user",
			EnvVars: []string{"REDIS_USER", "CONFIG_REDIS_USER"},
			Usage:   "Redis user",
		},
	}
)
