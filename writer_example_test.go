package ctxs_test

import (
	"bytes"
	"context"
	"fmt"

	"github.com/xuender/ctxs"
)

func ExampleWriter_Write() {
	dst := &bytes.Buffer{}
	ctx, cancel := context.WithCancel(context.Background())
	writer := ctxs.NewWriter(ctx, dst)

	fmt.Println(writer.Write([]byte("xx")))
	cancel()
	fmt.Println(writer.Write([]byte("xx")))

	// Output:
	// 2 <nil>
	// 0 context canceled
}
