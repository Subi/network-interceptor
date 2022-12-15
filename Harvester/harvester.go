package harvester

import (
	model "network/Model"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
)

type Harvester interface {
	Hijack()
}

type harvester struct {
	URL     string
	Method  string
	Headers model.ShapeHeaders
	Page    *rod.Page
	Browser *rod.Browser
}

func NewHarvester(url string, method string) Harvester {
	return &harvester{URL: url, Method: method}
}

func newBrowser() *rod.Browser {
	var browser *rod.Browser
	browser = rod.New().MustConnect()

	return browser
}

func newPage(browser *rod.Browser) *rod.Page {
	page := stealth.MustPage(browser)

	return page
}

func (h *harvester) initializeHijacking() {
	router := h.Page.HijackRequests()

	router.MustAdd("*", func(ctx *rod.Hijack) {
		if ctx.Request.Method() == "GET" {
			if ctx.Request.Type() == proto.NetworkResourceTypeImage || ctx.Request.Type() == proto.NetworkResourceTypeStylesheet {
				ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			}

			ctx.ContinueRequest(&proto.FetchContinueRequest{})
		}


		if strings.Contains(ctx.Request.URL().Path , "")
	})



}

func (h *harvester) Hijack() {
	h.Browser = newBrowser()
	h.Page = newPage(h.Browser)
	h.Page.MustNavigate(h.URL).MustWaitLoad()

	h.Page.MustScreenshot("harvester.png")
}
