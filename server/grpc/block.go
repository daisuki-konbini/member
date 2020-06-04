package grpc

import (
	"context"

	"go-common/app/service/main/member/api/grpc/v1"
)

// BlockInfo 查询封禁信息
func (s *MemberServer) BlockInfo(ctx context.Context, req *v1.MemberMidReq) (*v1.BlockInfoReply, error) {
	res, err := s.blockSvr.Infos(ctx, []int64{req.Mid})
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return v1.FromBlockInfo(s.blockSvr.DefaultUser(req.Mid)), nil
	}
	blockInfoReply := v1.FromBlockInfo(res[0])
	return blockInfoReply, nil
}

// BlockBatchInfo 批量查询封禁信息
func (s *MemberServer) BlockBatchInfo(ctx context.Context, req *v1.MemberMidsReq) (*v1.BlockBatchInfoReply, error) {
	res, err := s.blockSvr.Infos(ctx, req.Mids)
	if err != nil {
		return nil, err
	}
	blockInfos := make([]*v1.BlockInfoReply, 0, len(res))
	for i := range res {
		blockInfos[i] = v1.FromBlockInfo(res[i])
	}
	blockInfosReply := &v1.BlockBatchInfoReply{
		BlockInfos: blockInfos,
	}
	return blockInfosReply, nil
}
