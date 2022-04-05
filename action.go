package browsertest

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

type Action = interface {
	chromedp.Action
	Msg() string
}

func (bt Test) executeAction(ctx context.Context, action Action) {
	bt.Log(action.Msg())
	if err := chromedp.Run(ctx, action); err != nil {
		bt.Fatalf("%s: %s", action.Msg(), err)
	}
}

type BasicAction struct {
	chromedp.Action
	msg string
}

func (ba BasicAction) Msg() string {
	return ba.msg
}

func (bt Test) Navigate(url string) Action {
	return BasicAction{
		chromedp.Navigate(bt.baseURL + url),
		fmt.Sprintf("[Navigate] %q", url),
	}
}

func (bt Test) Click(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.Click(sel, opts...),
		fmt.Sprintf("[Click] %v", sel),
	}
}

func (bt Test) WaitReady(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.WaitReady(sel, opts...),
		fmt.Sprintf("[WaitReady] %v", sel),
	}
}

func (bt Test) WaitVisible(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.WaitVisible(sel, opts...),
		fmt.Sprintf("[WaitVisible] %v", sel),
	}
}

func (bt Test) WaitNotVisible(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.WaitNotVisible(sel, opts...),
		fmt.Sprintf("[WaitNotVisible] %v", sel),
	}
}

func (bt Test) WaitEnabled(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.WaitEnabled(sel, opts...),
		fmt.Sprintf("[WaitEnabled] %v", sel),
	}
}

func (bt Test) WaitSelected(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.WaitSelected(sel, opts...),
		fmt.Sprintf("[WaitSelected] %v", sel),
	}
}

func (bt Test) WaitNotPresent(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.WaitNotPresent(sel, opts...),
		fmt.Sprintf("[WaitNotPresent] %v", sel),
	}
}

func (bt Test) SendKeys(sel interface{}, v string, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.SendKeys(sel, v, opts...),
		fmt.Sprintf("[SendKeys] %v %q", sel, v),
	}
}

func (bt Test) Submit(sel interface{}, opts ...chromedp.QueryOption) Action {
	return BasicAction{
		chromedp.Submit(sel, opts...),
		fmt.Sprintf("[Submit] %v", sel),
	}
}

func (bt Test) ActionFunc(f chromedp.ActionFunc, msg string) Action {
	return BasicAction{f, msg}
}

func (bt Test) Poll(expression string, opts ...chromedp.PollOption) Action {
	return BasicAction{
		chromedp.ActionFunc(func(ctx context.Context) error {
			var result bool

			if err := chromedp.Run(ctx, chromedp.Poll(expression, &result, opts...)); err != nil {
				return err
			}

			if result == false {
				// I don't think this should be possible
				bt.Fatal("[Poll] Attempt to wait returned false")
			}

			return nil
		}),
		fmt.Sprintf("[Poll] %s", expression),
	}
}

type ActionWithTimeout struct {
	Action
	timeout time.Duration
}

func (awt ActionWithTimeout) Timeout() time.Duration {
	return awt.timeout
}

func (bt Test) WithTimeout(timeout time.Duration, action Action) Action {
	return ActionWithTimeout{
		action,
		timeout,
	}
}
