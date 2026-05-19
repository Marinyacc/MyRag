package internal

import (
	embedding "github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino-ext/components/model/ark"
	retrieveing "github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/cloudwego/eino/components/document"
)

type RagEngine struct {
	ChatModel *ark.ChatModel
	Embedder  *embedding.Embedder
	Indexer   *milvus.Indexer
	Retriever *retrieveing.Retriever
	Splitter  *document.Transformer
}

// internal包下的全局变量
var MyRag *RagEngine
