package entity

import "errors"

type ChatConfig struct {
	Model            *Model
	Temperature      float32
	TopP             float32
	N                int
	Stop             []string
	MaxTokens        int
	PresencePenalty  float32
	FrequencyPenalty float32
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMesages        []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == "ended" {
		return errors.New("chat is already ended, no more messages allowed")
	}
	for {
		if c.Config.Model.GetMaxToken() >= m.GetQtdTokens()+c.TokenUsage {
			c.Messages = append(c.Messages, m)
		}
	}
}
