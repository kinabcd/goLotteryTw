package goLotteryTw

import (
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Lottery38 struct {
	LotterySession
}

func NewLottery38() *Lottery38 {
	newLottery38 := new(Lottery38)
	return newLottery38
}

func (l *Lottery38) GetInfo(s *goquery.Selection) LotteryInfo {
	var info LotteryInfo
	strId := s.Find("span[id^=\"M638Control_history1_dlQuery_lbM638_DrawTerm_\"]").Eq(0).Text()
	strDate := s.Find("span[id^=\"M638Control_history1_dlQuery_M638_DDate_\"]").Eq(0).Text()
	strBall1 := s.Find("span[id^=\"M638Control_history1_dlQuery_SNo1_\"]").Eq(0).Text()
	strBall2 := s.Find("span[id^=\"M638Control_history1_dlQuery_SNo2_\"]").Eq(0).Text()
	strBall3 := s.Find("span[id^=\"M638Control_history1_dlQuery_SNo3_\"]").Eq(0).Text()
	strBall4 := s.Find("span[id^=\"M638Control_history1_dlQuery_SNo4_\"]").Eq(0).Text()
	strBall5 := s.Find("span[id^=\"M638Control_history1_dlQuery_SNo5_\"]").Eq(0).Text()
	strBall6 := s.Find("span[id^=\"M638Control_history1_dlQuery_SNo6_\"]").Eq(0).Text()
	intBall1, _ := strconv.ParseInt(strBall1, 10, 32)
	intBall2, _ := strconv.ParseInt(strBall2, 10, 32)
	intBall3, _ := strconv.ParseInt(strBall3, 10, 32)
	intBall4, _ := strconv.ParseInt(strBall4, 10, 32)
	intBall5, _ := strconv.ParseInt(strBall5, 10, 32)
	intBall6, _ := strconv.ParseInt(strBall6, 10, 32)
	info.Id = strId
	info.Date = strDate
	info.Balls = append(info.Balls, intBall1, intBall2, intBall3, intBall4, intBall5, intBall6)
	return info
}

func (l *Lottery38) GetUrl() string {
	return "http://www.taiwanlottery.com.tw/lotto/38M6/history.aspx"
}

func (l *Lottery38) GetIdValues(id string) url.Values {
	values := l.getDefaultValues()
	values.Set("M638Control_history1$DropDownList1", "4")
	values.Set("M638Control_history1$btnSubmit", "查詢")
	values.Set("M638Control_history1$chk", "radNO")
	values.Set("M638Control_history1$txtNO", id)
	return values
}

func (l *Lottery38) GetMonthValues(year string, month string) url.Values {
	values := l.getDefaultValues()
	values.Set("M638Control_history1$DropDownList1", "4")
	values.Set("M638Control_history1$btnSubmit", "查詢")
	values.Set("M638Control_history1$chk", "radYM")
	values.Set("M638Control_history1$dropYear", year)
	values.Set("M638Control_history1$dropMonth", month)
	return values
}
