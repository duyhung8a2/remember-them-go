introJet:
	jet -dsn="file://\Repositories\remember-them-go\database.db" -path=./gen

genBob:
	go run github.com/stephenafamo/bob/gen/bobgen-sqlite@latest -c ./bobgen.yaml