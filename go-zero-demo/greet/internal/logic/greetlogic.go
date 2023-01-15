package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	_ "github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.Ser
}

func (l *GreetLogic) Greet(req *types.Request) *ty {

}
