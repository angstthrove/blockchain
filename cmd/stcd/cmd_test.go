package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"freemasonry.cc/blockchain/app"
	evmosd "freemasonry.cc/blockchain/cmd/stcd/cmd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := evmosd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"evmos-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "evmos_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, "stcd", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := evmosd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"mykey2",
		"--coin-type",
		"118",
		"--index",
		"0",
		"--hd-path",
		"m/44'/118'/0'/0",
	})

	err := svrcmd.Execute(rootCmd, "stcd", app.DefaultNodeHome)
	require.Error(t, err)
}
