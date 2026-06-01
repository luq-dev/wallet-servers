module transaction

go 1.25.0

require (
	auth v0.0.0
	github.com/lib/pq v1.12.3
)

require github.com/golang-jwt/jwt/v5 v5.3.1 // indirect

replace auth => ../auth
