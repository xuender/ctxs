package ctxs_test

import (
	"bytes"
	"context"
	"fmt"

	ctxs "github.com/xuender/ctx"
)

func ExampleWriter_Writer() {
	dst := &bytes.Buffer{}
	ctx, cancel := context.WithCancel(context.Background())
	writer := ctxs.NewWriter(ctx, dst)

	fmt.Println(writer.Writer([]byte("xx")))
	cancel()
	fmt.Println(writer.Writer([]byte("xx")))

	// Output:
	// 2 <nil>
	// 0 context canceled
}
