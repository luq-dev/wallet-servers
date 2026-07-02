module message

go 1.26.4

require (
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/gorilla/websocket v1.5.3
	user v0.0.0
)

require github.com/lib/pq v1.12.3 // indirect

replace user => ../user
