package bedinde

import (
	"strings"
	"github.com/nyaruka/phonenumbers"
)

func FormatPhone(phone string, lang string) string {
	num, err := phonenumbers.Parse(phone, strings.ToUpper(lang))
	if err != nil {
		return ""
	}

	return phonenumbers.Format(num, phonenumbers.E164)
}
