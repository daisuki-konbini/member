package grpc

import (
	"context"
	"go-common/app/service/main/member/api/grpc/v1"
	"go-common/app/service/main/member/model"
)

// Exp get member exp info
func (s *MemberServer) Exp(ctx context.Context, req *v1.MidReq) (*v1.LevelInfoReply, error) {
	res, err := s.svr.Exp(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	var levelInfoReply = &v1.LevelInfoReply{
		Cur:     res.Cur,
		Min:     res.Min,
		NowExp:  res.NowExp,
		NextExp: res.NextExp,
	}

	return levelInfoReply, nil
}

// Level get member lebel info
func (s *MemberServer) Level(ctx context.Context, req *v1.MidReq) (*v1.LevelInfoReply, error) {
	res, err := s.svr.Level(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	var levelInfoReply = &v1.LevelInfoReply{
		Cur:     res.Cur,
		Min:     res.Min,
		NowExp:  res.NowExp,
		NextExp: res.NextExp,
	}

	return levelInfoReply, nil
}

// UpdateExp update member exp value
func (s *MemberServer) UpdateExp(ctx context.Context, req *v1.AddExpReq) (*v1.EmptyStruct, error) {
	err := s.svr.UpdateExp(ctx, &model.ArgAddExp{
		Mid:     req.Mid,
		Count:   req.Count,
		Reason:  req.Reason,
		Operate: req.Operate,
		IP:      req.Ip,
	})
	if err != nil {
		return nil, err
	}

	return &v1.EmptyStruct{}, nil
}

// ExpLog get member exp logs
func (s *MemberServer) ExpLog(ctx context.Context, req *v1.MidReq) (*v1.UserLogsReply, error) {
	res, err := s.svr.ExpLog(ctx, req.Mid, req.RealIP)
	if err != nil {
		return nil, err
	}
	userLogs := make([]*v1.UserLogReply, 0, len(res))
	for _, v := range res {
		var userLog = &v1.UserLogReply{
			Mid:     v.Mid,
			Ip:      v.IP,
			Ts:      v.TS,
			LogId:   v.LogID,
			Content: v.Content,
		}
		userLogs = append(userLogs, userLog)
	}
	userLogsReply := &v1.UserLogsReply{
		UserLogs: userLogs,
	}
	return userLogsReply, nil
}

// ExpStat get exp status
func (s *MemberServer) ExpStat(ctx context.Context, req *v1.MidReq) (*v1.ExpStatReply, error) {
	res, err := s.svr.Stat(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	expStatReply := &v1.ExpStatReply{
		Login: res.Login,
		Watch: res.Watch,
		Coin:  res.Coin,
		Share: res.Share,
	}

	return expStatReply, nil
}
