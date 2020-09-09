package fetcher

import (
	"context"
	"github.com/PuerkitoBio/goquery"
)

// Fetch は受け取ったURLからタイトルを取得して返す
func Fetch(ctx context.Context, src string) (string, error) {
	doc, err := goquery.NewDocument(src)
    if err != nil {
        return src, err
	}
	var title string
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		title = s.Find("title").Text()
	})

	return title, nil
}
