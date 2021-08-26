module github.com/ozoncp/ocp-vacancy-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.1 // indirect
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.6.0
	github.com/golang/snappy v0.0.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api v0.0.0-00010101000000-000000000000
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/rs/zerolog v1.23.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.2.3 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ozoncp/ocp-vacancy-api/pkg/ocp-vacancy-api => ./pkg/ocp-vacancy-api
