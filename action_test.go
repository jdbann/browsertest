package browsertest

import (
	"testing"
)

func TestActions(t *testing.T) {
	server := newTestServer(t)

	test := NewTest(t, server.URL)

	test.Run(
		test.Navigate("/actions"),
		test.WaitReady("body"),
		test.Click(`button[data-action="show-hidden-message"]`),
		test.WaitVisible(`div[data-target="hidden-message"]`),
		test.Click(`button[data-action="hide-hidden-message"]`),
		test.WaitNotVisible(`div[data-target="hidden-message"]`),
	)
}
