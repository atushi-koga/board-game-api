## AWS Lambda でgoを実行する

- .goファイルを作成

$ vi hello.go

- .goファイルのバイナリを作成（バイナリはLambda上で実行できるようにクロスコンパイル設定値を含めてビルドする）
- バイナリファイルに権限をつける
- zip化

```
GOARCH=amd64 GOOS=linux go build hello.go && chmod 777 hello &&  zip handler.zip ./hello
```

- zipファイルをコンソールからアップロード

- コンソールのハンドラ名を、lambda.Startで呼ばれる関数名に合わせる

- 保存して実行