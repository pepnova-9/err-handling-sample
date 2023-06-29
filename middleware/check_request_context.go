package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func CheckRequestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
		if errors.Is(r.Context().Err(), context.Canceled) {
			fmt.Printf("RequestContext: context canceled \n")
		} else {
			fmt.Printf("RequestContext: %v \n", r.Context().Err())
		}

	})
}
