# Penguin Park Server

https://github.com/chikuwa111/penguin-park のサーバー側です。

ローカル開発する際の注意点
- websocket/handler.goのallowOriginsを更新する
- main.goのListenAndServeTLSをListenAndServeにする
