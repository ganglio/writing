package ai

import (
	"github.com/openai/openai-go/v3"
	log "github.com/sirupsen/logrus"
)

func Chat(systemPrompt string, userPrompt string, schemaParam ...any) (string, error) {
	params := baseParams
	if systemPrompt != "" {
		params.Messages = append(params.Messages, openai.SystemMessage(systemPrompt))
	}
	params.Messages = append(params.Messages, openai.UserMessage(userPrompt))

	if len(schemaParam) > 0 {
		schema := openai.ResponseFormatJSONSchemaJSONSchemaParam{
			Name:        "structured_response",
			Description: openai.String("A structured response following the provided schema"),
			Schema:      schemaParam[0],
			Strict:      openai.Bool(true),
		}
		params.ResponseFormat = openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{JSONSchema: schema},
		}
	}

	for {
		completion, err := client.Chat.Completions.New(ctx, params)
		if err != nil {
			log.WithError(err).
				WithFields(map[string]any{
					"params": params,
				}).Error("failed to get chat completion")
			return "", err
		}

		response := completion.Choices[0].Message.Content
		stepFinishReason := completion.Choices[0].FinishReason

		log.WithField("choice", stepFinishReason).
			Debug("chat step finish reason")

		if stepFinishReason == "stop" {
			return response, nil
		}

		toolCalls := completion.Choices[0].Message.ToolCalls

		if len(toolCalls) == 0 {
			return response, nil
		}

		params.Messages = append(params.Messages, completion.Choices[0].Message.ToParam())

		processToolCalls(toolCalls, &params)
	}
}
