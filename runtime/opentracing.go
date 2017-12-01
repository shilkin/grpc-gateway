package runtime

import (
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
)

// TryToCopySpan copies opentracing span from one context to another if it's possible
func TryToCopySpan(to context.Context, from context.Context) context.Context {
	if span := opentracing.SpanFromContext(from); span != nil {
		return opentracing.ContextWithSpan(to, span)
	}
	return to
}

// tryToFinishSpan finishes opentracing span if it's possible
func tryToFinishSpan(ctx context.Context) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span.Finish()
	}
}
