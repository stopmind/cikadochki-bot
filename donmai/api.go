package donmai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	httpClient *http.Client
	domain     string
}

func (a *Api) GetPosts(tagsQuery string, count int, page int) ([]Post, error) {
	resp, err := a.httpClient.Get(fmt.Sprintf("https://%s/posts.json?tags=%s&limit=%d&page=%d", a.domain, tagsQuery, count, page))
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var posts []Post
	if err = json.Unmarshal(data, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func NewApi(domain string) Api {
	return Api{
		httpClient: http.DefaultClient,
		domain:     domain,
	}
}
