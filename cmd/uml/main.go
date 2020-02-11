package main

import (
	"fmt"
	"os"

	gopipeline "github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	"github.com/arpb2/C-3PO/pkg/controller/code"
	"github.com/arpb2/C-3PO/pkg/controller/session"
	"github.com/arpb2/C-3PO/pkg/controller/user"
	pipeline2 "github.com/arpb2/C-3PO/pkg/pipeline"
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
		code.CreateGetController(createDrawablePipeline("code_get_controller"), nil, nil),
		code.CreatePostController(createDrawablePipeline("code_post_controller"), nil, nil),
		code.CreatePutController(createDrawablePipeline("code_put_controller"), nil, nil),
		session.CreatePostController(createDrawablePipeline("session_post_controller"), nil, nil, nil),
		user.CreateGetController(createDrawablePipeline("user_get_controller"), nil, nil),
		user.CreatePostController(createDrawablePipeline("user_post_controller"), nil, nil),
		user.CreatePutController(createDrawablePipeline("user_put_controller"), nil, nil, nil),
		user.CreateDeleteController(createDrawablePipeline("user_delete_controller"), nil, nil),
	}
}

func main() {
	for _, controller := range getPipelinedControllers() {
		controller.Body(&http.Context{})
	}
}
