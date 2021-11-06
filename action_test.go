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
		test.SendKeys(`select[data-action="enable-reset-select"]`, "On"),
		test.WaitEnabled(`button[data-target="reset-select-button"]`),
		test.Click(`button[data-action="reset-select"]`),
		test.WaitSelected(`option[value="off"]`),
		test.Click(`button[data-action="remove"]`),
		test.WaitNotPresent(`button[data-target="to-remove"]`),
		test.SendKeys(`input[name="value"]`, "An interesting piece of information."),
		test.Submit(`input[name="value"]`),
		test.Text(`span[data-target="submitted-value"]`).Equals("An interesting piece of information."),
	)
}
