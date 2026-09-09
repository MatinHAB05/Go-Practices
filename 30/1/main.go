package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/spf13/viper"
)

type Config struct {
	Token string `mapstructure:"TOKEN"`
}

func main() {
	// Retrieve the API token from environment variables
	viper.AddConfigPath("./../")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("ERR: Failed to read config file: ", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("ERR: Failed to unmarshal config: ", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// 1. Initial configurations and default fall-back handler
	opts := []bot.Option{
		bot.WithDefaultHandler(notFoundHandler),
		bot.WithDebug(),
	}

	b, err := bot.New(cfg.Token, opts...)
	if err != nil {
		panic(err)
	}

	// 2. Register routes (Routing Table)
	// Exact match route for /start
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)

	// Exact match route for /help
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, helpHandler)

	// Prefix match route for inline buttons (matches any CallbackData starting with "select_")
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "select_", bot.MatchTypePrefix, callbackHandler)

	fmt.Println("Bot is running...")
	b.Start(ctx)
}

// ----------------- HANDLERS -----------------

// 1. /start command handler (includes sending an inline keyboard)
func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// Build inline keyboard
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Option 1 🟢", CallbackData: "select_opt1"},
				{Text: "Option 2 🔵", CallbackData: "select_opt2"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Hello! Welcome to the test bot. Please select an option:",
		ReplyMarkup: kb,
	})
}

// 2. /help command handler
func helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Bot commands:\n/start - Start and view options\n/help - Show help message",
	})
}

// 3. Inline button handler (Callback Routing)
func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// Stop the loading state on the user's client button
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "Selection recorded!",
	})

	// Extract clicked callback data
	data := update.CallbackQuery.Data

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   fmt.Sprintf("You clicked the button with data: `%s` 🔥", data),
	})
}

// 4. Unknown/unhandled message handler (Not Found / Default)
func notFoundHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Command or message not recognized. Please use /help for available options.",
	})
}
