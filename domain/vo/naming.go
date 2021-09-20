package vo

import (
	"strings"
)

// Naming ...
type Naming string

// String ...
func (r Naming) String() string {
	return string(r)
}

// LowerCase ...
func (r Naming) LowerCase() string {
	return strings.ToLower(string(r))
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
