package browsertest

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

type Test struct {
	*testing.T
	ctx     context.Context
	timeout time.Duration
	baseURL string
}

func NewTest(t *testing.T, baseURL string) Test {
	return Test{
		t,
		context.Background(),
		time.Second * 2,
		baseURL,
	}
}

func (bt Test) Run(actions ...Action) {
	ctx, cancel := chromedp.NewContext(bt.ctx)
	defer cancel()

	bt.executeAction(ctx, actions[0])

	for _, a := range actions[1:] {
		timeout := bt.timeout

		if awt, ok := a.(interface{ Timeout() time.Duration }); ok {
			timeout = awt.Timeout()
		}
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		bt.executeAction(ctx, a)
	}
}
