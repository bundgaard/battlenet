// Copyright (C) 2021 David Bundgaard
package http

import (
	"fmt"
	"io"
)

type MediaType string

func (m MediaType) Validate(v string) bool {
	switch m {
	case MediaTypeItems:
		return v == "small" || v == "large"
	case MediaTypeSkills:
		return v == "21" || v == "42" || v == "64"
	}
	return false
}

const (
	MediaTypeItems  MediaType = "items"
	MediaTypeSkills MediaType = "skills"
)

func (c Client) Media(mediatype MediaType, size, icon string) ([]byte, error) {
	if ok := mediatype.Validate(size); !ok {
		return nil, fmt.Errorf("%s was not in (%q %q) or between (%q <=> %q)", size, "small", "large", "21", "64")
	}

	resp, err := c.TokenRequest("GET", nil, "http://media.blizzard.com/d3/icons/%s/%s/%s.png", mediatype, size, icon)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
