package accounts

import (
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"github.com/tietang/dbx"
	"resk/infra/base"
	"resk/services"
)

//领域实例是有状态，每次使用都需要先实例化

type accountDomain struct {
	account    Account
	accountLog AccountLog
}

// 创建 accountLog 的 logNo
func (domain *accountDomain) createAccountLogNo() {
	// 暂时采用ksuid的ID生成策略来创建No
	// 后期优化成可读性比较好的分布式ID
	// 全局唯一的ID
	domain.accountLog.LogNo = ksuid.New().Next().String()
}

// 创建 accountNo 的逻辑
func (domain *accountDomain) createAccountNo() {
	domain.account.AccountNo = ksuid.New().Next().String()
}

// 创建流水记录
func (domain *accountDomain) createAccountLog() {
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	// 账户创建的流水，tradeNo 和 logNo 一致，其他 tradeNo 由外部传入
	domain.accountLog.TradeNo = domain.accountLog.LogNo

	// 流水中的交易主体信息
	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.Username = domain.account.Username.String
	domain.accountLog.AccountNo = domain.account.AccountNo

	// 流水中的交易对象信息
	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetUsername = domain.account.Username.String
	domain.accountLog.TargetAccountNo = domain.account.AccountNo

	// 交易金额
	domain.accountLog.Amount = domain.account.Balance
	domain.accountLog.Balance = domain.account.Balance

	// 交易变化属性
	domain.accountLog.Desc = "账户创建"
	domain.accountLog.ChangeType = services.AccountCreated
	domain.accountLog.ChangeFlag = services.AccountCreatedFlag
}

// 创建账户的业务
func (domain *accountDomain) Create(dto services.AccountDTO) (d *services.AccountDTO, err error) {
	// 创建 account 持久化对象
	domain.account = Account{}
	domain.account.FromDTO(&dto)
	domain.createAccountNo()
	// 创建 accountLog 持久化对象
	domain.createAccountLog()

	// 创建 DAO 对象
	accountDao := AccountDao{}
	accountLogDao := AccountLogDao{}
	err = base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		accountLogDao.runner = runner

		// 插入账户数据
		id, err := accountDao.Insert(&domain.account)
		if nil != err {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户失败")
		}

		// 插入流水数据
		id, err = accountLogDao.Insert(&domain.accountLog)
		if nil != err {
			return err
		}
		if id <= 0 {
			return errors.New("创建流水失败")
		}

		domain.account = *accountDao.GetOne(domain.account.AccountNo)
		return nil
	})
	d = domain.account.ToDTO()
	return
}
