package main

import (
	"fmt"
	"os"

	controller2 "github.com/arpb2/C-3PO/pkg/presentation/session/controller"
	controller3 "github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	controller4 "github.com/arpb2/C-3PO/pkg/presentation/user_level/controller"

	gopipeline "github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	pipeline2 "github.com/arpb2/C-3PO/pkg/infra/pipeline"
)

func createDrawablePipeline(fileName string) pipeline.HttpPipeline {
	file, err := os.Create(fmt.Sprintf("./build/uml/%s.svg", fileName))

	if err != nil {
		fmt.Println(os.Getwd())
		panic(err)
	}

	graphRenderer := gopipeline.CreateUMLActivityRenderer(gopipeline.UMLOptions{
		Type: gopipeline.UMLFormatSVG,
	})

	return pipeline2.CreateDrawablePipeline(file, graphRenderer)
}

func getPipelinedControllers() []controller.Controller {
	return []controller.Controller{
		controller4.CreateGetController(createDrawablePipeline("user_level_get_controller"), nil, nil),
		controller4.CreatePutController(createDrawablePipeline("user_level_put_controller"), nil, nil),
		controller2.CreatePostController(createDrawablePipeline("session_post_controller"), nil, nil, nil),
		controller3.CreateGetController(createDrawablePipeline("user_get_controller"), nil, nil),
		controller3.CreatePostController(createDrawablePipeline("user_post_controller"), nil, nil),
		controller3.CreatePutController(createDrawablePipeline("user_put_controller"), nil, nil, nil),
		controller3.CreateDeleteController(createDrawablePipeline("user_delete_controller"), nil, nil),
	}
}

func main() {
	for _, c := range getPipelinedControllers() {
		c.Body(&http.Context{})
	}
}
