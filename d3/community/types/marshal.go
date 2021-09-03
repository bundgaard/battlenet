// Copyright (C) 2021 David Bundgaard
package types

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

// ToBase64 convert to JSON and then to Base64
func (pr Profile) ToBase64() string {
	content, err := json.Marshal(pr)
	if err != nil {
		log.Fatal("ToBase64", err)
	}

	return base64.StdEncoding.EncodeToString(content)
}

// ToBase64 cnvert to JSON and further to Base64
func (hr Hero) ToBase64() string {
	content, err := json.Marshal(hr)
	if err != nil {
		log.Fatal("ToBase64", err)
	}
	return base64.StdEncoding.EncodeToString(content)

}

// NewHeroFromBase64 will create a hero out of text
func NewHeroFromBase64(data string) *Hero {
	content, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Println("failed to decode base64")
		return nil
	}

	var hero Hero
	if err := json.Unmarshal(content, &hero); err != nil {
		log.Println("failed to unmarshal hero")
		return nil
	}
	return &hero
}

// NewProfileFromBase64 will create a profile out of text
func NewProfileFromBase64(data string) *Profile {
	var p Profile
	content, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Println("failed to decode base64")
		return nil
	}
	if err := json.Unmarshal(content, &p); err != nil {
		log.Println("failed to unmarshal profile")
		return nil
	}
	return &p
}
