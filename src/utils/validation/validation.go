package validation

import "regexp"

var ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
