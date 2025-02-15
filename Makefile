gen:
	go run generator.go

clean:
	find . -name "gen_*.go" -type f -delete
