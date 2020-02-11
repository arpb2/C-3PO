package pipeline_test

import (
	"os"
	"testing"

	"github.com/arpb2/C-3PO/pkg/pipeline"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"
	"github.com/stretchr/testify/mock"
)

func TestDrawablePipeline_GivenOne_WhenRunningOnAStage_ThenItDrawsAndRendersIt(t *testing.T) {
	stage := new(pipeline2.MockStage)
	stage.On("Draw", mock.Anything)
	renderer := new(pipeline2.MockRenderer)
	renderer.On("Render", mock.Anything, os.Stdout).Return(nil)
	pipe := pipeline.CreateDrawablePipeline(os.Stdout, renderer)

	pipe.Run(nil, stage)

	stage.AssertExpectations(t)
	renderer.AssertExpectations(t)
}
