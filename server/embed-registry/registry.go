package embedregistry

import (
	"errors"
)

type EmbedRegistry struct {
	mapping map[string]string
}

func NewEmbedRegistry() *EmbedRegistry {
	m := &EmbedRegistry{
		mapping: make(map[string]string),
	}
	m.mapping["U102975477"] = "https://www.youtube.com/embed/uGGQGoht6ic"
	return m
}
func (e *EmbedRegistry) GetVideo(streamId string) (string, error) {
	val, ok := e.mapping[streamId]
	if !ok {
		return "", errors.New("streamId not found")
	}

	return val, nil
}

func (e *EmbedRegistry) UpdateVideo(streamId string, videoId string) {
	e.mapping[streamId] = videoId
}

func (e *EmbedRegistry) RemoveVideo(streamId string) {
	delete(e.mapping, streamId)
}
