package telegram

import (
	"context"
	"log/slog"

	"github.com/Arh0rn/gotgaibot1/internal/llm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const maxHistory = 10

//TODO: change this "kolhoz" in future

type Handler struct {
	Bot     *tgbotapi.BotAPI
	LLM     llm.LLM
	Context context.Context
	Logger  *slog.Logger
	Memory  map[int64][]llm.Message
}

func New(bot *tgbotapi.BotAPI, llmClient llm.LLM, ctx context.Context, logger *slog.Logger) *Handler {
	return &Handler{
		Bot:     bot,
		LLM:     llmClient,
		Context: ctx,
		Logger:  logger,
		Memory:  make(map[int64][]llm.Message),
	}
}

func (h *Handler) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.Bot.GetUpdatesChan(u)

	h.Logger.Info("Telegram bot started")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		userMsg := update.Message.Text

		h.Logger.Info("Received message", "chat_id", chatID, "text", userMsg)

		if _, exists := h.Memory[chatID]; !exists {
			h.Memory[chatID] = []llm.Message{
				{
					Role:    llm.RoleSystem,
					Content: h.LLM.GetLegend(),
				},
			}
		}

		h.Memory[chatID] = append(h.Memory[chatID], llm.Message{
			Role:    llm.RoleUser,
			Content: userMsg,
		})

		if len(h.Memory[chatID]) > maxHistory+1 {
			systemMsg := h.Memory[chatID][0]
			slog.Info("Old messages removed",
				"message", h.Memory[chatID][1],
			)
			recent := h.Memory[chatID][len(h.Memory[chatID])-maxHistory:]
			h.Memory[chatID] = append([]llm.Message{systemMsg}, recent...)
		}

		answer, err := h.LLM.GenerateResponse(h.Context, h.Memory[chatID])
		if err != nil {
			h.Logger.Error("LLM error", "error", err)
			answer = "Generation error."
		} else {
			h.Memory[chatID] = append(h.Memory[chatID], llm.Message{
				Role:    llm.RoleAssistant,
				Content: answer,
			})
		}

		msg := tgbotapi.NewMessage(chatID, answer)
		if _, err := h.Bot.Send(msg); err != nil {
			h.Logger.Error("Failed to send message", "error", err)
		}
	}
}
