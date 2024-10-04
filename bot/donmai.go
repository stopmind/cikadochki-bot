package bot

import (
	"cikadochki-bot/donmai"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
)

type Donmai struct {
	api       *donmai.Api
	tagsQuery string
}

func (d *Donmai) GetImage() (tgbotapi.RequestFileData, error) {
	posts, err := d.api.GetPosts(d.tagsQuery, 20, 0)
	if err != nil {
		return nil, err
	}

	return tgbotapi.FileURL(posts[rand.Intn(len(posts))].FileUrl), nil
}

func NewDonmai(api *donmai.Api, tagsQuery string) *Donmai {
	return &Donmai{
		api:       api,
		tagsQuery: tagsQuery,
	}
}
