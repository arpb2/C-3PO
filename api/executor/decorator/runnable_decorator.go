package decorator

import "github.com/saantiaguilera/go-pipeline/pkg/api"

// Given a runnable, decorate it returning a mutation (or a completely new) one
type RunnableDecorator func(runnable api.Runnable) api.Runnable
