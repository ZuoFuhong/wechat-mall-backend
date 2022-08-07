package log

import (
	"context"
	"fmt"
	"log"
	"wechat-mall-backend/consts"
)

func Debugf(format string, args ...interface{}) {
	log.Printf("DEBUG "+format+"\n", args...)
}

func Infof(format string, args ...interface{}) {
	log.Printf("INFO "+format+"\n", args...)
}

func Warnf(format string, args ...interface{}) {
	log.Printf("WARN "+format+"\n", args...)
}

func Errorf(format string, args ...interface{}) {
	log.Printf("ERROR "+format+"\n", args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func DebugContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value(consts.TraceKey).(string)
	prefix := fmt.Sprintf("DEBUG %s ", traceId)
	log.Printf(prefix+format+"\n", args...)
}

func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value(consts.TraceKey).(string)
	prefix := fmt.Sprintf("INFO %s ", traceId)
	log.Printf(prefix+format+"\n", args...)
}

func WarnContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value(consts.TraceKey).(string)
	prefix := fmt.Sprintf("WARN %s ", traceId)
	log.Printf(prefix+format+"\n", args...)
}

func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value(consts.TraceKey).(string)
	prefix := fmt.Sprintf("ERROR %s ", traceId)
	log.Printf(prefix+format+"\n", args...)
}
