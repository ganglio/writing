package ai

import (
	"github.com/invopop/jsonschema"
)

func generateSchema[T any]() interface{} {
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
		Location    string `json:"location" jsonschema_description:"The location of the character"`
		Description string `json:"description" jsonschema_description:"A brief description of what the character is doing or experiencing"`
	} `json:"details" jsonschema_description:"An array of character details, where each object represents a character with their name, location, and description"`
}

var (
	CharacterDetailsSchema = generateSchema[characterDetails]()
)
