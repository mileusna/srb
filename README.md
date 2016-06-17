# Srb for Go <a href="https://godoc.org/github.com/mileusna/srb"><img src="https://godoc.org/github.com/mileusna/srb?status.svg" alt="GoDoc"></a>

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

## License

[MIT](LICENSE)
