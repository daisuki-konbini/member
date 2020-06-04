package grpc

import (
	"context"
	"go-common/app/service/main/member/api/grpc/v1"
)

// Moral Get member moral info
func (s *MemberServer) Moral(ctx context.Context, req *v1.MemberMidReq) (*v1.MoralReply, error) {
	res, err := s.svr.Moral(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	moralReply := &v1.MoralReply{
		Mid:             res.Mid,
		Moral:           res.Moral,
		Added:           res.Added,
		Deducted:        res.Deducted,
		LastRecoverDate: res.LastRecoverDate,
	}

	return moralReply, nil
}

// MoralLog Get member moral logs
func (s *MemberServer) MoralLog(ctx context.Context, req *v1.MemberMidReq) (*v1.UserLogsReply, error) {
	res, err := s.svr.MoralLog(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	userLogs := make([]*v1.UserLogReply, 0, len(res))
	for _, v := range res {
		userLog := &v1.UserLogReply{
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

// AddMoral Add member's moral value
func (s *MemberServer) AddMoral(ctx context.Context, req *v1.UpdateMoralReq) (*v1.EmptyStruct, error) {
	err := s.svr.UpdateMoral(ctx, v1.ToArgUpdateMoral(req))
	if err != nil {
		return nil, err
	}
	return &v1.EmptyStruct{}, nil
}

// BatchAddMoral Batch add member's moral value
func (s *MemberServer) BatchAddMoral(ctx context.Context, req *v1.UpdateMoralsReq) (*v1.UpdateMoralsReply, error) {
	res, err := s.svr.UpdateMorals(ctx, v1.ToArgUpdateMorals(req))
	if err != nil {
		return nil, err
	}
	updateMoralsReply := &v1.UpdateMoralsReply{
		AfterMorals: res,
	}
	return updateMoralsReply, nil
}
