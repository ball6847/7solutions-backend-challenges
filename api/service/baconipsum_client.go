package service

import (
	"errors"
	"time"

	"github.com/imroc/req/v3"
	"github.com/patrickmn/go-cache"
)

type BaconipsumClient struct {
	http  *req.Client
	cache *cache.Cache
}

func NewBaconipsumClient(http *req.Client) *BaconipsumClient {
	return &BaconipsumClient{
		http:  http,
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (bacon *BaconipsumClient) Get(t string) (string, error) {
	if t != "meat-and-filler" && t != "all-meat" {
		return "", errors.New("invalid meat type for baconipsum api")
	}

	if val, found := bacon.cache.Get("BaconipsumClient." + t); found {
		return val.(string), nil
	}

	resp, err := bacon.http.R().
		Get("https://baconipsum.com/api/?type=" + t + "&paras=99&format=text")
	if err != nil {
		return "", err
	}

	text := resp.String()

	bacon.cache.Set("BaconipsumClient."+t, text, cache.DefaultExpiration)

	return text, nil
}
