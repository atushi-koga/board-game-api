## AWS Lambda でgoを実行する

- スクリプトを作成（hello.go）

- バイナリファイルを作成（hello）
the binary for Lambda is run on Amazon Linux
バイナリはLambda上で実行できるようにクロスコンパイル設定値を含めてビルドする。
GOARCH=amd64 GOOS=linux go build hello.go 

- バイナリファイルに権限をつける
chmod 777 hello

- zip化
zip handler.zip ./hello

- zipファイルをコンソールからアップロード

- コンソールのハンドラ名を、lambda.Startで呼ばれる関数名に合わせる

-　保存して実行