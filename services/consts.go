package services

// 转账状态
type TransferStatus uint8

const (
	// 转账失败
	TransferStatusFailure TransferStatus = -1
	// 余额不足
	TransferStatusNotEnough TransferStatus = 0
	// 转账成功
	TransferStatusSuccess TransferStatus = 1
)

//转账类型，0创建账户 >=1 进账 <=-1 支出
type ChangeType uint8

const (
	// 账户创建
	AccountCreated ChangeType = 0
	// 储值
	AccountStoreValue ChangeType = 1
	// 红包资金支出
	EnvelopeOut ChangeType = -2
	// 红包资金收入
	EnvelopeIn ChangeType = 2
	// 红包资金过期退款
	EnvelopeExpire ChangeType = 3
)

//转账标记
type ChangeFlag uint8

const (
	// 账户创建
	AccountCreatedFlag ChangeFlag = 0
	// 账户支出
	AccountOut ChangeFlag = -1
	// 账户收入
	AccountIn ChangeFlag = 1
)
