package pipeline

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	levelmodel "github.com/arpb2/C-3PO/pkg/domain/level/model"
	sessionmodel "github.com/arpb2/C-3PO/pkg/domain/session/model"
	usermodel "github.com/arpb2/C-3PO/pkg/domain/user/model"
	userlevelmodel "github.com/arpb2/C-3PO/pkg/domain/user_level/model"
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

func (c *ContextAware) GetAuthenticatedUser(key pipeline.Tag) (usermodel.AuthenticatedUser, error) {
	var user usermodel.AuthenticatedUser
	val, exists := c.Get(key)

	if !exists {
		return user, http.CreateInternalError()
	}

	if user, ok := val.(usermodel.AuthenticatedUser); ok {
		return user, nil
	}
	return user, http.CreateInternalError()
}

func (c *ContextAware) GetUser(key pipeline.Tag) (usermodel.User, error) {
	var user usermodel.User
	val, exists := c.Get(key)

	if !exists {
		return user, http.CreateInternalError()
	}

	if user, ok := val.(usermodel.User); ok {
		return user, nil
	}
	return user, http.CreateInternalError()
}

func (c *ContextAware) GetUserLevel(key pipeline.Tag) (userlevelmodel.UserLevel, error) {
	var userLevel userlevelmodel.UserLevel
	val, exists := c.Get(key)

	if !exists {
		return userLevel, http.CreateInternalError()
	}

	if userLevel, ok := val.(userlevelmodel.UserLevel); ok {
		return userLevel, nil
	}
	return userLevel, http.CreateInternalError()
}

func (c *ContextAware) GetUserLevelData(key pipeline.Tag) (userlevelmodel.UserLevelData, error) {
	var userLevelData userlevelmodel.UserLevelData
	val, exists := c.Get(key)

	if !exists {
		return userLevelData, http.CreateInternalError()
	}

	if userLevelData, ok := val.(userlevelmodel.UserLevelData); ok {
		return userLevelData, nil
	}
	return userLevelData, http.CreateInternalError()
}

func (c *ContextAware) GetLevel(key pipeline.Tag) (levelmodel.Level, error) {
	var level levelmodel.Level
	val, exists := c.Get(key)

	if !exists {
		return level, http.CreateInternalError()
	}

	if level, ok := val.(levelmodel.Level); ok {
		return level, nil
	}
	return level, http.CreateInternalError()
}

func (c *ContextAware) GetSession(key pipeline.Tag) (sessionmodel.Session, error) {
	var session sessionmodel.Session
	val, exists := c.Get(key)

	if !exists {
		return session, http.CreateInternalError()
	}

	if session, ok := val.(sessionmodel.Session); ok {
		return session, nil
	}
	return session, http.CreateInternalError()
}

func CreateContextAware(delegate pipeline.Context) *ContextAware {
	return &ContextAware{
		Context: delegate,
	}
}
