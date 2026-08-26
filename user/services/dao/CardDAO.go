package dao

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

const (
	netID = 12
	BIN   = 123456
)

type CardDTO struct {
	PAN  string `json:"PAN"`
	Exp  string `json:"exp"`
	Name string `json:"Name"`
	Cvv  string `json:"cvv"`
}

type CardDAO struct {
	Db *sql.DB
}

func NewCardDAO(db *sql.DB) *CardDAO {
	return &CardDAO{
		Db: db,
	}
}

func (dao *CardDAO) NewCard(ctx context.Context, account_number string) (*CardDTO, error) {
	tx, err := dao.Db.BeginTx(ctx, nil)

	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	var latest_pan int64
	latestPan_err := tx.QueryRowContext(ctx, "SELECT PAN FROM ORDER BY created_at DESC LIMIT 1").Scan(&latest_pan)
	if latestPan_err != nil {
		return nil, latestPan_err
	}
	max_id, conv_err := strconv.Atoi(fmt.Sprintf("%d", latest_pan)[8:14])
	if conv_err != nil {
		return nil, conv_err
	}

	var name string
	name_err := tx.QueryRowContext(ctx, "SELECT fullname FROM users WHERE email = (SELECT email FROM accounts WHERE accounts_number = $1)", account_number).Scan(&name)

	if name_err != nil {
		return nil, name_err
	}

	exp := time.Now().AddDate(3, 0, 0)

	card, card_err := NewCardInstance(netID, BIN, int64(max_id+1), exp, name)
	if card_err != nil {
		return nil, card_err
	}

	_, insert_err := tx.ExecContext(ctx, "INSERT (PAN, exp, card_account_number, card_active) INTO cards VALUES($1, $2, $3, $4)", card.PAN, card.Exp, account_number, true)

	if insert_err != nil {
		return nil, insert_err
	}

	return card, nil
}

// HELPER FUNCTIONS

func NewCardInstance(net int64, BIN int64, id int64, exp time.Time, name string) (*CardDTO, error) {

	first := fmt.Sprintf("%d%d%d", net, BIN, id)

	chksum := luhn_chksum(first)

	expiry := fmt.Sprintf("%02d/%s", exp.Month(), fmt.Sprintf("%d", exp.Year())[2:])

	return &CardDTO{
		PAN:  fmt.Sprintf("%s%d", first, chksum),
		Exp:  expiry,
		Name: name,
		Cvv:  calculateCVV(fmt.Sprintf("%s%d", first, chksum), expiry, name),
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

// calculate cvv from SHA256 sum. this is all a mock.
// definately has to change later
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

	return fmt.Sprintf("%s", cvv)
}

