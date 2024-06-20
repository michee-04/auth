module github.com/michee/authentificationApi

go 1.21.10

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v1.5.5
	github.com/go-chi/jwtauth/v5 v5.3.1
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.24.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.10

	// Dépendances indirectes suivantes
	filippo.io/edwards25519 v1.1.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/go-sql-driver/mysql v1.8.1
	github.com/goccy/go-json v0.10.3
	github.com/jackc/pgpassfile v1.0.0
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a
	github.com/jackc/pgx/v5 v5.5.5
	github.com/jackc/puddle/v2 v2.2.1
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5
	github.com/lestrrat-go/blackmagic v1.0.2
	github.com/lestrrat-go/httpcc v1.0.1
	github.com/lestrrat-go/httprc v1.0.5
	github.com/lestrrat-go/iter v1.0.2
	github.com/lestrrat-go/jwx/v2 v2.1.0
	github.com/lestrrat-go/option v1.0.1
	github.com/segmentio/asm v1.2.0
	golang.org/x/sync v0.7.0
	golang.org/x/sys v0.21.0
	golang.org/x/text v0.16.0
)