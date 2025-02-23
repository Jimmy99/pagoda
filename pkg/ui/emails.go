package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ConfirmEmailAddressEmail(username, url string) Node {
	return Group{
		Strong(Textf("Hello %s,", username)),
		Br(),
		P(Text("Please click on the following link to confirm your email address:")),
		Br(),
		A(Href(url), Text(url)),
	}
}
