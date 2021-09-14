package vo

import "strings"

type Naming string

func (r Naming) String() string {
	return string(r)
}

func (r Naming) LowerCase() string {
	return strings.ToLower(string(r))
}