package goLotteryTw

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type LotterySession struct {
	lastEventTarget        string
	lastEventArgument      string
	lastLastFocus          string
	lastViewState          string
	lastViewStateGenerator string
	lastViewStateEncrypted string
	lastEventValidation    string
}

func (ls *LotterySession) getDefaultValues() url.Values {
	values := make(url.Values)
	values.Set("__EVENTTARGET", ls.lastEventTarget)
	values.Set("__EVENTARGUMENT", ls.lastEventArgument)
	values.Set("__LASTFOCUS", ls.lastLastFocus)
	values.Set("__VIEWSTATE", ls.lastViewState)
	values.Set("__VIEWSTATEGENERATOR", ls.lastViewStateGenerator)
	values.Set("__VIEWSTATEENCRYPTED", ls.lastViewStateEncrypted)
	values.Set("__EVENTVALIDATION", ls.lastEventValidation)
	return values
}

func (ls *LotterySession) RefreshForm(doc *goquery.Document) {
	newEventTarget, existEventTarget := doc.Find("#__EVENTTARGET").Attr("value")
	newEventArgument, existEventArgument := doc.Find("#__EVENTARGUMENT").Attr("value")
	newLastFocus, existLastFocus := doc.Find("#__LASTFOCUS").Attr("value")
	newViewState, existViewState := doc.Find("#__VIEWSTATE").Attr("value")
	newViewStateGenerator, existViewStateGenerator := doc.Find("#__VIEWSTATEGENERATOR").Attr("value")
	newViewStateEncrypted, existViewStateEncrypted := doc.Find("#__VIEWSTATEENCRYPTED").Attr("value")
	newEventValidation, existEventValidation := doc.Find("#__EVENTVALIDATION").Attr("value")
	if existEventTarget {
		ls.lastEventTarget = newEventTarget
	}
	if existEventArgument {
		ls.lastEventArgument = newEventArgument
	}
	if existLastFocus {
		ls.lastLastFocus = newLastFocus
	}
	if existViewState {
		ls.lastViewState = newViewState
	}
	if existViewStateGenerator {
		ls.lastViewStateGenerator = newViewStateGenerator
	}
	if existViewStateEncrypted {
		ls.lastViewStateEncrypted = newViewStateEncrypted
	}
	if existEventValidation {
		ls.lastEventValidation = newEventValidation
	}
}
