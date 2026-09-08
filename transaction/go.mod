module transaction

go 1.26.4

require (
	storage v0.0.0
	user v0.0.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/lib/pq v1.12.3 // indirect
)

replace user => ../user

replace storage => ../storage
