package router

import (
	"context"

	"github.com/xtls/xray-core/common/errors"
)

type LeastPingStrategy struct {
	ctx context.Context
}

func (l *LeastPingStrategy) GetPrincipleTarget(strings []string) []string {
	return []string{l.PickOutbound(strings)}
}

func (l *LeastPingStrategy) InjectContext(ctx context.Context) {
	l.ctx = ctx
}

func (l *LeastPingStrategy) PickOutbound(strings []string) string {
	errors.LogInfo(l.ctx, "observatory removed, falling back to first available")
	if len(strings) > 0 {
		return strings[0]
	}
	return ""
}
