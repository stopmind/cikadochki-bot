package main

import (
	"cikadochki-bot/bot"
	"cikadochki-bot/donmai"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	b, err := bot.New(os.Getenv("TELEGRAM_APITOKEN"), "data.json")
	if err != nil {
		panic(err)
	}

	api := donmai.NewApi("safebooru.donmai.us")

	b.Sources = []bot.Source{
		bot.NewDonmai(&api, "higurashi_no_naku_koro_ni"),
	}

	go func() {
		for {
			err := b.SendImages()
			if err != nil {
				println(err.Error())
			}
			time.Sleep(time.Hour * 24)
		}
	}()

	go b.Run()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	err = b.Stop()
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
