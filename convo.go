package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Donnie/Airstrip/ptr"
)

// Handle handles a conversation by mapping expectors
// with contextual answer
func (convo *Convo) Handle(expect string, handler Expector) {
	convo.handlers[expect] = handler
}

func (convo *Convo) expectNext(gl *Global, expect string) string {
	if convo.Expect != nil {
		convo.handlers[*convo.Expect](gl, expect)
	}

	if convo.Expect != nil {
		gl.Orm.Save(&convo)
		return genQues(*convo.Expect)
	}
	gl.Orm.Delete(&convo)
	return "Record stored!"
}

func (convo *Convo) expectAccount(gl *Global, input string) {
	gl.Orm.Model(&Record{}).Where("id = ?", *convo.ContextID).Update("account", input)
	convo.Expect = ptr.String("amount")
}

func (convo *Convo) expectAmount(gl *Global, input string) {
	amountFlt, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return
	}
	gl.Orm.Model(&Record{}).Where("id = ?", *convo.ContextID).Update("amount", ptr.Int64(int64(amountFlt*100)))
	convo.Expect = ptr.String("currency")
}

func (convo *Convo) expectCurrency(gl *Global, input string) {
	currency := input
	if len(currency) != 3 {
		currency = "EUR"
	}
	gl.Orm.Model(&Record{}).Where("id = ?", *convo.ContextID).Update("currency", currency)
	convo.Expect = ptr.String("description")
}

func (convo *Convo) expectDescription(gl *Global, input string) {
	record := &Record{}
	gl.Orm.First(&record, *convo.ContextID)
	record.Description = &input
	gl.Orm.Save(&record)
	if *record.Type == "variable" {
		convo.Expect = ptr.String("date")
	} else {
		convo.Expect = ptr.String("from date")
	}
}

func (convo *Convo) expectDate(gl *Global, input string) {
	date := input
	layout := "2006-01-02 15:04"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		dateTime = time.Now()
	}
	gl.Orm.Model(&Record{}).Where("id = ?", *convo.ContextID).Update("date", dateTime)
	convo.Expect = nil
}

func (convo *Convo) expectFromDate(gl *Global, input string) {
	date := input
	layout := "Jan 2006"
	dateTime, _ := time.Parse(layout, date)
	gl.Orm.Model(&Record{}).Where("id = ?", *convo.ContextID).Update("from_date", dateTime)
	convo.Expect = ptr.String("till date")
}

func (convo *Convo) expectTillDate(gl *Global, input string) {
	record := &Record{}
	gl.Orm.First(&record, *convo.ContextID)
	date := input
	layout := "Jan 2006"
	dateTime, err := time.Parse(layout, date)
	if err == nil {
		dateTime = dateTime.AddDate(0, 1, -1)
		record.TillDate = &dateTime
		gl.Orm.Save(&record)
	}
	record.ID = nil
	if *record.Form == "lend" {
		record.Form = ptr.String("expense")
		gl.Orm.Create(&record)
	} else if *record.Form == "loan" {
		record.Form = ptr.String("gain")
		gl.Orm.Create(&record)
	}
	convo.Expect = nil
}

func genQues(ask string) string {
	return fmt.Sprintf("What is the %s?", ask)
}