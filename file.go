package i18n

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
)

type Options map[string]string

type Sections map[string]Options

func loadFile(filePath string) error {
	sections := make(Sections)
	_, err := toml.DecodeFile(filePath, &sections)
	if err != nil {
		return err
	}

	for section, options := range sections {
		if len(options) > 0 {
			locale := GetLocale(section)
			if locale == nil {
				locale = NewLocale(section)
			}
			for key, value := range options {
				locale.Add(key, value)
			}
			AddLocale(locale)
		}
	}
	return nil
}

// Loads translations from one or more files passing globPatterns to filepath.Glob.
func Load(globPattern string) error {
	paths, err := filepath.Glob(globPattern)
	if err != nil {
		return nil
	}

	for _, filePath := range paths {
		err := loadFile(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
