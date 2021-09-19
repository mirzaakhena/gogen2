package vo

import "strings"

type Naming string

func (r Naming) String() string {
	return string(r)
}

func (r Naming) LowerCase() string {
	return strings.ToLower(string(r))
}

func (r Naming) HasPrefix(str string) bool {
	return strings.HasPrefix(strings.ToLower(string(r)), str)
}

func (r Naming) HasOneOfThisPrefix(str ...string) bool {
	lc := r.LowerCase()
	for _, s := range str {
		if lc == s {
			return true
		}
	}
	return false
}

func (r Naming) IsEmpty() bool {
	return len(string(r)) == 0
}