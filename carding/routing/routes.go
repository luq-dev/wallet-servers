package routing

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	. "storage/database"
	"strconv"
	"time"
)

func RegisterRoutes() {
	http.HandleFunc("POST /", CardDetails)
	http.HandleFunc("POST /cards", AssignCard)
}

const (
	NET_ID = 9
	BIN    = 27041002
	CIPHER_KEY = ")7!$@#$&%^%"
)


// Respond with the user's card
func CardDetails(w http.ResponseWriter, r *http.Request) {

	var account_number string
	// var card Card

	DB.QueryRow("SELECT 1 FROM user_cards WHERE user = $1", account_number)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

// assign a new payment card to an account
func AssignCard(w http.ResponseWriter, r *http.Request) {

	var id int64
	var exp time.Time
	var account_number string
	var name string

	DB.Begin()
	DB.QueryRow("SELECT unique_id, exp FROM cards WHERE unique_id = MAX(unique_id)").Scan(&id, &exp)
	DB.QueryRow("SELECT users.fullname FROM users JOIN accounts ON accounts.user_id = users.email WHERE accounts.account_number = $1", account_number).Scan(&name)

	makeCard(NET_ID, BIN, id, exp, name)
}

type Card struct {
	PAN  string
	exp  string
	name string
	cvv  string
}

func makeCard(net int64, BIN int64, id int64, exp time.Time, name string) (*Card, error) {

	first := fmt.Sprintf("%d%d%d", net, BIN, id)

	chksum := luhn_chksum(first)

	expiry := fmt.Sprintf("%02d/%s", exp.Month(), fmt.Sprintf("%d", exp.Year())[2:])

	return &Card{
		PAN:  fmt.Sprintf("%s%d", first, chksum),
		exp:  expiry,
		name: name,
		cvv:  calculateCVV(fmt.Sprintf("%s%d", first, chksum), expiry, name),
	}, nil
}

func luhn_chksum(s string) int {
	sum := 0

	for i := range len(s) {
		if i%2 == 0 {
			if n := strconv.Itoa(((int(s[i]) - 48) * 2)); len(n) > 1 {
				sum += (int(n[0]) - 48) + (int(n[1]) - 48)
			} else {
				sum += (int(n[0]) - 48)
			}
		} else {
			sum += (int(s[i]) - 48)
		}
	}

	return 10 - sum%10
}

// calculate cvv from SHA256 sum. this is all a mock. definately
// has to change later
func calculateCVV(PAN string, exp string, name string) string {

	data := fmt.Sprintf("%s%s%s", PAN, exp, name)
	var cvv []byte

	h := sha256.New()

	h.Write([]byte(data))

	bs := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(bs)

	for i := range len(bs) {
		if int(bs[i]) >= 48 || int(bs[i]) <= 57 {
			cvv = append(cvv, bs[i])
			if len(cvv) == 3 {
				break
			}
		}
	}

	return fmt.Sprintf("%s",cvv)
}
