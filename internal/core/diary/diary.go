package diary

import (
	"emwell/internal/core/diary/converter"
	"emwell/internal/core/diary/repository"
	"emwell/internal/logger"
)

type Diary struct {
	logger        logger.ILogger
	converter     converter.Converter
	emotionalRepo repository.EmotionalDiaryRepository
}

func NewDiary(logger logger.ILogger, emotionalRepo repository.EmotionalDiaryRepository) *Diary {
	return &Diary{
		logger:        logger,
		emotionalRepo: emotionalRepo,
	}
}
