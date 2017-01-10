# ysok

## sweep: ファイル一括削除

1. [Releases · pinzolo/ysok](https://github.com/pinzolo/ysok/releases) から最新バージョンの自分の環境に合ったバイナリをダウンロードしてください。
1. ダウンロードしたバイナリを `ysok` にリネームして任意の場所に配置して下さい。 `PATH` が通った場所だと便利です。
1. [Tokens for Testing and Development \| Slack](https://api.slack.com/docs/oauth-test-tokens)から対象 Team のトークンを作成して下さい。ここでは、仮に `x-1234-abcd` とします。（本当はもっと長いです）
1. [https://slack\.com/api/auth\.test?token=x-1234-abcd](https://slack.com/api/auth.test?token=x-1234-abcd)にアクセスして `user_id` の値を確認して下さい。ここでは仮に `U1234` とします。
1. コマンドプロンプトやターミナルにて `ysok sweep -u U1234 -t x-1234-abcd` とすれば30日前より以前にアップロードしたファイルを全て削除します。（実際には3, 4で取得した値に置き換える必要があります）
