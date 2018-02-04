package goLotteryTw

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type LotteryInfo struct {
	Id     string
	Date   string
	Balls  []int64
	SBalls int
}

type LotterySource interface {
	GetInfo(s *goquery.Selection) LotteryInfo
	GetUrl() string
	GetIdValues(id string) url.Values
	GetMonthValues(year string, month string) url.Values
	RefreshForm(doc *goquery.Document)
}

var lotterySources = map[string]LotterySource{
	"38": NewLottery38(),
	"39": NewLottery39(),
	"49": NewLottery49(),
}

func QueryLast(game string) []LotteryInfo {
	source := getGameSource(game)
	if source == nil {
		return nil
	}

	response, _ := http.Get(source.GetUrl())

	doc, _ := goquery.NewDocumentFromReader(response.Body)
	source.RefreshForm(doc)

	s := doc.Find(".td_hm")
	ret := []LotteryInfo{}
	for i := 0; i < s.Size(); i += 1 {
		ret = append(ret, source.GetInfo(s.Eq(i)))
	}
	return ret
}

func QueryById(game string, id string) LotteryInfo {
	source := getGameSource(game)
	if source == nil {
		var info LotteryInfo
		return info
	}

	response, _ := http.PostForm(source.GetUrl(), source.GetIdValues(id))
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	if doc.Find(".td_hm").Size() == 0 {
		QueryLast(game)
		response, _ = http.PostForm(source.GetUrl(), source.GetIdValues(id))
		doc, _ = goquery.NewDocumentFromReader(response.Body)
	}
	source.RefreshForm(doc)

	return source.GetInfo(doc.Find(".td_hm"))
}

func QueryByMonth(game string, year string, month string) []LotteryInfo {
	source := getGameSource(game)
	if source == nil {
		return nil
	}

	response, _ := http.PostForm(source.GetUrl(), source.GetMonthValues(year, month))
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	s := doc.Find(".td_hm")
	if s.Size() == 0 {
		QueryLast(game)
		response, _ = http.PostForm(source.GetUrl(), source.GetMonthValues(year, month))
		doc, _ = goquery.NewDocumentFromReader(response.Body)
		s = doc.Find(".td_hm")
	}
	source.RefreshForm(doc)

	ret := []LotteryInfo{}
	for i := 0; i < s.Size(); i += 1 {
		ret = append(ret, source.GetInfo(s.Eq(i)))
	}
	return ret
}

func getGameSource(game string) LotterySource {
	return lotterySources[game]
}
