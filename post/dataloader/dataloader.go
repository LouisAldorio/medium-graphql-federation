package dataloader

import (
	"context"
	"net/http"
	"time"
	"myapp/service"
)

type ctxKeyType string

const (
	ctxKey ctxKeyType = "dataloaders"
)

type loaders struct {
	Posts *PostsLoader
}

func DataloaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		wait := 10 * time.Millisecond
		max := 150

		ldrs.Posts = &PostsLoader{
			wait: wait,
			maxBatch: max,
			fetch: service.PostsLoader,
		}

		dataloaderCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dataloaderCtx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}