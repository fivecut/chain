package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bandprotocol/chain/v2/x/restake/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	*Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) Keys(
	c context.Context,
	req *types.QueryKeysRequest,
) (*types.QueryKeysResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	keyStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyStoreKeyPrefix)

	filteredKeys, pageRes, err := query.GenericFilteredPaginate(
		k.cdc,
		keyStore,
		req.Pagination,
		func(key []byte, v *types.Key) (*types.Key, error) {
			return v, nil
		}, func() *types.Key {
			return &types.Key{}
		})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryKeysResponse{Keys: filteredKeys, Pagination: pageRes}, nil
}

func (k Querier) Rewards(
	c context.Context,
	req *types.QueryRewardsRequest,
) (*types.QueryRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	stakes := k.GetStakes(ctx, address)
	for _, stake := range stakes {
		k.ProcessStake(ctx, stake)
	}

	return &types.QueryRewardsResponse{
		Rewards: k.GetRewards(ctx, address),
	}, nil
}

func (k Querier) Locks(
	c context.Context,
	req *types.QueryLocksRequest,
) (*types.QueryLocksResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	var locks []types.Lock

	stakes := k.GetActiveStakes(ctx, address)
	for _, stake := range stakes {
		locks = append(locks, types.Lock{
			Key:    stake.Key,
			Amount: stake.Amount,
		})
	}

	return &types.QueryLocksResponse{Locks: locks}, nil
}
