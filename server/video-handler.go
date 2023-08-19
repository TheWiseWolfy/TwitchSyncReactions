package main

import (
	embedregistry "ebs_server/embed-registry"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddVideoRequest struct {
	VideoUrl string `json:"videoUrl"`
}

type VideoResponse struct {
	VideoUrl string `json:"videoUrl"`
}

type VideoHandler struct {
	embedRegistry *embedregistry.EmbedRegistry
}

func NewVideoHandler(embedRegistry *embedregistry.EmbedRegistry) *VideoHandler {
	return &VideoHandler{
		embedRegistry: embedRegistry,
	}
}

func (v *VideoHandler) GetVideo(c echo.Context) error {
	streamId := c.QueryParam("streamId")
	if streamId == "" {
		return c.String(http.StatusBadRequest, "")
	}
	var res VideoResponse
	videoId, err := v.embedRegistry.GetVideo(streamId)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	res.VideoUrl = videoId
	return c.JSON(http.StatusOK, res)
}

func (v *VideoHandler) UpdateVideo(c echo.Context) error {
	streamId := c.QueryParam("streamId")
	if streamId == "" {
		return c.String(http.StatusBadRequest, "")
	}
	var req AddVideoRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	v.embedRegistry.UpdateVideo(streamId, req.VideoUrl)
	return c.String(http.StatusOK, "Video updated")
}
