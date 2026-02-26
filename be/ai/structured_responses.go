package ai

import (
	"github.com/invopop/jsonschema"
)

func generateSchema[T any]() any {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

type characterDetails struct {
	Details []struct {
		Name        string `json:"name" jsonschema_description:"The name of the character"`
		Chapter     string `json:"chapter" jsonschema_description:"The last chapter in which the character appears"`
		Location    string `json:"location" jsonschema_description:"The last location of the character"`
		Description string `json:"description" jsonschema_description:"A brief description of what the character is doing or experiencing"`
	} `json:"details" jsonschema_description:"An array of character details, where each object represents a character with their name, location, and description"`
}

type storyTimeline struct {
	Timeline []struct {
		Description string   `json:"description" jsonschema_description:"A brief description of the event"`
		Chapter     string   `json:"chapter" jsonschema_description:"The chapter in which the event occurs"`
		Characters  []string `json:"characters" jsonschema_description:"An array of character names involved in the event"`
	} `json:"timeline" jsonschema_description:"An array of events, where each object represents an event with a description and the characters involved"`
}

var (
	CharacterDetailsSchema = generateSchema[characterDetails]()
	StoryTimelineSchema    = generateSchema[storyTimeline]()
)
