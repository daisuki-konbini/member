package grpc

import (
	"context"
	"go-common/app/service/main/member/api/grpc/v1"
	"go-common/app/service/main/member/model"
)

// AddUserMonitor add user monitor
func (s *MemberServer) AddUserMonitor(ctx context.Context, req *v1.AddUserMonitorReq) (*v1.EmptyStruct, error) {
	argAddUserMonitor := &model.ArgAddUserMonitor{
		Mid:      req.Mid,
		Operator: req.Operator,
		Remark:   req.Remark,
	}
	err := s.svr.AddUserMonitor(ctx, argAddUserMonitor)
	if err != nil {
		return nil, err
	}

	emptyStruct := &v1.EmptyStruct{}
	return emptyStruct, nil
}

// IsInMonitor check whether the member is in monitored status
func (s *MemberServer) IsInMonitor(ctx context.Context, req *v1.MidReq) (*v1.IsInMonitorReply, error) {
	res, err := s.svr.IsInMonitor(ctx, &model.ArgMid{
		Mid:    req.Mid,
		RealIP: req.RealIP,
	})
	if err != nil {
		return nil, err
	}

	isInMonitorReply := &v1.IsInMonitorReply{
		IsInMonitor: res,
	}
	return isInMonitorReply, nil
}
