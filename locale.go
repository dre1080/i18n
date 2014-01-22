package i18n

import (
	"fmt"
)

type TranslationMissing struct {
	Key string
}

func (t TranslationMissing) Error() string {
	return fmt.Sprintf("Translation missing for key `%s`", t.Key)
}

// Locale represents a locale with its Code and translations
type Locale struct {
	Code         string
	Translations map[string]*Translation
}

// Adds a key and its translation to locale.
func (locale *Locale) Add(key, value string) *Translation {
	translation := NewTranslation(key, value)
	locale.Translations[translation.Key()] = translation

	return translation
}

// Returns the translation for the specified key.
// The translation and args are passed to Sprintf.
// If you have a key "message" with translation "hello %s", calling Translate("message", "world") will return "hello world".
// An error is returned if the key is missing.
func (locale *Locale) Translate(key string, args ...interface{}) (string, error) {
	var value string
	translation := locale.Translations[key]

	if translation == nil {
		return value, &TranslationMissing{key}
	}

	return translation.Format(args...), nil
}

// Alias for Translate
func (locale *Locale) T(key string, args ...interface{}) (string, error) {
	return locale.Translate(key, args...)
}

// Returns a new locale with the specified code.
func NewLocale(code string) *Locale {
	locale := &Locale{
		Code:         code,
		Translations: make(map[string]*Translation),
	}

	return locale
}
