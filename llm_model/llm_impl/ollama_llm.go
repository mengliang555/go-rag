package llmimpl

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
	llmmodel "go-rag/llm_model"
)

type OllamaLLM struct {
	OllamaModel string
	llm         *ollama.LLM
}

func (o *OllamaLLM) GenerateFromSinglePrompt(ctx context.Context, promopt string) {
	//TODO implement me
	panic("implement me")
}

func (o *OllamaLLM) ShowModelList(ctx context.Context) []string {
	panic("implement me")
}

func InitOllama() *OllamaLLM {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1:8b"))
	if err != nil {
		panic(err)
	}
	ollamaLLM := &OllamaLLM{
		llm: llm,
	}
	return ollamaLLM
}

func init() {
	llmmodel.InitLLMModel(InitOllama())
}
