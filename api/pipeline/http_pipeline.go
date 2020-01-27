package pipeline

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	go_pipeline "github.com/saantiaguilera/go-pipeline/pkg/api"
)

type HttpPipeline interface {
	Run(context *http_wrapper.Context, stage go_pipeline.Stage)
}
