package embed

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
)

const modelName = "mxbai-embed-large:latest"

var ollamaM *ollama.LLM

func init() {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		panic(err)
	}
	ollamaM = llm
}

func EmbeddingString(ctx context.Context, val []string) ([][]float32, error) {
	return ollamaM.CreateEmbedding(ctx, val)
}

func EmbeddingStringStore(ctx context.Context, vector [][]float32) {

}

func NearestSimilarVector(ctx context.Context) {

}
