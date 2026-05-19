package internal

import (
	"context"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

var systemPrompt = `
# Role: Student Learning Assistant

# Language: Chinese

- When providing assistance:
  • Be clear and concise
  • Include practical examples when relevant
  • Reference documentation when helpful
  • Suggest improvements or next steps if applicable

here's documents searched for you:
==== doc start ====
	  {documents}
==== doc end ====
`

func Process(ctx context.Context, query string) string {
	results, err := MyRag.Retriever.Retrieve(ctx, query)
	if err != nil {
		panic(err)
	}
	tpl := prompt.FromMessages(schema.FString, []schema.MessagesTemplate{
		schema.SystemMessage(systemPrompt),
		schema.UserMessage("questiong:{query}"),
	}...)
	params := map[string]any{
		"doc":   results,
		"query": query,
	}
	messages, err := tpl.Format(ctx, params)
	resp, err := MyRag.ChatModel.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	return resp.Content
}
