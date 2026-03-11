FROM alpine:latest
WORKDIR /app
COPY personal-site .
COPY ui/static ./ui/static
COPY ui/templ ./ui/templ
EXPOSE 4000
CMD ["./personal-site"]
