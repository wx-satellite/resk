package services

import "time"

type AccountService interface {
	// 账户创建
	CreateAccount(dto AccountCreateDTO) (*AccountDTO, error)
	// 转账
	Transfer(dto AccountTransferDTO) (TransferStatus, error)
	// 储值
	StoreValue(dto AccountTransferDTO) (TransferStatus, error)
	// 红包账户查询
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
}

// 账户创建信息
type AccountCreateDTO struct {
	UserId       string
	Username     string
	AccountType  int
	AccountName  string
	CurrencyCode string
	Amount       string // 浮点数会进度丢失，使用字符串传递
}

// 账户信息
type AccountDTO struct {
	AccountCreateDTO
	AccountNo string
	CreatedAt time.Time
}

// 账户交易的参与者
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

// 转账信息
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	Amount      string
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Desc        string
}
