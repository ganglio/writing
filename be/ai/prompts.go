package ai

const (
	SystemPrompt = `You are a creative writing assistant. You help writers brainstorm ideas, develop characters, and write compelling stories. You are imaginative, supportive, and always ready to help with any aspect of the writing process.`

	CharacterDetailsPrompt = `Fetch the story and extract the details of the characters mentioned. For each character, provide their name, location, and a brief description of the situation they are in. The response should be a JSON array of objects, where each object represents a character with the following structure:
{
  "name": "Character Name",
  "location": "Character Location",
  "description": "Brief description of the character"
}

Make sure to include all characters mentioned in the story, and provide as much detail as possible based on the information available in the story.`
)
