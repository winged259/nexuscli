package model

import (
	"fmt"
)

type Registry struct {
	Hostname   string `json:"hostname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Repository string `json:"repository"`
}

type Repository struct {
	Images []string `json:"images"`
}

type Image struct {
	Name string   `json:"name"`
	Tag  []string `json:"tag"`
}

type Layer struct {
	Size   int    `json:"size"`
	Digest string `json:"digest"`
}

func (r Registry) GetToken() (string, error) {
	// Implement get token by username and password
	fmt.Println(r)
	return "", nil
}

func (r Repository) ListImages(token string) ([]Image, error) {
	// Implement list images from repository by token
	fmt.Println(token)
	return []Image{}, nil
}
