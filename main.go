package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{}

	b, err := bot.New(os.Getenv("BOT_API"), opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
	ticker := time.Tick(time.Hour)
	chatIDStr := os.Getenv("CHAT_ID")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		panic(err)
	}
	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Бот запущен"})
	if err != nil {
		panic(err)
	} else {
		log.Println(msg)
	}
	for range ticker {
		now := time.Now()
		weekday := now.Weekday()
		if weekday == time.Tuesday {
			p := &bot.SendPollParams{
				ChatID:      chatID,
				Question:    "Идешь на квиз?",
				Options:     []models.InputPollOption{{Text: "Я"}, {Text: "Не я"}, {Text: "Я кемелжемелевич"}},
				IsAnonymous: bot.False(),
			}

			msg, err := b.SendPoll(ctx, p)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(msg)
			}
		}
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
