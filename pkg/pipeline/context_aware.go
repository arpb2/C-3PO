package pipeline

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	api "github.com/arpb2/C-3PO/api/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

const (
	TagHttpReader     pipeline.Tag = "tag_http_reader"
	TagHttpWriter     pipeline.Tag = "tag_http_writer"
	TagHttpMiddleware pipeline.Tag = "tag_http_middleware"
)

type context struct {
	pipeline.Context
}

func (c *context) GetReader() (http.Reader, error) {
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

func (c *context) GetWriter() (http.Writer, error) {
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

func (c *context) GetMiddleware() (http.Middleware, error) {
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

func (c *context) GetAuthenticatedUser(key pipeline.Tag) (model.AuthenticatedUser, error) {
	var user model.AuthenticatedUser
	val, exists := c.Get(key)

	if !exists {
		return user, http.CreateInternalError()
	}

	if user, ok := val.(model.AuthenticatedUser); ok {
		return user, nil
	}
	return user, http.CreateInternalError()
}

func (c *context) GetUser(key pipeline.Tag) (model.User, error) {
	var user model.User
	val, exists := c.Get(key)

	if !exists {
		return user, http.CreateInternalError()
	}

	if user, ok := val.(model.User); ok {
		return user, nil
	}
	return user, http.CreateInternalError()
}

func (c *context) GetUserLevel(key pipeline.Tag) (model.UserLevel, error) {
	var userLevel model.UserLevel
	val, exists := c.Get(key)

	if !exists {
		return userLevel, http.CreateInternalError()
	}

	if userLevel, ok := val.(model.UserLevel); ok {
		return userLevel, nil
	}
	return userLevel, http.CreateInternalError()
}

func (c *context) GetSession(key pipeline.Tag) (model.Session, error) {
	var session model.Session
	val, exists := c.Get(key)

	if !exists {
		return session, http.CreateInternalError()
	}

	if session, ok := val.(model.Session); ok {
		return session, nil
	}
	return session, http.CreateInternalError()
}

func CreateContextAware(delegate pipeline.Context) api.Context {
	return &context{
		Context: delegate,
	}
}
