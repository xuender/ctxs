package ctxs

import (
	"context"
	"io"
)

// Reader is a context-aware reader that cancels reading if the context is done.
type Reader struct {
	ctx    context.Context // nolint: containedctx
	reader io.Reader
}

// NewReader creates a new context-aware reader.
func NewReader(ctx context.Context, reader io.Reader) *Reader {
	return &Reader{
		ctx:    ctx,
		reader: reader,
	}
}

// Read reads data into buf, stopping if the context is canceled.
func (r *Reader) Read(buf []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}

	return r.reader.Read(buf)
}
