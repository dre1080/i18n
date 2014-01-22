package i18n

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestNewTranslation(t *testing.T) {
	tr := NewTranslation("foo", "bar")
	assert.Equal(t, "foo", tr[0])
	assert.Equal(t, "bar", tr[1])
}

func TestNewTranslation_Key(t *testing.T) {
	tr := NewTranslation("foo", "bar")
	assert.Equal(t, "foo", tr.Key())
}

func TestNewTranslation_Value(t *testing.T) {
	tr := NewTranslation("foo", "bar")
	assert.Equal(t, "bar", tr.Value())
}

func TestTranslation_Translate(t *testing.T) {
	var tr *Translation

	tr = NewTranslation("greeting", "hello")
	assert.Equal(t, "hello", tr.Format())

	tr = NewTranslation("greeting", "hello %s")
	assert.Equal(t, "hello world", tr.Format("world"))
}
