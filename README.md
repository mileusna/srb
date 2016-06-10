# gocyr
[![GoDoc](https://godoc.org/github.com/mileusna/gocyr?status.svg)](https://godoc.org/github.com/mileusna/gocyr)

Converts string to/from Cyrillic/Latin for Serbian language, including some other utility functions for Serbian script.

## Examples

```Go
package main

import (
	"fmt"
	"github.com/mileusna/gocyr"
)

func main() {
    fmt.Println(gocyr.ToCyr("Kragujevac"))
    fmt.Println(gocyr.ToLat("кошарка"))
	fmt.Println(gocyr.HasCyr("Ima li ћирилице u ovom stringu?"))
    fmt.Println(gocyr.FixDj("Novak Djoković"))
    fmt.Println(gocyr.HTMLToCyr("<h1><a href='http://naslovi.net/'>Vesti dana</a></h1>"))
    fmt.Println(gocyr.ToASCII("Miloš"))
}
```
