run:
	templ generate
	go run ./cmd/web

build:
	templ generate
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o personal-site ./cmd/web/

deploy:
	make build
	fly deploy
