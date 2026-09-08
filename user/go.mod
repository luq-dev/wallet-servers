module user

go 1.26.4

require (
	github.com/golang-jwt/jwt/v5 v5.3.1
	golang.org/x/crypto v0.52.0
)

require storage v0.0.0

require github.com/lib/pq v1.12.3 // indirect

replace transaction => ../transaction

replace storage => ../storage
