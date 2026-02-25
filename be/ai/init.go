package ai

import (
	"context"

	"github.com/openai/openai-go/v3"
	log "github.com/sirupsen/logrus"
)

var (
	client     openai.Client
	ctx        context.Context
	baseParams = openai.ChatCompletionNewParams{
		Tools:    tools,
		Seed:     openai.Int(0),
		Model:    openai.ChatModelGPT5_2,
		Messages: []openai.ChatCompletionMessageParamUnion{},
	}
)

func init() {
	log.Debug("initializing AI module")

	client = openai.NewClient()
	ctx = context.Background()
}
