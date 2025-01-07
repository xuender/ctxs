package ctxs_test

import (
	"bytes"
	"context"
	"fmt"

	ctxs "github.com/xuender/ctx"
)

func ExampleCopy() {
	ctx, cancel := context.WithCancel(context.Background())
	dst := &bytes.Buffer{}
	src := bytes.NewBufferString("data")

	fmt.Println(ctxs.Copy(ctx, dst, src))

	cancel()

	fmt.Println(ctxs.Copy(ctx, dst, src))

	// Output:
	// 4 <nil>
	// 0 context canceled
}
