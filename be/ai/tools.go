package ai

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"writing/config"

	"github.com/openai/openai-go/v3"
	log "github.com/sirupsen/logrus"
)

var env = config.GetEnv()

var tools = []openai.ChatCompletionToolUnionParam{
	openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
		Name:        "ls",
		Description: openai.String("list the files in the current project"),
		Parameters: openai.FunctionParameters{
			"type": "object",
			"properties": map[string]any{
				"recursive": map[string]string{
					"type": "boolean",
				},
			},
		},
	}),
	openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
		Name:        "get_file_content",
		Description: openai.String("Get the content of a file in the current project"),
		Parameters: openai.FunctionParameters{
			"type": "object",
			"properties": map[string]any{
				"file": map[string]string{
					"type": "string",
				},
			},
			"required": []string{"file"},
		},
	}),
	openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
		Name:        "set_file_content",
		Description: openai.String("Update or create a file in the current project"),
		Parameters: openai.FunctionParameters{
			"type": "object",
			"properties": map[string]any{
				"file": map[string]string{
					"type": "string",
				},
				"content": map[string]string{
					"type": "string",
				},
			},
			"required": []string{"file", "content"},
		},
	}),
	openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
		Name:        "delete_file",
		Description: openai.String("Delete a file in the current project"),
		Parameters: openai.FunctionParameters{
			"type": "object",
			"properties": map[string]any{
				"file": map[string]string{
					"type": "string",
				},
			},
			"required": []string{"file"},
		},
	}),
}

func set_file_content(file string, content string) error {
	log.WithFields(map[string]any{
		"file":    file,
		"content": content,
	}).Debug("set_file_content")

	err := os.WriteFile(env.WorkFolder+"/"+file, []byte(content), 0644)
	if err != nil {
		log.WithFields(map[string]any{
			"file":    file,
			"content": content,
			"error":   err,
		}).Error("failed to write file")
		return err
	}
	return nil
}

func get_file_content(file string) (string, error) {
	log.WithFields(map[string]any{
		"file": file,
	}).Debug("get_file_content")

	content, err := os.ReadFile(env.WorkFolder + "/" + file)
	if err != nil {
		log.WithFields(map[string]any{
			"file":  file,
			"error": err,
		}).Error("failed to read file")
		return "", err
	}

	return string(content), nil
}

func delete_file(file string) error {
	log.WithFields(map[string]any{
		"file": file,
	}).Debug("delete_file")

	err := os.Remove(env.WorkFolder + "/" + file)
	if err != nil {
		log.WithFields(map[string]any{
			"file":  file,
			"error": err,
		}).Error("failed to delete file")
		return err
	}
	return nil
}

func ls(recursive bool) ([]string, error) {
	var files []string
	err := fs.WalkDir(os.DirFS(env.WorkFolder), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && !strings.HasPrefix(path, ".") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.WithFields(map[string]any{
			"error":     err,
			"recursive": recursive,
			"folder":    env.WorkFolder,
		}).Error("failed to list files")
		return nil, err
	}

	log.WithFields(map[string]any{
		"recursive": recursive,
		"files":     files,
	}).Debug("ls")

	return files, nil
}

func processToolCalls(toolCalls []openai.ChatCompletionMessageToolCallUnion, params *openai.ChatCompletionNewParams) {
	for _, toolCall := range toolCalls {
		switch toolCall.Function.Name {
		case "ls":
			var args struct {
				Recursive bool `json:"recursive"`
			}
			err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
			if err != nil {
				log.WithFields(map[string]any{
					"error": err,
				}).Error("failed to unmarshal ls args")
				continue
			}
			files, err := ls(args.Recursive)
			if err != nil {
				log.WithFields(map[string]any{
					"error": err,
				}).Error("failed to list files")
				continue
			}
			params.Messages = append(params.Messages, openai.ToolMessage(fmt.Sprintf(`["%s"]`, strings.Join(files, `","`)), toolCall.ID))
		case "get_file_content":
			var args struct {
				File string `json:"file"`
			}
			err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
			if err != nil {
				log.WithFields(map[string]any{
					"error": err,
				}).Error("failed to unmarshal get_file_content args")
				continue
			}
			content, err := get_file_content(args.File)
			if err != nil {
				log.WithFields(map[string]any{
					"file":  args.File,
					"error": err,
				}).Error("failed to get file content")
				continue
			}
			params.Messages = append(params.Messages, openai.ToolMessage(content, toolCall.ID))
		case "set_file_content":
			var args struct {
				File    string `json:"file"`
				Content string `json:"content"`
			}
			err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
			if err != nil {
				log.WithFields(map[string]any{
					"error": err,
				}).Error("failed to unmarshal set_file_content args")
				continue
			}
			err = set_file_content(args.File, args.Content)
			if err != nil {
				log.WithFields(map[string]any{
					"file":  args.File,
					"error": err,
				}).Error("failed to set file content")
				continue
			}
			params.Messages = append(params.Messages, openai.ToolMessage("file content set successfully", toolCall.ID))
		case "delete_file":
			var args struct {
				File string `json:"file"`
			}
			err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
			if err != nil {
				log.WithFields(map[string]any{
					"error": err,
				}).Error("failed to unmarshal delete_file args")
				continue
			}
			err = delete_file(args.File)
			if err != nil {
				log.WithFields(map[string]any{
					"file":  args.File,
					"error": err,
				}).Error("failed to delete file")
				continue
			}
			params.Messages = append(params.Messages, openai.ToolMessage("file deleted successfully", toolCall.ID))
		default:
			log.WithFields(map[string]any{
				"name": toolCall.Function.Name,
			}).Error("unknown tool call")
		}
	}
}
