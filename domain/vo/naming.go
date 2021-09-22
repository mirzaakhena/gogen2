package vo

import (
	"regexp"
	"strings"
	"unicode"
)

// Naming ...
type Naming string

// String ...
func (r Naming) String() string {
	return string(r)
}

// HasPrefix ...
func (r Naming) HasPrefix(str string) bool {
	return strings.HasPrefix(strings.ToLower(string(r)), str)
}

// HasOneOfThisPrefix ...
func (r Naming) HasOneOfThisPrefix(str ...string) bool {
	lc := r.LowerCase()
	for _, s := range str {
		if strings.HasPrefix(lc, s) {
			return true
		}
	}
	return false
}

// IsEmpty ...
func (r Naming) IsEmpty() bool {
	return len(string(r)) == 0
}

// CamelCase is
func (r Naming) CamelCase() string {

	name := string(r)

	// hardcoded is bad
	// But we can figure out later
	{
		if name == "IPAddress" {
			return "ipAddress"
		}

		if name == "ID" {
			return "id"
		}
	}

	out := []rune(name)
	out[0] = unicode.ToLower([]rune(name)[0])
	return string(out)
}

// UpperCase is
func (r Naming) UpperCase() string {
	name := string(r)
	return strings.ToUpper(name)
}

// LowerCase is
func (r Naming) LowerCase() string {
	name := string(r)
	return strings.ToLower(name)
}

var matchFirstCapSpaceCase = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCapSpaceCase = regexp.MustCompile("([a-z0-9])([A-Z])")

// SpaceCase is
func (r Naming) SpaceCase() string {
	str := string(r)

	snake := matchFirstCapSpaceCase.ReplaceAllString(str, "${1} ${2}")
	snake = matchAllCapSpaceCase.ReplaceAllString(snake, "${1} ${2}")
	return strings.ToLower(snake)
}

// PascalCase is
func (r Naming) PascalCase() string {
	name := string(r)
	rs := []rune(name)
	return strings.ToUpper(string(rs[0])) + string(rs[1:])
}

var matchFirstCapSnakeCase = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCapSnakeCase = regexp.MustCompile("([a-z0-9])([A-Z])")

// SnakeCase is
func (r Naming) SnakeCase() string {
	str := string(r)
	snake := matchFirstCapSnakeCase.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCapSnakeCase.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
