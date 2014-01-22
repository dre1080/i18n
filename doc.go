/*
Package i18n implements i18n with simple config file or custom backends.
Files are simple config files based on https://github.com/mojombo/toml.

You can load a file with i18n.Load:

    err := i18n.Load("sample.conf")

You can also load multiple files passing a glob pattern:

    i18n.Load("/path/to/locals/*.conf")

Add all your translations under the "en" section:

  [en]
  greeting = "hello"

Inside a file you can have multiple sections. Each section will be a different locale:

  // file sample.conf

  [en]
  greeting = "hello"

  [it]
  greeting = "ciao"

  [es]
  greeting = "hola"

Check the example folder for a complete example https://github.com/dre1080/i18n/tree/master/example
*/
package i18n
