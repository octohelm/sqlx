package testutil

import (
	"context"
	"testing"

	"github.com/go-courier/logr"
	"github.com/go-courier/logr/slog"
)

func NewContext(t testing.TB) context.Context {
	t.Helper()
	ctx := context.Background()
	return logr.WithLogger(ctx, slog.Logger(slog.Default()))
}
