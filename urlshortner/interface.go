package urlshortner

import "context"

type UrlShortnerService interface {
	encodeUrl(ctx context.Context, urlInput UrlEncoder) (string, error)
	decodeUrl(ctx context.Context, code string) string
}
