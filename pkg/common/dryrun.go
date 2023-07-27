package common

import (
	"context"
	"fmt"
)

type dryrunContextKey string

const dryrunContextKeyVal = dryrunContextKey("dryrun")

// Dryrun returns true if the current context is dryrun
func Dryrun(ctx context.Context) bool {
	val := ctx.Value(dryrunContextKeyVal)
	if val != nil {
		if dryrun, ok := val.(bool); ok {
			return dryrun
		}
	}
	return false
}

// WithDryrun adds a value to the context for dryrun
func WithDryrun(ctx context.Context, dryrun bool) context.Context {
	fmt.Println("WithDryrun dryrun ", dryrun)
	return context.WithValue(ctx, dryrunContextKeyVal, dryrun)
}
