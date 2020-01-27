package pipeline

import (
	"github.com/arpb2/C-3PO/api/http"
	gopipeline "github.com/saantiaguilera/go-pipeline/pkg/api"
)

type HttpPipeline interface {
	Run(context *http.Context, stage gopipeline.Stage)
}
