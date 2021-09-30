package browsertest

import (
	"context"
	"testing"
	"time"
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
