package pipeline

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/saantiaguilera/go-pipeline/pkg"
)

type HttpPipeline interface {
	Run(context *http_wrapper.Context, stage pkg.Stage)
}
