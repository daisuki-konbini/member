package grpc

import (
	"context"
	"go-common/app/service/main/member/api/grpc/v1"
)

// RealnameStatus get the member realname status
func (s *MemberServer) RealnameStatus(ctx context.Context, req *v1.MemberMidReq) (*v1.RealnameStatusReply, error) {
	res, err := s.svr.RealnameStatus(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	var realnameStatusReply = &v1.RealnameStatusReply{
		RealnameStatus: int8(res),
	}

	return realnameStatusReply, nil
}

// RealnameApplyStatus get member realname apply status
func (s *MemberServer) RealnameApplyStatus(ctx context.Context, req *v1.MemberMidReq) (*v1.RealnameApplyInfoReply, error) {
	res, err := s.svr.RealnameApplyStatus(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	var realnameStatusReply = &v1.RealnameApplyInfoReply{
		Status: int8(res.Status),
		Remark: res.Remark,
	}

	return realnameStatusReply, nil
}

// RealnameTelCapture mobilePhone realname certification
func (s *MemberServer) RealnameTelCapture(ctx context.Context, req *v1.MemberMidReq) (*v1.EmptyStruct, error) {
	_, err := s.svr.RealnameTelCapture(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	return &v1.EmptyStruct{}, nil
}

// RealnameApply apply for realname certification
func (s *MemberServer) RealnameApply(ctx context.Context, req *v1.ArgRealnameApplyReq) (*v1.EmptyStruct, error) {
	err := s.svr.RealnameApply(ctx, req.Mid, int(req.CaptureCode), req.Realname, req.CardType, req.CardCode, req.Country, req.HandIMGToken, req.FrontIMGToken, req.BackIMGToken)
	if err != nil {
		return nil, err
	}

	return &v1.EmptyStruct{}, nil
}

// RealnameDetail detail about realname by mid
func (s *MemberServer) RealnameDetail(ctx context.Context, req *v1.MemberMidReq) (*v1.RealnameDetailReply, error) {
	res, err := s.svr.RealnameDetail(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	var realnameDetail = &v1.RealnameDetailReply{
		Realname: res.Realname,
		Card:     res.Card,
		CardType: int8(res.CardType),
		Status:   int8(res.Status),
		Gender:   res.Gender,
		HandImg:  res.HandIMG,
	}

	return realnameDetail, nil
}

// RealnameStrippedInfo is
func (s *MemberServer) RealnameStrippedInfo(ctx context.Context, req *v1.MemberMidReq) (*v1.RealnameStrippedInfoReply, error) {
	return s.svr.RealnameStrippedInfo(ctx, req.Mid)
}

// MidByRealnameCard is
func (s *MemberServer) MidByRealnameCard(ctx context.Context, req *v1.MidByRealnameCardsReq) (*v1.MidByRealnameCardReply, error) {
	return s.svr.MidByRealnameCard(ctx, req)
}
