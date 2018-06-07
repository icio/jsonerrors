README.md: *.go cmd/jsonerrors/main.go Gopkg.*
	@tmp=`mktemp`; \
	trap 'rm $$tmp' EXIT; \
	go run cmd/jsonerrors/main.go > $$tmp && mv $$tmp README.md
