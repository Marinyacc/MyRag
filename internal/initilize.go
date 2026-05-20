package internal

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	embedding "github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino-ext/components/model/ark"
	retrieveing "github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/joho/godotenv"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

func Init(ctx context.Context) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//语言模型
	ChatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("API_KEY"),
		Model:  os.Getenv("CHAT_MODEL"),
	})
	if err != nil {
		panic(err)
	}

	//嵌入模型
	apiType := embedding.APITypeMultiModal
	Embedder, err := embedding.NewEmbedder(ctx, &embedding.EmbeddingConfig{
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("EMBEDDING_MODEL"),
		APIType: &apiType,
	})
	if err != nil {
		panic(err)
	}

	//分词器
	Splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: false,
	})
	if err != nil {
		panic(err)
	}

	//向量数据库
	//表名
	collection := "TEST"
	//表字段
	fields := []*entity.Field{
		{
			Name:     "id",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "255",
			},
			PrimaryKey: true,
		},
		{
			Name:     "vector",
			DataType: entity.FieldTypeBinaryVector,
			TypeParams: map[string]string{
				//使用doubao-embedding-vison,返回浮点数且维数为2048
				"dim": "65536",
			},
		},
		{
			Name:     "content",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "8192",
			},
		},
		{
			Name:     "metadata",
			DataType: entity.FieldTypeJSON,
		},
	}
	cli := InitClient(ctx)
	Indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     *cli,
		Collection: collection,
		Fields:     fields,
		Embedding:  Embedder,
	})
	if err != nil {
		panic(err)
	}

	//检索组件
	Retriever, err := retrieveing.NewRetriever(ctx, &retrieveing.RetrieverConfig{
		Client:      *cli,
		Collection:  collection,
		VectorField: "vector",
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
		TopK:      5,
		Embedding: Embedder,
	})
	if err != nil {
		panic(err)
	}

	MyRag = &RagEngine{
		ChatModel: ChatModel,
		Embedder:  Embedder,
		Indexer:   Indexer,
		Retriever: Retriever,
		Splitter:  &Splitter,
	}

	docInit(ctx)
}
