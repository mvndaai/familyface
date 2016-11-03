name=familyface

build:
	env GOOS=darwin  GOARCH=amd64 go build -o bin/$(name).darwin $(name).go
	env GOOS=linux   GOARCH=amd64 go build -o bin/$(name).linux  $(name).go
	env GOOS=windows GOARCH=amd64 go build -o bin/$(name).exe    $(name).go
