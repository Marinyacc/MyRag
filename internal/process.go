package internal

import (
	"context"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

var systemPrompt = `
# Role:Learning and Searching Assistant

# Language: Chinese

here's documents searched for you:
==== doc start ====
	  {documents}
==== doc end ====
`

func Process(ctx context.Context, query string) *schema.Message {
	results, err := MyRag.Retriever.Retrieve(ctx, query)
	if err != nil {
		panic(err)
	}
	tpl := prompt.FromMessages(schema.FString, []schema.MessagesTemplate{
		schema.SystemMessage(systemPrompt),
		schema.UserMessage("question:{query}"),
	}...)
	params := map[string]any{
		"documents": results,
		"query":     query,
	}
	messages, err := tpl.Format(ctx, params)
	if err != nil {
		panic(err)
	}
	resp, err := MyRag.ChatModel.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	return resp
}
