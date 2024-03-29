package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Peersyst/exrp/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveValidator(goCtx context.Context, msg *types.MsgRemoveValidator) (*types.MsgRemoveValidatorResponse, error) {
	if k.authority != msg.Authority {
		return nil, errors.Wrapf(gov.ErrInvalidSigner, "expected %s got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.ExecuteRemoveValidator(ctx, msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveValidatorResponse{}, nil
}
