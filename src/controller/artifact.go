package controller

import (
	"github.com/UPrefer/StorageService/model"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ArtifactController struct {
	artifactService service.IArtifactService
}

func NewArtifactController(artifactService service.IArtifactService) *ArtifactController {
	return &ArtifactController{
		artifactService: artifactService,
	}
}

func (artifactController *ArtifactController) Post(ctx *gin.Context) {
	var (
		artifactToCreate model.ArtifactDTO
		createdArtifact  *model.ArtifactDTO
		err              error
	)

	err = ctx.BindJSON(&artifactToCreate)
	if err != nil {
		return
	}

	createdArtifact, err = artifactController.artifactService.CreateArtifact(&artifactToCreate)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
	ctx.Header("Location", ctx.Request.RequestURI+"/"+createdArtifact.Uuid)
}

func (artifactController *ArtifactController) Get(ctx *gin.Context) {
	var foundArtifact, err = artifactController.artifactService.ReadArtifact(ctx.Param("artifact_id"))

	if err != nil {
		log.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if foundArtifact == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, foundArtifact)
}
