export PATH := $(GOPATH)/bin:$(PATH)

LDFLAGS := -s -w

all: build

build: app

app:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./bin/rdc_linux_amd64 ./cmd/rdc
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./bin/rds_linux_amd64 ./cmd/rds
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./bin/rdc_darwin_amd64 ./cmd/rdc
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./bin/rds_darwin_amd64 ./cmd/rds
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./bin/rdc_windows_amd64 ./cmd/rdc
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./bin/rds_windows_amd64 ./cmd/rds
	env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$(LDFLAGS)" -o ./bin/rdc_linux_mipsle ./cmd/rdc
	env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$(LDFLAGS)" -o ./bin/rds_linux_mipsle ./cmd/rds
