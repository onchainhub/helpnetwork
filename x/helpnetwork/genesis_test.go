package helpnetwork_test

import (
	"testing"

	keepertest "github.com/onchainengineer/helpnetwork/testutil/keeper"
	"github.com/onchainengineer/helpnetwork/testutil/nullify"
	"github.com/onchainengineer/helpnetwork/x/helpnetwork"
	"github.com/onchainengineer/helpnetwork/x/helpnetwork/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelpnetworkKeeper(t)
	helpnetwork.InitGenesis(ctx, *k, genesisState)
	got := helpnetwork.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
