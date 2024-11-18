package constants

const AppName = "ecommerce-application"
const ModuleName = "customer-service"
const GrpcServerPort = "8082"
const JwtExpiryHours = 1000 // 41 Days

// Env Variables
const (
	DatabaseUrlEnvName = "DATABASE_URL"
)

// Otel
const (
	OtelEnableEnv       = "OTEL_ENABLED"
	OtelServiceNameEnv  = "OTEL_SERVICE_NAME"
	OtelCollectorEnv    = "OTEL_COLLECTOR_URL"
	OtelInsecureModeEnv = "OTEL_INSECURE_MODE"
)
