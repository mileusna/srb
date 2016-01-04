package gocyr_test

import (
	"testing"

	"github.com/mileusna/gocyr"
)

/*
   TestCyr tests gocyr package
*/
func TestCyr(t *testing.T) {

	var exp, res, param string

	param = "кашњење у плаћању пореза"
	if !gocyr.IsCyr(param) {
		t.Errorf("IsCyr() returns false for string `%s`\n", param)
	}

	param = "nema ćirilice ovde"
	if gocyr.IsCyr(param) {
		t.Errorf("IsCyr() returns true for string `%s`\n", param)
	}

	exp = "Novak Đoković"
	res = gocyr.ToLat("Новак Ђоковић")
	if res != exp {
		t.Errorf("ToLat() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "грађевинске дозволе"
	res = gocyr.ToCyr("građevinske dozvole")
	if res != exp {
		t.Errorf("ToCyr() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "Novak Đoković"
	res = gocyr.FixDj("Novak Djoković")
	if res != exp {
		t.Errorf("FixDj() result `%s` doesn't match expected result `%s`\n", res, exp)
	}

	exp = "<h1><a href='http://www.naslovi.net/'>Вести дана - Наслови.нет</a></h1>"
	res = gocyr.HtmlToCyr("<h1><a href='http://www.naslovi.net/'>Vesti dana - Naslovi.net</a></h1>")
	if res != exp {
		t.Errorf("HtmlToCyr() result `%s` doesn't match expected result `%s`\n", res, exp)
	}
}
