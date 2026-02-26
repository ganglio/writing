package ai

const (
	SystemPrompt = `You are a creative writing assistant. You help writers brainstorm ideas, develop characters, and write compelling stories. You are imaginative, supportive, and always ready to help with any aspect of the writing process.`

	CharacterDetailsPrompt = `Fetch the story and extract the details of the characters mentioned. For each character, provide their name, last location, and a brief description of the situation they are in. The response should be a JSON array of objects, where each object represents a character with the following structure:
{
  "name": "Character Name",
  "chapter": "Last Chapter in which the character appears",
  "location": "Last Character Location",
  "description": "Brief description of what the character is doing or experiencing. Keep it concise, ideally less than 20 words."
}

Make sure to include all characters mentioned in the story, and provide as much detail as possible based on the information available in the story.`

	CharacterDetailsPromptWithFile = `Fetch the file "%s" and extract the details of all the characters mentioned. For each character, provide their name, last location, and a brief description of the situation they are in. The response should be a JSON array of objects, where each object represents a character with the following structure:
{
  "name": "Character Name",
  "chapter": "Last Chapter in which the character appears",
  "location": "Last Character Location",
  "description": "Brief description of what the character is doing or experiencing. Keep it concise, ideally less than 20 words."
}

Make sure to include all characters mentioned in the story, and provide as much detail as possible based on the information available in the story.`

	StoryTimelinePrompt = `Fetch the story and extract the timeline of events. For each event, provide a brief description and the characters involved. The response should be a JSON array of objects, where each object represents an event with the following structure:
{
  "description": "Brief description of the event. Keep it concise, ideally less than 20 words.",
  "characters": ["Character Name 1", "Character Name 2", ...]
}`

	StoryTimelinePromptWithFile = `Fetch the file "%s" and extract the timeline of events. For each event, provide a brief description and the characters involved. The response should be a JSON array of objects, where each object represents an event with the following structure:
{
  "description": "Brief description of the event. Keep it concise, ideally less than 20 words.",
  "Chapter": "Chapter in which the event occurs",
  "characters": ["Character Name 1", "Character Name 2", ...]
}`
)
