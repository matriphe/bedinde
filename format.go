package bedinde

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatNumber(number uint64, lang string) string {
	holder := message.NewPrinter(language.Make(lang))
	return holder.Sprintf("%d", number)
}
