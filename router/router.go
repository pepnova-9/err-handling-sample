package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pepnova-9/err-handling-sample/controller"
	"github.com/pepnova-9/err-handling-sample/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Sample1
	// 外部スタック付きerrorsパッケージを使わない。各レイヤーでerrorをラップしてメッセージを付与していく
	sample1 := controller.Sample1{}
	r.HandleFunc("/sample1/users/{userID}", sample1.UpdateUserHandler).Methods(http.MethodPut)

	// Sample2
	// 外部スタック付きpkg/errorsパッケージを使う: 全てのレイヤーでpkg/errors.Wrapしていく方法
	// メリット:
	// - とりあえずpkg/errors.Wrapで必ずラップすれば良いというルールさえ守れば、出力が冗長でもスタックトレースは全部吐くことができる。
	// デメリット:
	// - どのレイヤーでもpkg/errors.Wrapしないといけない。もしpkg外のエラーでラップするとスタックトレースが吐けなくなる。(errorのループ回してerrors.withStackを遡る方法もある)
	// - ラップしたerrorごとにスタックトレースが吐き出されるので出力が見辛い
	sample2 := controller.Sample2{}
	r.HandleFunc("/sample2/users/{userID}", sample2.UpdateUserHandler).Methods(http.MethodPut)

	// Sample3
	// 外部スタック付きpkg/errorsパッケージを使う: 最下位のレイヤーでのみpkg/errorsでラップする。それ以外ではreturn errする方法
	// メリット:
	// - Sample2と違って最下位のスタックトレースのみが出力されるのでみやすい。
	// デメリット:
	// - pkg/errorsでラップし忘れたらスタックトレースも出ないし、間のレイヤーでreturn errしているので悲惨なことになる。
	// - 間のレイヤーで別のerrorでラップしちゃうとスタックトレースが吐けない。これはSample2と同じでerrorをループ回してerrors.withStackまで遡る方法で回避はできそう。
	// - ユースケース層でのエラーやドメイン層でのエラーの時に困る。どこが"最下位"なのか考えてエラーを返さないといけない。一方でどこが"最下位でないのか"を意識しないといけない。
	// - ユースケースやドメインでのエラーは"既知(ハンドリングできる)"エラーだからスタックトレースはいらないかもしれないが、その場合は間のレイヤーでメッセージつけてラップしたくなる。コードベースで統一感を出すのが難しくチームでの共通認識を取るコストがかかる。
	// - この方法を取るなら、エラーログ出力する際にはerrorのループを回してerrors.withStackがあればスタックトレースを吐く仕組みがないと厳しい印象
	sample3 := controller.Sample3{}
	r.HandleFunc("/sample3/users/{userID}", sample3.UpdateUserHandler).Methods(http.MethodPut)

	r.Use(middleware.RecoverPanic, middleware.CheckRequestContext)
	return r
}
