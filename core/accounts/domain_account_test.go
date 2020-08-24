package accounts

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"resk/services"
	_ "resk/testx"
	"testing"
)

func TestAccountDomain_Create(t *testing.T) {
	dto := &services.AccountDTO{}
	dto.Username = ksuid.New().Next().String()
	dto.UserId = ksuid.New().Next().String()
	dto.Balance = decimal.NewFromFloat(100)
	dto.Status = 1

	domain := new(accountDomain)
	Convey("账户创建", t, func() {
		r, err := domain.Create(*dto)

		So(err, ShouldBeNil)
		So(r, ShouldNotBeNil)
		So(r.Balance.String(), ShouldEqual, dto.Balance.String())
		So(r.UserId, ShouldEqual, dto.UserId)
		So(r.Username, ShouldEqual, dto.Username)
		So(r.Status, ShouldEqual, dto.Status)
	})
}
