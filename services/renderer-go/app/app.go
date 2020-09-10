package app

import (
	"context"

	pb_fetcher "github.com/wafuwafu/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

type App struct {
	FetcherClient pb_fetcher.FetcherClient
}

func NewApp(fetcherClient pb_fetcher.FetcherClient) *App {
	return &App{fetcherClient}
}

func (a *App) Fetch(ctx context.Context, src string) (string, error) {
	reply, err := a.FetcherClient.Fetcher(ctx, &pb_fetcher.FetcherRequest{Src: src})
	if err != nil {
		return "", err
	}
	return reply.Title, nil
}
