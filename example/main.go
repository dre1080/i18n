package main

import (
	"fmt"
	"github.com/dre1080/i18n"
)

func t(key string) string {
	value, err := i18n.T(key)
	if err == nil {
		return value
	}

	return fmt.Sprintf("translation missing: `%s.%s`", i18n.CurrentLocale().Code, key)
}

func main() {
	i18n.Load("locales/*.conf")
	i18n.SetLocale("en")

	fmt.Println("en:")
	fmt.Printf("  Greeting: %s \n", t("greeting"))
	fmt.Printf("  Greeting 2: %s \n", t("greeting 2"))

	i18n.SetLocale("it")

	fmt.Println("it:")
	fmt.Printf("  Greeting: %s \n", t("greeting"))
	fmt.Printf("  Greeting 2: %s \n", t("greeting 2"))

	fmt.Println("es:")
	err := i18n.SetLocale("es")
	if err != nil {
		fmt.Println(err.Error())
	}
}
