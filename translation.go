package i18n

import (
	"fmt"
)

// Translation is a pair of strings, the key and the value, which is the actual translation.
type Translation [2]string

func (t *Translation) Key() string {
	return t[0]
}

func (t *Translation) Value() string {
	return t[1]
}

// Returns fmt.Sprintf(t.Value(), args...)
func (t *Translation) Format(args ...interface{}) string {
	return fmt.Sprintf(t.Value(), args...)
}

// Returns a new Translation.
func NewTranslation(key, value string) *Translation {
	return &Translation{key, value}
}
