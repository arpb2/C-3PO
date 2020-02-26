package main

import (
	"fmt"
	"os"

	pipeline3 "github.com/arpb2/C-3PO/cmd/uml/infrastructure/pipeline"

	levelcontroller "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	sessioncontroller "github.com/arpb2/C-3PO/pkg/presentation/session/controller"
	usercontroller "github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	userlevelcontroller "github.com/arpb2/C-3PO/pkg/presentation/user_level/controller"

	gopipeline "github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
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

	return pipeline3.CreateDrawablePipeline(file, graphRenderer)
}

func getPipelinedControllers() []controller.Controller {
	return []controller.Controller{
		levelcontroller.CreateGetController(createDrawablePipeline("level_get_controller"), nil),
		levelcontroller.CreatePutController(createDrawablePipeline("level_put_controller"), nil, nil),

		userlevelcontroller.CreateGetController(createDrawablePipeline("user_level_get_controller"), nil, nil),
		userlevelcontroller.CreatePutController(createDrawablePipeline("user_level_put_controller"), nil, nil),

		sessioncontroller.CreatePostController(createDrawablePipeline("session_post_controller"), nil, nil, nil),

		usercontroller.CreateGetController(createDrawablePipeline("user_get_controller"), nil, nil),
		usercontroller.CreatePostController(createDrawablePipeline("user_post_controller"), nil, nil),
		usercontroller.CreatePutController(createDrawablePipeline("user_put_controller"), nil, nil, nil),
		usercontroller.CreateDeleteController(createDrawablePipeline("user_delete_controller"), nil, nil),
	}
}

func main() {
	for _, c := range getPipelinedControllers() {
		c.Body(&http.Context{})
	}
}
