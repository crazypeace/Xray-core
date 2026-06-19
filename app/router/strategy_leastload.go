package router

import (
	"context"

	"github.com/xtls/xray-core/common/dice"
	"github.com/xtls/xray-core/common/errors"
)

// LeastLoadStrategy represents a least load balancing strategy
type LeastLoadStrategy struct {
	ctx      context.Context
	settings *StrategyLeastLoadConfig
}

func NewLeastLoadStrategy(settings *StrategyLeastLoadConfig) *LeastLoadStrategy {
	return &LeastLoadStrategy{settings: settings}
}

func (l *LeastLoadStrategy) GetPrincipleTarget(strings []string) []string {
	return strings
}

func (l *LeastLoadStrategy) InjectContext(ctx context.Context) {
	l.ctx = ctx
}

func (l *LeastLoadStrategy) PickOutbound(candidates []string) string {
	count := len(candidates)
	if count == 0 {
		return ""
	}
	errors.LogInfo(l.ctx, "observatory removed, falling back to random selection")
	return candidates[dice.Roll(count)]
}
