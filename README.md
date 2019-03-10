# wheel - Implementacja genratora tokenów Cerb Wheel w języku Go.

Obsługuje jedynie wyliczenie tokena z domyślnymi parametrami: 

* tolerancja czasowa = 60 sekund
* domyślna strefa czasowa
* TOTP
* brak obsługi szyfowanego SEEDa (przy użyciu PINu)

Przykładowe uycie:

```go
package main

import (
	"fmt"

	"github.com/srozb/wheel"
)

func main() {
	t := wheel.NewToken("Seed w formie szesnastkowej")
	t.Generate()
	tokenSting := t.GetTokenString()
	fmt.Printf(tokenSting)
}
```

Seed można wydobyć z bazy sqlite z aplikacji na Androida, na przykład przy użyciu triku `adb backup`. 

Bardziej kompletna implementacja dostępna jest [tutaj](https://github.com/srozb/grumpytoken), ale pisałem ją milion lat temu i źródła są mocno nieczytelne. 

Brak obsługi DST (czasu letniego), więc po zmianie czasu może wyliczać niepoprawne wartości. Jeśli tak będzie to się poprawi ;-)

