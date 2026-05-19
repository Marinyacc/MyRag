package internal

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cloudwego/eino/schema"
)

func docInit(ctx context.Context) {
	root := "../document"
	var docs []*schema.Document
	cnt := 0

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".md" || ext == ".markdown" {
			cnt++
			b, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			docs = append(docs, &schema.Document{
				ID:      "doc:" + strconv.Itoa(cnt),
				Content: string(b),
			})
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	splitter := *MyRag.Splitter
	results, err := splitter.Transform(ctx, docs)
	if err != nil {
		panic(err)
	}

	var Docs []*schema.Document
	var tmp string
	var Count int
	for _, doc := range results {
		if tmp != doc.ID {
			tmp = doc.ID
			Count = 1
		}
		Docs = append(Docs, &schema.Document{
			ID:      doc.ID + ":" + strconv.Itoa(Count),
			Content: doc.Content,
		})
	}

	_, err = MyRag.Indexer.Store(ctx, Docs)
	if err != nil {
		panic(err)
	}
}
