package accounts

// 数据库持久层

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type Account struct {
	Id           int64           `db:"id,omitempty"`
	AccountNo    string          `db:"account_no,uni"`
	AccountName  string          `db:"account_name"`
	AccountType  int             `db:"account_type"`
	CurrencyCode string          `db:"currency_code"`
	UserId       string          `db:"user_id"`
	Username     sql.NullString  `db:"username"`
	Balance      decimal.Decimal `db:"balance"`
	Status       int             `db:"status"`
	CreatedAt    *time.Time      `db:"created_at,omitempty"`
	UpdatedAt    *time.Time      `db:"updated_at,omitempty"`
}
