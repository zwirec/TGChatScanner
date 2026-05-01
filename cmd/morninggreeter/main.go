package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	TGBotAPI "github.com/zwirec/TGChatScanner/TGBotApi"
)

const (
	greetingHour   = 10
	greetingMinute = 0
	timezone       = "Europe/Moscow"
	greeting       = "Привет! Доброе утро ☀️"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	chatIDRaw := os.Getenv("CHAT_ID")
	if chatIDRaw == "" {
		log.Fatal("CHAT_ID is not set")
	}
	chatID, err := strconv.ParseInt(chatIDRaw, 10, 64)
	if err != nil {
		log.Fatalf("CHAT_ID must be an integer: %v", err)
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("failed to load timezone %s: %v", timezone, err)
	}

	api := TGBotAPI.NewBotAPI(token)
	log.Printf("morning greeter started: chat=%d tz=%s at %02d:%02d", chatID, timezone, greetingHour, greetingMinute)

	for {
		next := nextGreetingTime(time.Now().In(loc), greetingHour, greetingMinute)
		wait := time.Until(next)
		log.Printf("next greeting at %s (in %s)", next.Format(time.RFC3339), wait.Truncate(time.Second))
		time.Sleep(wait)

		if err := sendGreeting(api, chatID, greeting); err != nil {
			log.Printf("send failed: %v", err)
		} else {
			log.Printf("greeting sent")
		}
	}
}

func nextGreetingTime(now time.Time, hour, minute int) time.Time {
	candidate := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
	if !candidate.After(now) {
		candidate = candidate.AddDate(0, 0, 1)
	}
	return candidate
}

func sendGreeting(api *TGBotAPI.BotAPI, chatID int64, text string) error {
	resp, err := api.SendMessage(chatID, text, true)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return &telegramError{Status: resp.StatusCode, Body: string(body)}
	}
	return nil
}

type telegramError struct {
	Status int
	Body   string
}

func (e *telegramError) Error() string {
	return "telegram api error: status=" + strconv.Itoa(e.Status) + " body=" + e.Body
}
