package dao

import (
	"context"
	"testing"

	"go-common/app/service/main/member/model"

	"github.com/smartystreets/goconvey/convey"
)

func TestDaoExp(t *testing.T) {
	var (
		c   = context.TODO()
		mid = int64(4780461)
	)
	convey.Convey("Exp", t, func(ctx convey.C) {
		exp, err := d.Exp(c, mid)
		ctx.Convey("Error should be nil", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("exp should not be nil", func(ctx convey.C) {
			t.Logf("exp:%+v", exp)
			ctx.So(exp, convey.ShouldNotBeNil)
		})
	})
}

func TestDaoExps(t *testing.T) {
	var (
		c    = context.TODO()
		mids = []int64{4780461, 2}
	)
	convey.Convey("Exps", t, func(ctx convey.C) {
		exps, err := d.Exps(c, mids)
		ctx.Convey("Error should be nil", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("exps should not be nil", func(ctx convey.C) {
			ctx.So(exps, convey.ShouldNotBeNil)
		})
	})
}

func TestDaoExpLog(t *testing.T) {
	var (
		c   = context.TODO()
		mid = int64(0)
		ip  = ""
	)
	convey.Convey("ExpLog", t, func(ctx convey.C) {
		l, err := d.ExpLog(c, mid, ip)
		ctx.Convey("Error should be nil", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("p1 should not be nil", func(ctx convey.C) {
			ctx.So(l, convey.ShouldNotBeNil)
		})
	})
}

func TestDaoasExpLog(t *testing.T) {
	var (
		res = &model.SearchResult{}
	)
	convey.Convey("asExpLog", t, func(ctx convey.C) {
		p1 := asExpLog(res)
		ctx.Convey("p1 should not be nil", func(ctx convey.C) {
			t.Logf("log:%+v", p1)
			ctx.So(p1, convey.ShouldNotBeNil)
		})
	})
}
