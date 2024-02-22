
.PHONY: default
default: test

.PHONY: test
test:
	go test -count=1 -race ./...

.PHONY: bin
bin:
	go build -trimpath -o bin/kolko

#wyczysz zaleznosci i indeksuj
.PHONY: tidy
tidy:
	go mod tidy

#sciagnij zaleznosci
.PHONY: vendor
vendor:
	go mod vendor

.PHONY: dep
dep: tidy vendor

#updejtuj zaleznosci (liby uzywane w projekcie)
.PHONY: get
get:
	go get -u ./...

.PHONY: update
update: get tidy vendor
