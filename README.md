# srb [![GoDoc](https://godoc.org/github.com/mileusna/gocyr?status.svg)](https://godoc.org/github.com/mileusna/gocyr)

Package srb converts string to/from Cyrillic/Latin for Serbian language, including some other utility functions for Serbian script.

## Examples

```Go
package main

import (
	"fmt"
	"github.com/mileusna/srb"
)

func main() {
    fmt.Println(srb.ToCyr("Kragujevac"))
    fmt.Println(srb.ToLat("кошарка"))
	fmt.Println(srb.HasCyr("Ima li ћирилице u ovom stringu?"))
    fmt.Println(srb.FixDj("Novak Djoković"))
    fmt.Println(srb.HTMLToCyr("<h1><a href='http://naslovi.net/'>Vesti dana</a></h1>"))
    fmt.Println(srb.ToASCII("Miloš"))
}
```
