package llmmodel

import "context"

// LLM 基础功能接口定义
type LLMModel interface {
	GenerateFromSinglePrompt(ctx context.Context, promopt string)
	ShowModelList(ctx context.Context) []string
}

var llmModel LLMModel

func InitLLMModel(llmModelImpl LLMModel) {
	llmModel = llmModelImpl
}

func GetLLMModel() LLMModel {
	return llmModel
}
