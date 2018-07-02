build:
	go build && ./super

test:
	go test ./... -cover

db:
	sqlite3 data/sqlite.db