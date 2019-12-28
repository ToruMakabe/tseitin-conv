APP=tseitin-conv

.PHONY: build
build: clean
	go build -o ${APP} main.go

.PHONY: clean
clean:
	go clean

.PHONY: build-formula
build-formula:
	${MAKE} -C formula build