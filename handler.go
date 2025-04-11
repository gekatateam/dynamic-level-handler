package logger

import (
	"context"
	"log/slog"
	"math"
)

var _ slog.Handler = (*DynamicLevelHandler)(nil)

const LevelUnassigned slog.Level = math.MaxInt

func OverrideLevel(h slog.Handler, newLevel slog.Leveler) {
	if dlh, ok := h.(*DynamicLevelHandler); ok {
		if newLevel != nil {
			dlh.Override(newLevel)
		}
	}
}

func New(h slog.Handler) *DynamicLevelHandler {
	return &DynamicLevelHandler{
		basic:         h,
		assignedLevel: LevelUnassigned,
	}
}

type DynamicLevelHandler struct {
	basic         slog.Handler
	assignedLevel slog.Leveler
}

func (h *DynamicLevelHandler) Override(newLevel slog.Leveler) {
	h.assignedLevel = newLevel
}

func (h *DynamicLevelHandler) Handle(ctx context.Context, record slog.Record) error {
	return h.basic.Handle(ctx, record)
}

func (h *DynamicLevelHandler) Enabled(ctx context.Context, level slog.Level) bool {
	if h.assignedLevel.Level() == LevelUnassigned {
		return h.basic.Enabled(ctx, level)
	}

	return level >= h.assignedLevel.Level()
}

func (h *DynamicLevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &DynamicLevelHandler{
		basic:         h.basic.WithAttrs(attrs),
		assignedLevel: h.assignedLevel,
	}
}

func (h *DynamicLevelHandler) WithGroup(name string) slog.Handler {
	return &DynamicLevelHandler{
		basic:         h.basic.WithGroup(name),
		assignedLevel: h.assignedLevel,
	}
}
