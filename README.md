


### Build

- Run `GOARCH=amd64 GOOS=linux go build cmd/main.go` at the project's source folder
- Run `mv main build/`

### Prepare for deployment

- Run `zip -jrm build/main.zip build/main`
