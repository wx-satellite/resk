package accounts

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

// 数据库访问层
type AccountDao struct {
	runner *dbx.TxRunner
}

// 查询数据库持久化对象的单实例
func (dao *AccountDao) GetOne(accountNo string) (account *Account) {
	account = &Account{AccountNo: accountNo}
	// 根据 unique 或者 pk 字段查询
	ok, err := dao.runner.GetOne(account)
	if nil != err {
		logrus.Error(err)
		return nil
	}
	if !ok {
		logrus.Error(fmt.Sprintf("accountNo：%v not found", accountNo))
		return nil
	}
	return
}

// 根据用户id和账户类型获取持久化实例
func (dao *AccountDao) GetByUserId(uid string, accountType int) (account *Account) {
	account = &Account{}
	sql := "select * from account where user_id = ? and account_type = ?"
	ok, err := dao.runner.Get(account, sql, uid, accountType)
	if nil != err {
		logrus.Error(err)
		return nil
	}
	if !ok {
		logrus.Error(fmt.Sprintf("userId：%v，accountType：%v not found", uid, accountType))
		return nil
	}
	return
}

// 账户数据的插入
func (dao *AccountDao) Insert(account *Account) (id int64, err error) {
	rs, err := dao.runner.Insert(account)
	if err != nil {
		return
	}
	return rs.LastInsertId()
}

// 账户数据的更新
// amount 如果是负数则扣除余额，反之增加余额
func (dao *AccountDao) Update(accountNo string, amount decimal.Decimal) (rows int64, err error) {
	// balance >= -1 * CAST(? as decimal(30,6)) 乐观锁
	// 保证余额足够的时候更新，余额不足的时候sql语句不更新，以此保证 balance 为正数
	sql := "update account set balance += CAST(? as decimal(30,6)) where account_no = ? and balance >= -1 * CAST(? as decimal(30,6))"
	rs, err := dao.runner.Exec(sql, amount.String(), accountNo, amount.String())
	if err != nil {
		logrus.Error(err)
		return
	}
	return rs.RowsAffected()
}

// 账户状态更新
func (dao *AccountDao) UpdateStatus(accountNo string, status int64) (rows int64, err error) {
	sql := "update account set status = ? where account_no = ?"
	rs, err := dao.runner.Exec(sql, status, accountNo)
	if err != nil {
		logrus.Error(err)
		return
	}
	return rs.RowsAffected()
}
