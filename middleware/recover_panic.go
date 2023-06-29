package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// panicが発生したらrecoverして500を返す
			if e := recover(); e != nil {
				var err error
				switch e := e.(type) {
				case error:
					err = e
				default:
					err = fmt.Errorf("%v", e)
				}

				// 予期しないエラーなのでエラーログを出力する
				// もしpkg/errorsのエラーなら+vで出力すればスタックがログとして出せる。
				// 他のエラーでラップされていたら遡って出力しないとスタックトレースが出力されない。
				// Sentryは一部のスタック付きerrorパッケージの仕様を汲み取っているのでスタックがSentryに送信されている。(errorを遡って出力してくれている)
				// https://github.com/getsentry/sentry-go/blob/master/stacktrace.go#L74-L101
				// Sentryを使っていない場合はスタックをログとして出力してあげないと追えなくなる。
				log.Printf("panic error: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
