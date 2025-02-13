package ctxs

import (
	"context"
	"io"
)

// Writer is a context-aware writer that cancels writing if the context is done.

type Writer struct {
	ctx    context.Context // nolint: containedctx
	writer io.Writer
}

// NewWriter creates a new context-aware writer.
func NewWriter(ctx context.Context, writer io.Writer) *Writer {
	return &Writer{
		ctx:    ctx,
		writer: writer,
	}
}

// Write writes data from buf, stopping if the context is canceled.
func (r *Writer) Write(buf []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}

	return r.writer.Write(buf)
}
