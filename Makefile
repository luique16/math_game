build:
	go build -o math_game cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -f math_game
