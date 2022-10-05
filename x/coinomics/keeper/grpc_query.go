package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/haqq-network/haqq/x/coinomics/types"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Era(
	c context.Context,
	_ *types.QueryEraRequest,
) (*types.QueryEraResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	era := k.GetEra(ctx)

	return &types.QueryEraResponse{Era: era}, nil
}

func (k Keeper) InflationRate(
	c context.Context,
	_ *types.QueryInflationRateRequest,
) (*types.QueryInflationRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	inflation := k.GetInflation(ctx)

	return &types.QueryInflationRateResponse{InflationRate: inflation}, nil
}

func (k Keeper) TotalTargetSupply(
	c context.Context,
	_ *types.QueryTotalTargetSupplyRequest,
) (*types.QueryTotalTargetSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	totalTarget := k.GetTargetTotalSupply(ctx)

	return &types.QueryTotalTargetSupplyResponse{TotalTargetSupply: totalTarget}, nil
}

func (k Keeper) EraTargetSupply(
	c context.Context,
	_ *types.QueryEraTargetSupplyRequest,
) (*types.QueryEraTargetSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	eraTargetSupply := k.GetEraTargetSupply(ctx)

	return &types.QueryEraTargetSupplyResponse{EraTargetSupply: eraTargetSupply}, nil
}

// Params returns params of the mint module.
func (k Keeper) Params(
	c context.Context,
	_ *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
