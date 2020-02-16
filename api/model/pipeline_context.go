package model

import "github.com/saantiaguilera/go-pipeline"

type PipelineContext interface {
	GetAuthenticatedUser(key pipeline.Tag) (AuthenticatedUser, error)
	GetUser(key pipeline.Tag) (User, error)
	GetUserLevel(key pipeline.Tag) (UserLevel, error)
	GetSession(key pipeline.Tag) (Session, error)
}
