package goLotteryTw

import (
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Lottery39 struct {
	LotterySession
}

func NewLottery39() *Lottery39 {
	newLottery39 := new(Lottery39)
	return newLottery39
}

func (l *Lottery39) GetInfo(s *goquery.Selection) LotteryInfo {
	var info LotteryInfo
	var strId string
	var strDate string
	var strBall1 string
	var strBall2 string
	var strBall3 string
	var strBall4 string
	var strBall5 string
	if s.Find("span[id^=\"M539Control_history1_dlQuery_M539_DrawTerm_\"]").Size() > 0 {
		strId = s.Find("span[id^=\"M539Control_history1_dlQuery_M539_DrawTerm_\"]").Eq(0).Text()
		strDate = s.Find("span[id^=\"M539Control_history1_dlQuery_M539_DDate_\"]").Eq(0).Text()
		strBall1 = s.Find("span[id^=\"M539Control_history1_dlQuery_SNo1_\"]").Eq(0).Text()
		strBall2 = s.Find("span[id^=\"M539Control_history1_dlQuery_SNo2_\"]").Eq(0).Text()
		strBall3 = s.Find("span[id^=\"M539Control_history1_dlQuery_SNo3_\"]").Eq(0).Text()
		strBall4 = s.Find("span[id^=\"M539Control_history1_dlQuery_SNo4_\"]").Eq(0).Text()
		strBall5 = s.Find("span[id^=\"M539Control_history1_dlQuery_SNo5_\"]").Eq(0).Text()
	} else {
		strId = s.Find("span[id^=\"M539Control_history1_dlQuery_Label2_\"]").Eq(0).Text()
		strDate = s.Find("span[id^=\"M539Control_history1_dlQuery_Label3_\"]").Eq(0).Text()
		strBall1 = s.Find("span[id^=\"M539Control_history1_dlQuery_Label4_\"]").Eq(0).Text()
		strBall2 = s.Find("span[id^=\"M539Control_history1_dlQuery_Label5_\"]").Eq(0).Text()
		strBall3 = s.Find("span[id^=\"M539Control_history1_dlQuery_Label6_\"]").Eq(0).Text()
		strBall4 = s.Find("span[id^=\"M539Control_history1_dlQuery_Label7_\"]").Eq(0).Text()
		strBall5 = s.Find("span[id^=\"M539Control_history1_dlQuery_Label8_\"]").Eq(0).Text()

	}
	intBall1, _ := strconv.ParseInt(strBall1, 10, 32)
	intBall2, _ := strconv.ParseInt(strBall2, 10, 32)
	intBall3, _ := strconv.ParseInt(strBall3, 10, 32)
	intBall4, _ := strconv.ParseInt(strBall4, 10, 32)
	intBall5, _ := strconv.ParseInt(strBall5, 10, 32)
	info.Id = strId
	info.Date = strDate
	info.Balls = append(info.Balls, intBall1, intBall2, intBall3, intBall4, intBall5)
	return info
}

func (l *Lottery39) GetUrl() string {
	return "http://www.taiwanlottery.com.tw/lotto/39M5/history.aspx"
}

func (l *Lottery39) GetIdValues(id string) url.Values {
	values := l.getDefaultValues()
	values.Set("M539Control_history1$DropDownList1", "10")
	values.Set("M539Control_history1$btnSubmit", "查詢")
	values.Set("M539Control_history1$chk", "radNO")
	values.Set("M539Control_history1$txtNO", id)
	return values
}
func (l *Lottery39) GetMonthValues(year string, month string) url.Values {
	values := l.getDefaultValues()
	values.Set("M539Control_history1$DropDownList1", "10")
	values.Set("M539Control_history1$btnSubmit", "查詢")
	values.Set("M539Control_history1$chk", "radYM")
	values.Set("M539Control_history1$dropYear", year)
	values.Set("M539Control_history1$dropMonth", month)
	return values
}
