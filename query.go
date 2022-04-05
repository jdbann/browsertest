package browsertest

import (
	"context"
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
)

type Query struct {
	queryFunc      func(context.Context) (string, error)
	msg            string
	expectedResult bool
}

type queryActionFunc func(interface{}, *string, ...chromedp.QueryOption) chromedp.QueryAction

func newQuery(msg string, query queryActionFunc, sel interface{}, opts ...chromedp.QueryOption) Query {
	return Query{
		queryFunc: func(ctx context.Context) (string, error) {
			var queryResult string
			if err := query(sel, &queryResult, opts...).Do(ctx); err != nil {
				return "", err
			}

			return queryResult, nil
		},
		msg:            msg,
		expectedResult: true,
	}
}

func (bq Query) Contains(expected string) Action {
	return BasicAction{
		chromedp.ActionFunc(func(ctx context.Context) error {
			actual, err := bq.queryFunc(ctx)
			if err != nil {
				return err
			}

			if strings.Contains(actual, expected) != bq.expectedResult {
				return fmt.Errorf("expected %q to contain %q", actual, expected)
			}

			return nil
		}),
		fmt.Sprintf("%s [Contains] %q", bq.msg, expected),
	}
}

func (bq Query) Equals(expected string) Action {
	return BasicAction{
		chromedp.ActionFunc(func(ctx context.Context) error {
			actual, err := bq.queryFunc(ctx)
			if err != nil {
				return err
			}

			if (strings.TrimSpace(actual) == expected) != bq.expectedResult {
				return fmt.Errorf("expected %q to equal %q", expected, actual)
			}

			return nil
		}),
		fmt.Sprintf("%s [Equals] %q", bq.msg, expected),
	}
}

func (bq Query) Not() Query {
	bq.expectedResult = false
	return bq
}

func (bt Test) Text(sel interface{}, opts ...chromedp.QueryOption) Query {
	return newQuery(fmt.Sprintf("[Text] %q", sel), chromedp.Text, sel, opts...)
}

func (bt Test) InnerHTML(sel interface{}, opts ...chromedp.QueryOption) Query {
	return newQuery(fmt.Sprintf("[InnerHTML] %q", sel), chromedp.InnerHTML, sel, opts...)
}

func (bt Test) Value(sel interface{}, opts ...chromedp.QueryOption) Query {
	return newQuery(fmt.Sprintf("[Value] %q", sel), chromedp.Value, sel, opts...)
}
