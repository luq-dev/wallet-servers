module transaction

go 1.26.4

require user v0.0.0

require github.com/golang-jwt/jwt/v5 v5.3.1 // indirect

replace user => ../user

replace message => ../message

replace finance => ../finance

replace storage => ../storage
