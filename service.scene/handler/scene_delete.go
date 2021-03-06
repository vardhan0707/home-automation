package handler

import (
	"github.com/jakewright/home-automation/libraries/go/database"
	"github.com/jakewright/home-automation/libraries/go/errors"
	"github.com/jakewright/home-automation/libraries/go/slog"
	scenedef "github.com/jakewright/home-automation/service.scene/def"
	"github.com/jakewright/home-automation/service.scene/domain"
)

// HandleDeleteScene deletes a scene and associated actions
func HandleDeleteScene(req *scenedef.DeleteSceneRequest) (*scenedef.DeleteSceneResponse, error) {
	if req.SceneId == 0 {
		return nil, errors.BadRequest("scene_id empty")
	}

	// Delete the scene
	if err := database.Delete(&domain.Scene{}, req.SceneId); err != nil {
		return nil, err
	}

	slog.Infof("Deleted scene %d", req.SceneId)
	return &scenedef.DeleteSceneResponse{}, nil
}
