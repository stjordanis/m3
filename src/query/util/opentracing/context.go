// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package opentracing provides more utilities for dealing with opentracing and context.Context's.
package opentracing

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// SpanFromContextOrRoot is the same as opentracing.SpanFromContext, but instead of returning nil, it starts a
// a new root span if ctx doesn't already have an associated span, instead of returning nil. Use this over
// opentracing.StartSpanFromContext if you need access to the current span, but don't want to start a child span.
// Note: the span returned by this function won't really be configured; if you want a proper span, start one
// at the root and pass it in.
func SpanFromContextOrRoot(ctx context.Context) (opentracing.Span, context.Context) {
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		sp = opentracing.StartSpan("SpanFromContextOrRoot - dummy")
	}

	ctx = opentracing.ContextWithSpan(ctx, sp)
	return sp, ctx
}

// Time is a log.Field for time.Time values. It translates to RF3339 formatted time strings.
func Time(key string, t time.Time) log.Field {
	return log.String(key, t.Format(time.RFC3339))
}

// Duration is a log.Field for Duration values. It translates to the standard Go duration format (Duration.String()).
func Duration(key string, t time.Duration) log.Field {
	return log.String(key, fmt.Sprint(t))
}
