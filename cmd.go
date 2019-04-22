// The command prints "Hello!" in multiple languages.
//
// Usage:
//       hello [lang]
// Currently supported languages: en, ua, fr, jp.
// When language is not set, it's chosen randomly.
//
// Example:
//       $ hello
//       こんにちは
//       $ hello fr
//       Salut!
package main // import "rmazur.io/hello"

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var greetings = map[string]string{
	"en": "Hello!",
	"ua": "Вітаю!",
	"fr": "Salut!",
	"jp": "こんにちは",
}

func main() {

	var lang string
	if len(os.Args) > 1 {
		lang = strings.ToLower(os.Args[1])
		if _, know := greetings[lang]; !know {
			fmt.Fprintf(os.Stderr, "I don't speak %s\n", lang)
			os.Exit(1)
		}
	} else {
		keys := make([]string, len(greetings))
		i := 0
		for k := range greetings {
			keys[i] = k
			i++
		}

		lang = keys[rand.Int31n(int32(len(keys)))]
	}

	fmt.Println(greetings[lang])
}
