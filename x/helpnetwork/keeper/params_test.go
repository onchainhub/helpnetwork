package keeper_test

import (
	"testing"

	testkeeper "github.com/onchainengineer/helpnetwork/testutil/keeper"
	"github.com/onchainengineer/helpnetwork/x/helpnetwork/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HelpnetworkKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
