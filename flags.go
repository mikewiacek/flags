// Package flags provides custom flag types that are compaitble with Go's
// standard flag package.
package flags

import (
	"flag"
	"fmt"
	"strings"
)

type stringSlice []string

func newStringSlice(defaults []string, f *[]string) *stringSlice {
	// Setup *f to be the default values passed in val.
	*f = defaults

	// return a *stringSlice cast of f.
	return (*stringSlice)(f)
}

// String helps implement flag.Value for stringSlice.
func (s *stringSlice) String() string {
	sb := strings.Builder{}
	for i, v := range *s {
		sb.WriteString(fmt.Sprintf("%s", v))
		if i < len(*s)-1 {
			sb.WriteRune(',')
		}
	}
	return sb.String()
}

// Set helps implement flag.Value for stringSlice.
func (s *stringSlice) Set(input string) error {
	*s = (stringSlice)(strings.Split(input, ","))
	return nil
}

// StringSlice implements a custom flag that accepts a list of strings in CSV form.
// No escaping of commas is supported.
func StringSlice(name string, defaultValues []string, usage string) *[]string {
	v := append([]string(nil), defaultValues...)
	ss := newStringSlice(defaultValues, &v)
	flag.Var(ss, name, usage)
	return &v
}
