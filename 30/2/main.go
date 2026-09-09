package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Config struct {
	Token string `mapstructure:"TOKEN"`
}

func main() {
	customDebugHandler, cl := initLogger(false)
	defer cl()
	cfg := initConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// 3. Configure bot options using bot.WithDebugHandler
	opts := []bot.Option{
		bot.WithDebug(),                          // Enable debug mode
		bot.WithDebugHandler(customDebugHandler), // Redirect debug output
		bot.WithDefaultHandler(notFoundHandler),
	}

	b, err := bot.New(cfg.Token, opts...)
	if err != nil {
		log.Fatalf("ERR: Failed to create bot instance: %v", err)
	}

	// Register routes
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, helpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/salavat", bot.MatchTypeExact, salavatHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/credits", bot.MatchTypeExact, creditHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/progress", bot.MatchTypeExact, progressHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "select_", bot.MatchTypePrefix, callbackHandler)

	fmt.Println("Bot is running...")
	b.Start(ctx)
}

// ----------------- HANDLERS -----------------

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Bot commands:\n/start - Start and view options\n/help - Show help message",
	})

}

func salavatHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "الله هم صله الله محمد و آل محمد",
	})

}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "Selection recorded!",
	})

	data := update.CallbackQuery.Data
	// m, _ :=
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   fmt.Sprintf("You clicked the button with data: `%s` 🔥", data),
	})

	// sm, _ := json.Marshal(m)
	// fmt.Println(string(sm))
}

func notFoundHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Command or message not recognized. Please use /help for available options.",
	})
}

func creditHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "MatinHAB05 - 2026",
	})
}

func progressHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	pg := 0
	rm, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   strconv.Itoa(pg) + " %",
	})
	log.Println(rm.ID)
	go func() {
		for {
			if pg >= 100 {
				return
			}
			time.Sleep(time.Second * time.Duration((rand.Int() % 3)))
			pg = pg + 5*(rand.Int()%6)
			// res, err :=
			b.EditMessageText(ctx, &bot.EditMessageTextParams{
				ChatID:    update.Message.Chat.ID,
				MessageID: rm.ID,
				Text:      strconv.Itoa(pg) + " %",
			})
			// log.Println(res, err)
		}
	}()
}
