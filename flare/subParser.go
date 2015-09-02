package flare

// sub-parser provides a very basic ini-style reader.
// every [header] calls the section callback
// every key= value calls the keyvalue callback
// all other non-blank lines calls rawline
type subParser struct {
	section  func(string) bool
	keyvalue keyValueParser
	rawline  func(string) bool
}

type keyValueParser struct {
	name string
	cb   func(string, string) bool
}
