package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/onchainengineer/helpnetwork/testutil/keeper"
	"github.com/onchainengineer/helpnetwork/x/helpnetwork/keeper"
	"github.com/onchainengineer/helpnetwork/x/helpnetwork/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HelpnetworkKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
