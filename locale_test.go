package i18n

import (
  "testing"
  assert "github.com/pilu/miniassert"
)

func TestNewLocale(t *testing.T) {
  locale := NewLocale("en")
  assert.Equal(t, "en", locale.Code)
  assert.Equal(t, 0, len(locale.translations))
}

func TestLocale_Add(t *testing.T) {
  locale := NewLocale("en")
  tr := locale.Add("foo", "bar")
  assert.Equal(t, 1, len(locale.translations))
  assert.Equal(t, tr, locale.translations["foo"])
}

func TestLocale_Translate(t *testing.T) {
  locale := NewLocale("en")
  locale.Add("greeting", "hello %s")
  value, err := locale.Translate("greeting", "world")
  assert.Nil(t, err)
  assert.Equal(t, "hello world", value)
}

func TestLocale_T(t *testing.T) {
  locale := NewLocale("en")
  locale.Add("greeting", "hello %s")
  value, err := locale.T("greeting", "world")
  assert.Nil(t, err)
  assert.Equal(t, "hello world", value)
}
