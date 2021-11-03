package browsertest

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

type Action struct {
	chromedp.Action
	msg string
}

func (bt Test) executeAction(ctx context.Context, action Action) {
	bt.Log(action.msg)
	if err := chromedp.Run(ctx, action); err != nil {
		bt.Fatalf("%s: %s", action.msg, err)
	}
}

func (bt Test) Run(actions ...Action) {
	ctx, cancel := chromedp.NewContext(bt.ctx)
	defer cancel()

	bt.executeAction(ctx, actions[0])

	for _, a := range actions[1:] {
		ctx, cancel := context.WithTimeout(ctx, bt.timeout)
		defer cancel()

		bt.executeAction(ctx, a)
	}
}

func (bt Test) Navigate(url string) Action {
	return Action{
		chromedp.Navigate(bt.baseURL + url),
		fmt.Sprintf("[Navigate] %q", url),
	}
}

func (bt Test) Click(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.Click(sel, opts...),
		fmt.Sprintf("[Click] %v", sel),
	}
}

func (bt Test) WaitReady(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.WaitReady(sel, opts...),
		fmt.Sprintf("[WaitReady] %v", sel),
	}
}

func (bt Test) WaitVisible(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.WaitVisible(sel, opts...),
		fmt.Sprintf("[WaitVisible] %v", sel),
	}
}

func (bt Test) WaitNotVisible(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.WaitNotVisible(sel, opts...),
		fmt.Sprintf("[WaitNotVisible] %v", sel),
	}
}

func (bt Test) WaitEnabled(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.WaitEnabled(sel, opts...),
		fmt.Sprintf("[WaitEnabled] %v", sel),
	}
}

func (bt Test) WaitSelected(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.WaitSelected(sel, opts...),
		fmt.Sprintf("[WaitSelected] %v", sel),
	}
}

func (bt Test) WaitNotPresent(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.WaitNotPresent(sel, opts...),
		fmt.Sprintf("[WaitNotPresent] %v", sel),
	}
}

func (bt Test) SendKeys(sel interface{}, v string, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.SendKeys(sel, v, opts...),
		fmt.Sprintf("[SendKeys] %v %q", sel, v),
	}
}

func (bt Test) Submit(sel interface{}, opts ...chromedp.QueryOption) Action {
	return Action{
		chromedp.Submit(sel, opts...),
		fmt.Sprintf("[Submit] %v", sel),
	}
}

func (bt Test) ActionFunc(f chromedp.ActionFunc, msg string) Action {
	return Action{f, msg}
}

func (bt Test) Poll(expression string, opts ...chromedp.PollOption) Action {
	return Action{
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
