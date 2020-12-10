package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maticnetwork/heimdall/helper"
	hmTypes "github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/heimdall/x/staking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
	contractCaller helper.IContractCaller
}

// NewQueryServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper, contractCaller helper.IContractCaller) types.QueryServer {
	return &Querier{Keeper: keeper, contractCaller: contractCaller}
}

var _ types.QueryServer = Querier{}

// Validator queries validator info for given validator addr
func (k Querier) Validator(c context.Context, req *types.QueryValidatorRequest) (*types.QueryValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	validatorID := hmTypes.ValidatorID(req.ValidatorId)
	if req.ValidatorId == 0 {
		return nil, status.Error(codes.InvalidArgument, "validator ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	validator, found := k.GetValidatorFromValID(ctx, validatorID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "validator %s not found", req.ValidatorId)
	}

	return &types.QueryValidatorResponse{Validator: &validator}, nil
}
