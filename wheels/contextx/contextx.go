package contextx

import (
	"context"

	"github.com/v3nooonn/gheels/tools/constant"
)

type Helper interface {
	WithValue(k constant.ContextKey, v any) Helper
	Context() context.Context
}

type ContextHelper struct {
	context context.Context
}

func WithContext(ctx context.Context) Helper {
	return &ContextHelper{context: ctx}
}

func (h *ContextHelper) WithValue(k constant.ContextKey, v any) Helper {
	h.setContext(context.WithValue(h.Context(), k, v))
	return h
}

func (h *ContextHelper) Context() context.Context {
	return h.context
}

func (h *ContextHelper) setContext(ctx context.Context) *ContextHelper {
	h.context = ctx
	return h
}
