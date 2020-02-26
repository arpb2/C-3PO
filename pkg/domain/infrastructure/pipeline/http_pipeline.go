package pipeline

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/saantiaguilera/go-pipeline"
)

type HttpPipeline interface {
	Run(context *http.Context, stage pipeline.Stage)
}
