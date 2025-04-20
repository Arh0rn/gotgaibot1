package app

import (
	"context"
	"fmt"
	"github.com/Arh0rn/gotgaibot1/internal/delivery/telegram"
	"github.com/Arh0rn/gotgaibot1/internal/llm"
	"github.com/Arh0rn/gotgaibot1/internal/llm/openai"
	"github.com/Arh0rn/gotgaibot1/pkg/config"
	"github.com/Arh0rn/gotgaibot1/pkg/logger"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

type App struct {
	cfg *config.Config
	ctx context.Context
	log *slog.Logger

	bot *tgbotapi.BotAPI

	llm llm.LLM
}

func NewApp(ctx context.Context) (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	log := logger.InitLogger(cfg.Env)
	slog.SetDefault(log)
	log.Debug(fmt.Sprintf("%+v", cfg))

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramConfig.Token)
	if err != nil {
		return nil, err
	}
	log.Debug(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	model := openai.New(cfg.LLMConfig)

	app := &App{
		cfg: cfg,
		ctx: ctx,
		log: log,
		bot: bot,
		llm: model,
	}

	return app, nil
}

func (a *App) Run() error {
	handler := telegram.New(a.bot, a.llm, a.ctx, a.log)
	handler.Run()
	return nil
	//TODO: add graceful shutdown
}
