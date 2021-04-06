package jaeger

import "time"

// ReporterConfig configures the reporter. All fields are optional.
type ReporterConfig struct {
	// QueueSize controls how many spans the reporter can keep in memory before it starts dropping
	// new spans. The queue is continuously drained by a background go-routine, as fast as spans
	// can be sent out of process.
	// Can be provided by FromEnv() via the environment variable named JAEGER_REPORTER_MAX_QUEUE_SIZE
	QueueSize int `yaml:"queue-size" json:"queue-size"`

	// BufferFlushInterval controls how often the buffer is force-flushed, even if it's not full.
	// It is generally not useful, as it only matters for very low traffic services.
	// Can be provided by FromEnv() via the environment variable named JAEGER_REPORTER_FLUSH_INTERVAL
	BufferFlushInterval time.Duration `yaml:"buffer-flush-interval" json:"buffer-flush-interval"`

	// LogSpans, when true, enables LoggingReporter that runs in parallel with the main reporter
	// and logs all submitted spans. Main Configuration.Logger must be initialized in the code
	// for this option to have any effect.
	// Can be provided by FromEnv() via the environment variable named JAEGER_REPORTER_LOG_SPANS
	LogSpans bool `yaml:"log-spans" json:"log-spans"`

	// LocalAgentHostPort instructs reporter to send spans to jaeger-agent at this address.
	// Can be provided by FromEnv() via the environment variable named JAEGER_AGENT_HOST / JAEGER_AGENT_PORT
	LocalAgentHostPort string `yaml:"local-agent-host-port" json:"local-agent-host-port"`

	// DisableAttemptReconnecting when true, disables udp connection helper that periodically re-resolves
	// the agent's hostname and reconnects if there was a change. This option only
	// applies if LocalAgentHostPort is specified.
	// Can be provided by FromEnv() via the environment variable named JAEGER_REPORTER_ATTEMPT_RECONNECTING_DISABLED
	DisableAttemptReconnecting bool `yaml:"disable-attempt-reconnecting" json:"disable-attempt-reconnecting"`

	// AttemptReconnectInterval controls how often the agent client re-resolves the provided hostname
	// in order to detect address changes. This option only applies if DisableAttemptReconnecting is false.
	// Can be provided by FromEnv() via the environment variable named JAEGER_REPORTER_ATTEMPT_RECONNECT_INTERVAL
	AttemptReconnectInterval time.Duration `yaml:"attempt-reconnect-interval" json:"attempt-reconnect-interval"`

	// CollectorEndpoint instructs reporter to send spans to jaeger-collector at this URL.
	// Can be provided by FromEnv() via the environment variable named JAEGER_ENDPOINT
	CollectorEndpoint string `yaml:"collector-endpoint" json:"collector-endpoint"`

	// User instructs reporter to include a user for basic http authentication when sending spans to jaeger-collector.
	// Can be provided by FromEnv() via the environment variable named JAEGER_USER
	User string `yaml:"user" json:"user"`

	// Password instructs reporter to include a password for basic http authentication when sending spans to
	// jaeger-collector.
	// Can be provided by FromEnv() via the environment variable named JAEGER_PASSWORD
	Password string `yaml:"password" json:"password"`

	// HTTPHeaders instructs the reporter to add these headers to the http request when reporting spans.
	// This field takes effect only when using HTTPTransport by setting the CollectorEndpoint.
	HTTPHeaders map[string]string `yaml:"http-headers" json:"http-headers"`
}