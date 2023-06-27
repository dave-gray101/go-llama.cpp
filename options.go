package llama

import (
	common "github.com/go-skynet/go-common"
)

// TODO: should this just be inlined over in llama.go now?
// var DefaultModelInitializationOptions common.InitializationOptions = common.InitializationOptions{
// 	ContextSize: 512,
// 	Seed:        0,
// 	F16Memory:   false,
// 	MLock:       false,
// 	Embeddings:  false,
// 	MMap:        true,
// 	LowVRAM:     false,
// }

// var MergeInitializationOptionsWithDefaults = common.GetMergeInitializationOptionsFnFromDefault(DefaultModelInitializationOptions)

var DefaultPredictOptions common.PredictTextOptions = common.PredictTextOptions{
	Seed:              -1,
	Threads:           4,
	Tokens:            128,
	Penalty:           1.1,
	Repeat:            64,
	Batch:             512,
	NKeep:             64,
	TopK:              40,
	TopP:              0.95,
	TailFreeSamplingZ: 1.0,
	TypicalP:          1.0,
	Temperature:       0.8,
	FrequencyPenalty:  0.0,
	PresencePenalty:   0.0,
	Mirostat:          0,
	MirostatTAU:       5.0,
	MirostatETA:       0.1,
	MMap:              true,
}

var MergePredictOptionsWithDefaults = common.GetMergePredictTextOptionsFnFromDefault(DefaultPredictOptions)
