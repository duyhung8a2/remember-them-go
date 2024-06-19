introJet:
	jet -dsn="file://\Repositories\remember-them-go\database.db" -path=./gen

genBob:
	go run github.com/stephenafamo/bob/gen/bobgen-sqlite@latest -c ./bobgen.yaml

initDb:
	rm db/database.db 2>/dev/null && sqlite3 db/database.db < db/db_schema.sql && sqlite3 db/database.db < db/db_init.sql 

resetDb: initDb genBob


install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	~/go/bin/swagger generate spec -o ./swagger.yaml --scan-models