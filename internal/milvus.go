package internal

import (
	"context"

	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
)

func InitClient(ctx context.Context) *cli.Client {
	client, err := cli.NewClient(ctx, cli.Config{
		Address: "localhost:19530",
		DBName: "default",
		Username: "root",
		Password: "",
	})
	if err != nil {
		panic(err)
	}
	return &client
}
