package accounts

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"resk/infra/base"
	"resk/services"
	"testing"
)

func TestAccountLogDao_Get(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountLogDao{
			runner: runner,
		}
		Convey("通过logNo、tradeNo 查询", t, func() {
			logNo := ksuid.New().String()
			tradeNo := ksuid.New().String()
			a := &AccountLog{
				LogNo:      logNo,
				TradeNo:    tradeNo,
				Status:     1,
				AccountNo:  ksuid.New().String(),
				UserId:     ksuid.New().String(),
				Username:   "测试用户",
				Amount:     decimal.NewFromFloat(100),
				Balance:    decimal.NewFromFloat(100),
				ChangeFlag: services.AccountCreatedFlag,
				ChangeType: services.AccountCreated,
			}
			// 插入
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			Convey("根据 logNo 查询", func() {
				a1 := dao.GetOne(a.LogNo)
				So(a1, ShouldNotBeNil)
				So(a1.Id, ShouldEqual, id)
			})

			Convey("根据 tradeNo 查询", func() {
				a2 := dao.GetByTradeNo(tradeNo)
				So(a2, ShouldNotBeNil)
				So(a2.Id, ShouldEqual, id)
			})
		})

		return nil
	})

	if nil != err {
		logrus.Error(err)
	}
}
