build:
	GOARCH=amd64 GOOS=linux go build -o go-consultor ./cmd/lambda/.
	zip go-consultor.zip go-consultor

clean:
	rm go-consultor go-consultor.zip