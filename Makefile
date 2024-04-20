introJet:
	jet -dsn="file://\Repositories\remember-them-go\database.db" -path=./gen

genBob:
	go run github.com/stephenafamo/bob/gen/bobgen-sqlite@latest -c ./bobgen.yaml

initDb:
	rm database.db 2>/dev/null && sqlite3 database.db < db_schema.sql && sqlite3 database.db < db_init.sql 

resetDb: initDb genBob