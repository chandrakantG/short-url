package urlshortner

import (
	"context"
	"errors"
	"math/rand"

	"short-url/base62"
)

const BASEURL = "http://localhost:8082/"

type urlService struct {
}

type Store struct {
	db map[uint64]urlData
}

func NewUrlStore() *Store {
	store := &Store{db: map[uint64]urlData{}}
	return store
}

func (k Store) Get(key uint64) urlData {
	return k.db[key]
}

func (k Store) Put(key uint64, value urlData) {
	k.db[key] = value
}

type urlData struct {
	ShortURL string `json:"short_url"`
	URL      string `json:"url"`
}

var storeUrlData = NewUrlStore()

func (us urlService) isUsed(checkUrl string) bool {
	for _, data := range storeUrlData.db {
		if data.URL == checkUrl {
			return true
		}
	}
	return false
}

func (us urlService) encodeUrl(ctx context.Context, urlInput UrlEncoder) (string, error) {
	if us.isUsed(urlInput.Url) {
		return "url already used", errors.New("url already used")
	}
	id := rand.Uint64()
	encodeStr := base62.Encode(id)
	var item urlData
	item.URL = urlInput.Url
	item.ShortURL = BASEURL + encodeStr
	storeUrlData.Put(id, item)
	return item.ShortURL, nil
}

func (us urlService) decodeUrl(ctx context.Context, code string) string {
	encodeNum, _ := base62.Decode(code)
	urlData := storeUrlData.Get(encodeNum)
	if urlData.URL != "" {
		return urlData.URL
	}
	return "url does not exist"
}

func NewService() UrlShortnerService {
	return urlService{}
}
