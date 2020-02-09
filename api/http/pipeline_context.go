package http

type PipelineContext interface {
	GetReader() (Reader, error)
	GetWriter() (Writer, error)
	GetMiddleware() (Middleware, error)
}
