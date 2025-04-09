package embed

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
)

const modelName = "mxbai-embed-large:latest"

var ollamaM *ollama.LLM

func init() {
	llm, err := ollama.New(ollama.WithModel(modelName))
	if err != nil {
		panic(err)
	}
	ollamaM = llm
}

func MustEmbeddingSingleString(ctx context.Context, val string) []float32 {
	embedding, err := ollamaM.CreateEmbedding(ctx, []string{val})
	if err != nil {
		panic(err)
	}
	if len(embedding) == 0 {
		panic("no embedding found")
	}
	if len(embedding[0]) == 0 {
		panic("no embedding found")
	}
	return embedding[0]
}

func EmbeddingString(ctx context.Context, val []string) ([][]float32, error) {
	return ollamaM.CreateEmbedding(ctx, val)
}

func EmbeddingStringStore(ctx context.Context, vector [][]float32) {

}

func NearestSimilarVector(ctx context.Context) {

}
