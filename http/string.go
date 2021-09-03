// Copyright (C) 2021 David Bundgaard
package http

import "strings"

func Normalize(battleTag string) string {
	return strings.Replace(battleTag, "#", "%23", -1)
}
