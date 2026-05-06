package errorDomain

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
)

const XRequestIdHeader = "X-Request-Id"
const XTenantIDHeader = "X-Tenant-ID"
const AuthorizationHeader = "Authorization"
const MerchantApiKeyHeader = "Merchant-Api-Key"
const StoreApiKeyHeader = "Store-Api-Key"
const TraceIdKey = "traceId"
const requestBodyKey = "requestBody"
const XTenantIdKey = "xTenantId"
const MerchantApiKey = "merchantApiKey"
const StoreApiKey = "storeApiKey"
const errorCodeKey = "errorCode"
const errorRawKey = "errorRaw"
const projectKey = "project"
const envKey = "env"

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.FunctionKey = "func"

	opts := []zap.Option{
		//zap.AddCaller(),
		zap.AddCallerSkip(1),
	}

	var err error
	logger, err = config.Build(opts...)
	if err != nil {
		fmt.Println("Failed to create logger")
		return
	}
}

func AddEnvs(ctx *context.Context, fields []zap.Field) []zap.Field {
	project := os.Getenv("PROJECT")
	env := os.Getenv("ENV")
	if project != "" {
		fields = append(fields, zap.String(projectKey, project))
	}
	if env != "" {
		fields = append(fields, zap.String(envKey, env))
	}
	if ctx == nil {
		return fields
	}

	traceID, hasTraceID := (*ctx).Value(TraceIdKey).(string)
	if hasTraceID {
		fields = append(fields, zap.String(TraceIdKey, traceID))
	}
	xTenantID, hasXTenantID := (*ctx).Value(XTenantIdKey).(string)
	if hasXTenantID {
		fields = append(fields, zap.String(XTenantIdKey, xTenantID))
	}
	merchantID, hasMerchantID := (*ctx).Value(MerchantApiKey).(string)
	if hasMerchantID {
		fields = append(fields, zap.String(MerchantApiKey, merchantID))
	}
	storeID, hasStoreID := (*ctx).Value(StoreApiKey).(string)
	if hasStoreID {
		fields = append(fields, zap.String(StoreApiKey, storeID))
	}
	return fields
}

func Info(ctx *context.Context, message string) {
	if logger == nil {
		return
	}
	fields := make([]zap.Field, 0)
	fields = AddEnvs(ctx, fields)
	logger.Info(message, fields...)
}

func Error(ctx *context.Context, message string, errorCode string, errorRaw string) {
	if logger == nil {
		return
	}
	fields := make([]zap.Field, 0)
	fields = AddEnvs(ctx, fields)
	if errorCode != "" {
		fields = append(fields, zap.String(errorCodeKey, errorCode))
	}
	if errorRaw != "" {
		fields = append(fields, zap.String(errorRawKey, errorRaw))
	}
	logger.Error(message, fields...)
}

func Fatal(ctx *context.Context, message string, errorCode string, errorRaw string) {
	if logger == nil {
		return
	}
	fields := make([]zap.Field, 0)
	fields = AddEnvs(ctx, fields)
	if errorCode != "" {
		fields = append(fields, zap.String(errorCodeKey, errorCode))
	}
	if errorRaw != "" {
		fields = append(fields, zap.String(errorRawKey, errorRaw))
	}

	//logger.Fatal(message, fields...)
	logger.Error(message, fields...)
}

func EndHttpLog(
	ctx context.Context, httpStatus int, message string,
) {
	if logger == nil {
		return
	}
	fields := make([]zap.Field, 0)
	fields = AddEnvs(&ctx, fields)

	logType := LevelInfo
	if httpStatus >= 400 && httpStatus < 500 {
		logType = LevelWarning
	} else if httpStatus >= 500 {
		logType = LevelError
	}

	switch logType {
	case LevelInfo:
		logger.Info(message, fields...)
		break
	case LevelError:
		logger.Error(message, fields...)
		break
	case LevelWarning:
		logger.Warn(message, fields...)
		break
	}
}

func StartHttpLog(
	ctx *context.Context, body string,
) {
	if logger == nil {
		return
	}
	fields := make([]zap.Field, 0)
	fields = AddEnvs(ctx, fields)
	fields = append(fields, zap.String(requestBodyKey, body))
	logger.Info("Request received", fields...)
}
