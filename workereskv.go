package main

import (
	"context"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

type KV struct {
	API  *cloudflare.API
	ID   string
}

func NewKV() *KV {
	api, err := cloudflare.New(os.Getenv("CF_API"), os.Getenv("CF_EMAIL"), cloudflare.UsingOrganization(os.Getenv("CF_ORG_ID")))
	if err != nil {
		panic(err)
	}

	return &KV{
		API:  api,
		ID:   os.Getenv("CF_WORKERSKV_ID"),
	}
}

func (k *KV) write(key, value string) error {
	_, err := k.API.WriteWorkersKV(context.Background(), k.ID, key, []byte(value))
	return err
}

func (k *KV) read(key string) ([]byte, error) {
	return k.API.ReadWorkersKV(context.Background(), k.ID, key)
}
