package pipeline

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/saantiaguilera/go-pipeline"
)

type Context interface {
	pipeline.Context
	http.PipelineContext
	model.Context
}
