package pipeline

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model/classroom"
	"github.com/arpb2/C-3PO/pkg/domain/model/level"
	"github.com/arpb2/C-3PO/pkg/domain/model/session"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/saantiaguilera/go-pipeline"
)

const (
	TagHttpReader     pipeline.Tag = "tag_http_reader"
	TagHttpWriter     pipeline.Tag = "tag_http_writer"
	TagHttpMiddleware pipeline.Tag = "tag_http_middleware"
)

type ContextAware struct {
	pipeline.Context
}

func (c *ContextAware) GetReader() (http.Reader, error) {
	var reader http.Reader
	val, exists := c.Get(TagHttpReader)

	if !exists {
		return reader, http.CreateInternalError()
	}

	if reader, ok := val.(http.Reader); ok {
		return reader, nil
	}
	return reader, http.CreateInternalError()
}

func (c *ContextAware) GetWriter() (http.Writer, error) {
	var writer http.Writer
	val, exists := c.Get(TagHttpWriter)

	if !exists {
		return writer, http.CreateInternalError()
	}

	if writer, ok := val.(http.Writer); ok {
		return writer, nil
	}
	return writer, http.CreateInternalError()
}

func (c *ContextAware) GetMiddleware() (http.Middleware, error) {
	var middleware http.Middleware
	val, exists := c.Get(TagHttpMiddleware)

	if !exists {
		return middleware, http.CreateInternalError()
	}

	if middleware, ok := val.(http.Middleware); ok {
		return middleware, nil
	}
	return middleware, http.CreateInternalError()
}

func (c *ContextAware) GetAuthenticatedUser(key pipeline.Tag) (user.AuthenticatedUser, error) {
	var u user.AuthenticatedUser
	val, exists := c.Get(key)

	if !exists {
		return u, http.CreateInternalError()
	}

	if u, ok := val.(user.AuthenticatedUser); ok {
		return u, nil
	}
	return u, http.CreateInternalError()
}

func (c *ContextAware) GetClassroom(key pipeline.Tag) (classroom.Classroom, error) {
	var cr classroom.Classroom
	val, exists := c.Get(key)

	if !exists {
		return cr, http.CreateInternalError()
	}

	if cr, ok := val.(classroom.Classroom); ok {
		return cr, nil
	}
	return cr, http.CreateInternalError()
}

func (c *ContextAware) GetUser(key pipeline.Tag) (user.User, error) {
	var u user.User
	val, exists := c.Get(key)

	if !exists {
		return u, http.CreateInternalError()
	}

	if u, ok := val.(user.User); ok {
		return u, nil
	}
	return u, http.CreateInternalError()
}

func (c *ContextAware) GetUserLevel(key pipeline.Tag) (user.Level, error) {
	var userLevel user.Level
	val, exists := c.Get(key)

	if !exists {
		return userLevel, http.CreateInternalError()
	}

	if userLevel, ok := val.(user.Level); ok {
		return userLevel, nil
	}
	return userLevel, http.CreateInternalError()
}

func (c *ContextAware) GetUserLevelData(key pipeline.Tag) (user.LevelData, error) {
	var userLevelData user.LevelData
	val, exists := c.Get(key)

	if !exists {
		return userLevelData, http.CreateInternalError()
	}

	if userLevelData, ok := val.(user.LevelData); ok {
		return userLevelData, nil
	}
	return userLevelData, http.CreateInternalError()
}

func (c *ContextAware) GetLevel(key pipeline.Tag) (level.Level, error) {
	var l level.Level
	val, exists := c.Get(key)

	if !exists {
		return l, http.CreateInternalError()
	}

	if l, ok := val.(level.Level); ok {
		return l, nil
	}
	return l, http.CreateInternalError()
}

func (c *ContextAware) GetSession(key pipeline.Tag) (session.Session, error) {
	var s session.Session
	val, exists := c.Get(key)

	if !exists {
		return s, http.CreateInternalError()
	}

	if s, ok := val.(session.Session); ok {
		return s, nil
	}
	return s, http.CreateInternalError()
}

func CreateContextAware(delegate pipeline.Context) *ContextAware {
	return &ContextAware{
		Context: delegate,
	}
}
