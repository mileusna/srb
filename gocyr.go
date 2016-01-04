/*
 */
package gocyr

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	cyr          = "ЕРТЗУИОПШЂАСДФГХЈКЛЧЋЦВБНМЖертзуиопшђасдфгхјклчћцвбнмж"
	toLat        = strings.NewReplacer("Е", "E", "Р", "R", "Т", "T", "З", "Z", "У", "U", "И", "I", "О", "O", "П", "P", "Ш", "Š", "Ђ", "Đ", "А", "A", "С", "S", "Д", "D", "Ф", "F", "Г", "G", "Х", "H", "Ј", "J", "К", "K", "Л", "L", "Ч", "Č", "Ћ", "Ć", "Ц", "C", "В", "V", "Б", "B", "Н", "N", "М", "M", "Ж", "Ž", "е", "e", "р", "r", "т", "t", "з", "z", "у", "u", "и", "i", "о", "o", "п", "p", "ш", "š", "ђ", "đ", "а", "a", "с", "s", "д", "d", "ф", "f", "г", "g", "х", "h", "ј", "j", "к", "k", "л", "l", "ч", "č", "ћ", "ć", "ц", "c", "в", "v", "б", "b", "н", "n", "м", "m", "ж", "ž", "Љ", "Lj", "Њ", "Nj", "Џ", "Dž", "љ", "lj", "њ", "nj", "џ", "dž")
	toCyr        = strings.NewReplacer("Lj", "Љ", "LJ", "Љ", "Nj", "Њ", "NJ", "Њ", "Dž", "Џ", "DŽ", "Џ", "lj", "љ", "nj", "њ", "dž", "џ", "E", "Е", "R", "Р", "T", "Т", "Z", "З", "U", "У", "I", "И", "O", "О", "P", "П", "Š", "Ш", "Đ", "Ђ", "A", "А", "S", "С", "D", "Д", "F", "Ф", "G", "Г", "H", "Х", "J", "Ј", "K", "К", "L", "Л", "Č", "Ч", "Ć", "Ћ", "C", "Ц", "V", "В", "B", "Б", "N", "Н", "M", "М", "Ž", "Ж", "e", "е", "r", "р", "t", "т", "z", "з", "u", "у", "i", "и", "o", "о", "p", "п", "š", "ш", "đ", "ђ", "a", "а", "s", "с", "d", "д", "f", "ф", "g", "г", "h", "х", "j", "ј", "k", "к", "l", "л", "č", "ч", "ć", "ћ", "c", "ц", "v", "в", "b", "б", "n", "н", "m", "м", "ž", "ж")
	fixDj        = strings.NewReplacer("Dj", "Đ", "DJ", "Đ", "dj", "đ")
	htmlRegex, _ = regexp.Compile("(^|>)([^<]*)(<|$)")
)

/*
   ToLat converts string from cyrillic to latin
*/
func ToLat(s string) string {
	return toLat.Replace(s)
}

/*
   ToCyr converts string from latin to cyrillic
*/
func ToCyr(s string) string {
	return toCyr.Replace(s)
}

/*
   FixDj replaces Dj with Đ, uppercase and lowercase
*/
func FixDj(s string) string {
	return fixDj.Replace(s)
}

/*
   IsCyr returns true if string contains at least one cyrillic character
*/
func IsCyr(s string) bool {
	return (strings.IndexAny(s, cyr) != -1)
}

/*
   HtmlToCyr converts string that contains HTML tags from latin to cyrillic, preserving HTML tags in latin
*/
func HtmlToCyr(html string) string {
	return htmlRegex.ReplaceAllStringFunc(html, ToCyr)
}

/*
   HtmlToLat converts string that contains html tags from cyrillic to latin
   Not needed actually
*/
// func HtmlToLat(html string) string {
// 	return htmlRegex.ReplaceAllStringFunc(html, ToLat)
// }

func dummy() {
	lat := []rune("ERTZUIOPŠĐASDFGHJKLČĆCVBNMŽertzuiopšđasdfghjklčćcvbnmž")
	cyr := []rune("ЕРТЗУИОПШЂАСДФГХЈКЛЧЋЦВБНМЖертзуиопшђасдфгхјклчћцвбнмж")
	for i := 0; i < len(lat); i++ {
		fmt.Printf("\"%s\", \"%s\", ", string(lat[i]), string(cyr[i]))

	}
}
