.PHONY: start
start: .compose-build .compose-up .migrate

.PHONY: .compose-build
.compose-build:
	docker-compose build

.PHONY: .compose-up
.compose-up:
	docker-compose up -d

.PHONY: stop
stop: .compose-stop

.PHONY: .compose-stop
.compose-stop:
	docker-compose stop
	
.PHONY: migrate
migrate: .migrate

.PHONY: .migrate
.migrate:
	 goose -dir ./migrations postgres "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

.PHONY: build
build: vendor-proto .generate .build

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-vacancy-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-vacancy-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-vacancy-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-vacancy-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				--validate_out lang=go:pkg/ocp-vacancy-api \
				api/ocp-vacancy-api/ocp_vacancy_api.proto
		mv pkg/ocp-vacancy-api/github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api/* pkg/ocp-vacancy-api/
		rm -rf pkg/ocp-vacancy-api/github.com
		mkdir -p cmd/ocp-vacancy-api
		cd pkg/ocp-vacancy-api && ls go.mod || go mod init github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api && go mod tidy

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-vacancy-api cmd/ocp-vacancy-api/main.go

.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/grpc-server/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-vacancy-api
		cp api/ocp-vacancy-api/ocp_vacancy_api.proto vendor.protogen/api/ocp-vacancy-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate