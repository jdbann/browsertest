package browsertest

import (
	"testing"
)

func TestQueries(t *testing.T) {
	server := newTestServer(t)

	test := NewTest(t, server.URL)

	test.Run(
		test.Navigate("/queries"),
		test.WaitReady("body"),

		test.Text("div#textEquals").Equals("A div containing text"),
		test.Text("div#textEquals").Not().Equals("An li containing numbers"),
		test.Text("div#textContains").Contains("secretmessage"),
		test.Text("div#textContains").Not().Contains("coherent text"),

		test.InnerHTML("div#innerHTMLEquals").Equals(`<span class="is-hidden">A hidden message</span>`),
		test.InnerHTML("div#innerHTMLEquals").Not().Equals("<em>An exposed message</em>"),
		test.InnerHTML("div#innerHTMLContains").Contains("</em> messages</strong>"),
		test.InnerHTML("div#innerHTMLContains").Not().Contains("<strong>Important messages</strong>"),

		test.Value("input#valueEquals").Equals("simple message"),
		test.Value("input#valueEquals").Not().Equals("elaborate story"),
		test.Value("input#valueContains").Contains("codedmessage"),
		test.Value("input#valueContains").Not().Contains("obvious clue"),
	)
}
