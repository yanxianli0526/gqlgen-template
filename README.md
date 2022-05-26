# go-gql-server

這是用 Golang GraphQL API Server (template)

### 設定 env

可以參直接修改`config/env.go`(這邊要注意一下 env 裡面的 HTTPPort 記得要用 4000 不然可能會遇到其他的小麻煩)

### postgresql 設定

確認 config/env.go 裡面的變數和本機端一樣(帳號,密碼,資料庫名稱,port...等)
如果有需要migration 可以在database/migration/jobs裡實作seed,實作完成後在可以在database/migration/main呼叫實作的seed(可以參考main裡面的gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration)

### 執行方式

1. `make run` 
2. 用網頁打開 http://localhost:4000 
3. 在網頁上使用這段`{categories{ id, name }}`進行查詢,有查到就畢業了.沒查到可以確認一下資料庫裡面的categories有沒有資料,可能是auto migrate有問題
### How to modify schema

1. 修改 `internal/gql/schemas`
2. 修改 `internal/models`
3. 修改 `gqlgen.yml`
4. `make graph-build` (這一步很重要 它會透過 schema.graphql 跟 gqlgen.yml,進而生成三個檔案
5. 修改 GQLGen 產生的 Resolvers
6. 可以參考 resolvers/Menu 和 models/Menu 裡面有新增的實作



