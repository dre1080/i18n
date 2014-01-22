package i18n

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestAddLocale(t *testing.T) {
	defer resetLocales()

	locale := NewLocale("en")
	assert.Equal(t, 0, len(locales))

	AddLocale(locale)
	assert.Equal(t, 1, len(locales))
	assert.Equal(t, locale, locales["en"])
}

func TestGetLocale(t *testing.T) {
	defer resetLocales()

	locale := NewLocale("en")
	AddLocale(locale)

	assert.Equal(t, locale, GetLocale("en"))
}

func TestSetLocale(t *testing.T) {
	defer resetLocales()

	locale := NewLocale("en")
	AddLocale(locale)

	var err error

	err = SetLocale("en")
	assert.Nil(t, err)
	assert.Equal(t, locale, currentLocale)

	err = SetLocale("foo")
	assert.NotNil(t, err)
}

func TestTranslate(t *testing.T) {
	defer resetLocales()

	en := NewLocale("en")
	en.Add("greeting", "hello")
	en.Add("greeting2", "hello2")
	AddLocale(en)

	it := NewLocale("it")
	it.Add("greeting", "ciao")
	AddLocale(it)

	var value string
	var err error

	err = SetLocale("en")
	assert.Nil(t, err)
	assert.Equal(t, en, currentLocale)

	value, err = Translate("greeting")
	assert.Nil(t, err)
	assert.Equal(t, "hello", value)

	value, err = Translate("greeting2")
	assert.Nil(t, err)
	assert.Equal(t, "hello2", value)

	err = SetLocale("it")
	assert.Nil(t, err)
	assert.Equal(t, it, currentLocale)

	value, err = Translate("greeting")
	assert.Nil(t, err)
	assert.Equal(t, "ciao", value)

	value, err = Translate("greeting2")
	assert.NotNil(t, err)
	assert.Type(t, "*i18n.TranslationMissing", err)

	err = SetLocale("es")
	assert.NotNil(t, err)
	assert.Type(t, "*i18n.LocaleMissing", err)
	assert.True(t, currentLocale == nil)

	value, err = Translate("greeting")
	assert.NotNil(t, err)
	assert.Equal(t, "Current locale is not set", err.Error())
}
