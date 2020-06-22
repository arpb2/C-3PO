package main

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/presentation/classroom"
	"os"

	pipeline2 "github.com/arpb2/C-3PO/pkg/infrastructure/pipeline"
	"github.com/arpb2/C-3PO/pkg/presentation/level"
	"github.com/arpb2/C-3PO/pkg/presentation/session"
	"github.com/arpb2/C-3PO/pkg/presentation/user"

	gopipeline "github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
)

const (
	ParamLevelId = "level_id"
	ParamUserId  = "user_id"
	ParamClassroomId = "classroom_id"
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

func getPipelinedBodies() []http.Handler {
	return []http.Handler{
		level.CreateGetHandler(ParamLevelId, createDrawablePipeline("level_get_controller"), nil),
		level.CreatePutHandler(ParamLevelId, createDrawablePipeline("level_put_controller"), nil),

		user.CreateGetLevelHandler(ParamUserId, ParamLevelId, createDrawablePipeline("user_level_get_controller"), nil),
		user.CreatePutLevelHandler(ParamUserId, ParamLevelId, createDrawablePipeline("user_level_put_controller"), nil),

		session.CreatePostHandler(createDrawablePipeline("session_post_controller"), nil, nil, nil),

		user.CreateGetUserHandler(ParamUserId, createDrawablePipeline("user_get_controller"), nil),
		user.CreatePostUserHandler(createDrawablePipeline("user_post_controller"), nil, nil),
		user.CreatePutUserHandler(ParamUserId, createDrawablePipeline("user_put_controller"), nil, nil),
		user.CreateDeleteUserHandler(ParamUserId, createDrawablePipeline("user_delete_controller"), nil),

		classroom.CreateGetHandler(ParamClassroomId, createDrawablePipeline("classroom_get_controller"), nil),
		classroom.CreatePutHandler(ParamClassroomId, createDrawablePipeline("classroom_put_controller"), nil),
	}
}

func main() {
	for _, b := range getPipelinedBodies() {
		b(&http.Context{})
	}
}
