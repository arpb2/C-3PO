package model

import "github.com/saantiaguilera/go-pipeline"

type Context interface {
	GetAuthenticatedUser(key pipeline.Tag) (AuthenticatedUser, error)
	GetUser(key pipeline.Tag) (User, error)
	GetCode(key pipeline.Tag) (Code, error)
	GetSession(key pipeline.Tag) (Session, error)
}
