package ctxs

import (
	"context"
	"io"
)

// Copy transfers data from src to dst using the provided context for cancellation or timeout.
func Copy(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, NewReader(ctx, src))
}
