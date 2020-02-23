package decorator

import "github.com/saantiaguilera/go-pipeline"

// Given a runnable, decorate it returning a mutation (or a completely new) one
type RunnableDecorator func(runnable pipeline.Runnable) pipeline.Runnable
