package i18n

import (
	"errors"
	"fmt"
)

type LocaleMissing struct {
	Code string
}

func (m LocaleMissing) Error() string {
	return fmt.Sprintf("Missing `%s` locale", m.Code)
}

var (
	locales       map[string]*Locale
	currentLocale *Locale
)

func init() {
	resetLocales()
}

func resetLocales() {
	locales = make(map[string]*Locale)
}

// Adds a locale to the available locales.
func AddLocale(locale *Locale) {
	locales[locale.Code] = locale
}

// Return the locale with the specified code
func GetLocale(code string) *Locale {
	return locales[code]
}

// Sets the current locale. Returns a LocaleMissing error if no locale has been added with the specified code.
func SetLocale(code string) error {
	currentLocale = GetLocale(code)

	if currentLocale == nil {
		return &LocaleMissing{code}
	}

	return nil
}

// Returns the current locale.
func CurrentLocale() *Locale {
	return currentLocale
}

// Calls Translate on the current locale.
// An error is returned if the current locale is not set or if the translation is missing in the current locale.
func Translate(key string, args ...interface{}) (string, error) {
	if currentLocale != nil {
		return currentLocale.Translate(key, args...)
	}

	return "", errors.New("Current locale is not set")
}

// Alias for Translate.
func T(key string, args ...interface{}) (string, error) {
	return Translate(key, args...)
}
