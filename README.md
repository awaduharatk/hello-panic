# Golangのpanicとrecoverのmemo

- panicが発生する関数の後に宣言したDeferは実行されない

- 一回recoverした後に再度panicを発生させることが可能