run:
	templ generate
	go run ./cmd/web

build:
	templ generate
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o personal-site ./cmd/web/

deploy-staging:
	fly deploy --config fly.staging.toml

deploy-production:
	fly deploy

deploy: deploy-production
