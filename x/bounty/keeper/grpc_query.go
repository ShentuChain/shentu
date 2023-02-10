package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

var _ types.QueryServer = Keeper{}

// Hosts implements the Query/Hosts gRPC method
func (k Keeper) Hosts(c context.Context, req *types.QueryHostsRequest) (*types.QueryHostsResponse, error) {
	//TODO implement me
	panic("implement me")
}

// Host implements
func (k Keeper) Host(c context.Context, req *types.QueryHostRequest) (*types.QueryHostResponse, error) {
	//TODO implement me
	panic("implement me")
}

// Programs implements the Query/Programs gRPC method
func (k Keeper) Programs(c context.Context, req *types.QueryProgramsRequest) (*types.QueryProgramsResponse, error) {
	var programs types.Programs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	programStore := prefix.NewStore(store, types.ProgramsKey)

	pageRes, err := query.FilteredPaginate(programStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var p types.Program
		if err := k.cdc.Unmarshal(value, &p); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		if accumulate {
			programs = append(programs, p)
		}

		return true, nil

	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryProgramsResponse{
		Programs:   programs,
		Pagination: pageRes,
	}, nil
}

// Program returns program details based on ProgramId
func (k Keeper) Program(c context.Context, req *types.QueryProgramRequest) (*types.QueryProgramResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.ProgramId == 0 {
		return nil, status.Error(codes.InvalidArgument, "program-id can not be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	program, found := k.GetProgram(ctx, req.ProgramId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "program %d doesn't exist", req.ProgramId)
	}

	return &types.QueryProgramResponse{Program: program}, nil
}

func (k Keeper) Findings(c context.Context, req *types.QueryFindingsRequest) (*types.QueryFindingsResponse, error) {
	var findings types.Findings
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	programStore := prefix.NewStore(store, types.FindingKey)

	pageRes, err := query.FilteredPaginate(programStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var finding types.Finding
		if err := k.cdc.Unmarshal(value, &finding); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		matchProgramID, matchSubmitter := true, true
		// match program-id
		if req.ProgramId != 0 {
			matchProgramID = req.ProgramId == finding.ProgramId
		}
		// match submitter address
		if len(req.SubmitterAddress) > 0 {
			matchSubmitter = req.SubmitterAddress == finding.SubmitterAddress

		}

		if matchProgramID && matchSubmitter {
			if accumulate {
				findings = append(findings, finding)
			}
		}

		return true, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFindingsResponse{
		Findings:   findings,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) Finding(c context.Context, req *types.QueryFindingRequest) (*types.QueryFindingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.FindingId == 0 {
		return nil, status.Error(codes.InvalidArgument, "finding-id can not be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	finding, found := k.GetFinding(ctx, req.FindingId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "finding %d doesn't exist", req.FindingId)
	}

	return &types.QueryFindingResponse{Finding: finding}, nil
}

func (k Keeper) NextProgramID(c context.Context, req *types.QueryNextProgramIDRequest) (*types.QueryNextProgramIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	npid := k.GetNextProgramID(ctx)
	return &types.QueryNextProgramIDResponse{NextProgramId: npid}, nil
}
