module transaction

go 1.26.4

require (
	message v0.0.0-00010101000000-000000000000
	user v0.0.0
	finance v0.0.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/lib/pq v1.12.3 // indirect
)

replace user => ../user
replace message => ../message
replace finance => ../finance
