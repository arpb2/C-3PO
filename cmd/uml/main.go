package main

import (
	"fmt"
	"os"

	gopipeline "github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	"github.com/arpb2/C-3PO/pkg/controller/session"
	"github.com/arpb2/C-3PO/pkg/controller/user"
	"github.com/arpb2/C-3PO/pkg/controller/user_level"
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
		user_level.CreateGetController(createDrawablePipeline("user_level_get_controller"), nil, nil),
		user_level.CreatePostController(createDrawablePipeline("user_level_post_controller"), nil, nil),
		user_level.CreatePutController(createDrawablePipeline("user_level_put_controller"), nil, nil),
		session.CreatePostController(createDrawablePipeline("session_post_controller"), nil, nil, nil),
		user.CreateGetController(createDrawablePipeline("user_get_controller"), nil, nil),
		user.CreatePostController(createDrawablePipeline("user_post_controller"), nil, nil),
		user.CreatePutController(createDrawablePipeline("user_put_controller"), nil, nil, nil),
		user.CreateDeleteController(createDrawablePipeline("user_delete_controller"), nil, nil),
	}
}

func main() {
	for _, c := range getPipelinedControllers() {
		c.Body(&http.Context{})
	}
}
