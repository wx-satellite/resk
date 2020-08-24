package accounts

import (
	"github.com/shopspring/decimal"
	"resk/services"
	"time"
)

type AccountLog struct {
	Id              int64               `db:"id,omitempty"`
	LogNo           string              `db:"log_no,uni"`
	TradeNo         string              `db:"trade_no"`
	UserId          string              `db:"user_id"`
	AccountNo       string              `db:"account_no"`
	Username        string              `db:"username"`
	TargetAccountNo string              `db:"target_account_no"`
	TargetUserId    string              `db:"target_user_id"`
	TargetUsername  string              `db:"target_username"`
	Amount          decimal.Decimal     `db:"amount"`
	Balance         decimal.Decimal     `db:"balance"`
	ChangeType      services.ChangeType `db:"change_type"`
	ChangeFlag      services.ChangeFlag `db:"change_flag"`
	Status          int                 `db:"status"`
	Desc            string              `db:"desc"`
	CreatedAt       *time.Time          `db:"created_at,omitempty"`
}
