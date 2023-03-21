package repository

import "errors"

var (
	ErrorNotFound      = errors.New("id not found")
	ErrorAlreadyExists = errors.New("id already exists")
)
