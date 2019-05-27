package srb_test

import (
	"testing"
	"time"

	"github.com/mileusna/srb"
)

/*
   TestCyr tests gocyr package
*/
func TestCyr(t *testing.T) {

	var exp, res, param string

	param = "кашњење у плаћању пореза"
	if !srb.HasCyr(param) {
		t.Errorf("HasCyr() returns false for string `%s`\n", param)
	}

	param = "nema ćirilice ovde"
	if srb.HasCyr(param) {
		t.Errorf("HasCyr() returns true for string `%s`\n", param)
	}

	exp = "Novak Đoković"
	res = srb.ToLat("Новак Ђоковић")
	if res != exp {
		t.Errorf("ToLat() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "грађевинске дозволе"
	res = srb.ToCyr("građevinske dozvole")
	if res != exp {
		t.Errorf("ToCyr() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "ексцлусив ББЦ невс дисцовери"
	res = srb.ToCyr("exclusiv BBC news discovery")
	if res != exp {
		t.Errorf("ToCyr() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "Novak Đoković"
	res = srb.FixDj("Novak Djoković")
	if res != exp {
		t.Errorf("FixDj() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "<h1><a href='http://www.naslovi.net/'>Вести дана - Наслови.нет</a></h1>"
	res = srb.HTMLToCyr("<h1><a href='http://www.naslovi.net/'>Vesti dana - Naslovi.net</a></h1>")
	if res != exp {
		t.Errorf("HTMLToCyr() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "Milos Djordjevic Zarkov cukundeda"
	res = srb.ToASCII("Miloš Đorđević Žarkov cukundeda")
	if res != exp {
		t.Errorf("ToASCII() result `%s` doesn't match expected result `%s`\n", res, exp)
	}
}

type testData struct {
	date    string
	month   string
	weekday string
}

func TestDates(t *testing.T) {

	var testData = []testData{
		testData{
			date:    "2019-02-02",
			month:   "februar",
			weekday: "subota",
		},
		testData{
			date:    "2019-02-03",
			month:   "februar",
			weekday: "nedelja",
		},
		testData{
			date:    "2018-07-11",
			month:   "jul",
			weekday: "sreda",
		},
		testData{
			date:    "2018-08-14",
			month:   "avgust",
			weekday: "utorak",
		},
		testData{
			date:    "2018-09-14",
			month:   "septembar",
			weekday: "petak",
		},
		testData{
			date:    "2018-10-09",
			month:   "oktobar",
			weekday: "subota",
		},
		testData{
			date:    "2018-11-19",
			month:   "novembar",
			weekday: "ponedeljak",
		},
		testData{
			date:    "2018-12-30",
			month:   "decembar",
			weekday: "nedelja",
		},
		testData{
			date:    "2018-01-04",
			month:   "januar",
			weekday: "četvrtak",
		},
		testData{
			date:    "2018-02-14",
			month:   "februar",
			weekday: "sreda",
		},
		testData{
			date:    "2018-03-22",
			month:   "mart",
			weekday: "četvrtak",
		},
		testData{
			date:    "2018-04-22",
			month:   "april",
			weekday: "nedelja",
		},
		testData{
			date:    "2018-05-26",
			month:   "maj",
			weekday: "subota",
		},
		testData{
			date:    "2018-06-19",
			month:   "jun",
			weekday: "utorak",
		},
	}

	for _, td := range testData {
		d, _ := time.Parse("2006-01-02", td.date)
		if srb.Month(d) != td.month {
			t.Errorf("Month should be %s not %s", td.month, srb.Month(d))
		}
	}
}
