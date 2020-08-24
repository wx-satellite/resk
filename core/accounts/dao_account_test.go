package accounts

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"resk/infra/base"
	_ "resk/testx"
	"testing"
)

func TestAccountDao_GetOne(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		Convey("通过AccountNo查询账户数据", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().String(),
				AccountName: "测试账户",
				UserId:      ksuid.New().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			na := dao.GetOne(a.AccountNo)

			So(na, ShouldNotBeNil)
			So(na.Balance.String(), ShouldEqual, a.Balance.String())
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)
		})
		return nil
	})

	if nil != err {
		logrus.Error(err)
	}

}

func TestAccountDao_GetByUserId(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		Convey("通过用户id和账户类型查询账户数据", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().String(),
				AccountName: "测试账户",
				UserId:      ksuid.New().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
				AccountType: 2,
			}
			id, err := dao.Insert(a)

			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, id)

			na := dao.GetByUserId(a.UserId, a.AccountType)
			So(na, ShouldNotBeNil)
			So(na.Balance.String(), ShouldEqual, a.Balance.String())
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)
		})
		return nil
	})

	if nil != err {
		logrus.Error(err)
	}
}

func TestAccountDao_Update(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		balance := decimal.NewFromFloat(100)
		Convey("更新账户余额", t, func() {
			a := &Account{
				Balance:     balance,
				Status:      1,
				AccountNo:   ksuid.New().String(),
				AccountName: "测试账户",
				UserId:      ksuid.New().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			// 增加余额
			Convey("增加余额", func() {
				amount := decimal.NewFromFloat(10)
				rows, err := dao.Update(a.AccountNo, amount)
				So(err, ShouldBeNil)
				So(rows, ShouldEqual, 1)

				na := dao.GetOne(a.AccountNo)
				So(na, ShouldNotBeNil)
				So(na.Balance.String(), ShouldEqual, balance.Add(amount).String())
			})
			// 扣减余额，余额充足
			Convey("扣减余额，余额充足", func() {
				a1 := dao.GetOne(a.AccountNo)
				So(a1, ShouldNotBeNil)

				change := decimal.NewFromFloat(-10)

				row, err := dao.Update(a.AccountNo, change)

				So(err, ShouldBeNil)
				So(row, ShouldEqual, 1)

				a2 := dao.GetOne(a.AccountNo)
				So(a2, ShouldNotBeNil)

				So(a2.Balance.String(), ShouldEqual, a1.Balance.Add(change).String())
			})
			// 扣减余额，余额不足
			Convey("扣减余额，余额不足", func() {
				a1 := dao.GetOne(a.AccountNo)
				So(a1, ShouldNotBeNil)

				change := decimal.NewFromFloat(-300)

				row, err := dao.Update(a.AccountNo, change)

				So(err, ShouldBeNil)
				So(row, ShouldEqual, 0)

				a2 := dao.GetOne(a.AccountNo)
				So(a2, ShouldNotBeNil)

				So(a2.Balance.String(), ShouldEqual, a1.Balance.String())
			})
		})
		return nil
	})
	if nil != err {
		logrus.Error(err)
	}
}
