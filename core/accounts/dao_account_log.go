package accounts

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type AccountLogDao struct {
	runner *dbx.TxRunner
}

// 通过流水编号查询流水记录
func (dao *AccountLogDao) GetOne(logNo string) (accountLog *AccountLog) {
	accountLog = &AccountLog{
		LogNo: logNo,
	}
	ok, err := dao.runner.GetOne(accountLog)
	if nil != err {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return
}

// 通过交易编号查询流水记录
func (dao *AccountLogDao) GetByTradeNo(tradeNo string) (accountLog *AccountLog) {
	accountLog = &AccountLog{}
	// * 实际需要替换成 需要查询的字段列表
	sql := "select * from account_log where trade_no = ?"
	ok, err := dao.runner.Get(accountLog, sql, tradeNo)
	if nil != err {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return
}

// 流水记录的写入
func (dao *AccountLogDao) Insert(accountLog *AccountLog) (id int64, err error) {
	res, err := dao.runner.Insert(accountLog)
	if nil != err {
		logrus.Error(err)
		return
	}
	return res.LastInsertId()
}
