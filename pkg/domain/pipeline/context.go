package pipeline

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/saantiaguilera/go-pipeline"
)

type Context interface {
	pipeline.Context

	http.PipelineContext
	model.PipelineContext
}
