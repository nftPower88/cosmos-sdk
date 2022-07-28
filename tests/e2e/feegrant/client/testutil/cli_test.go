//go:build e2e
// +build e2e

package testutil

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	clienttestutil "github.com/cosmos/cosmos-sdk/x/feegrant/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/feegrant/testutil"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg, err := network.DefaultConfigWithAppConfig(testutil.AppConfig)
	require.NoError(t, err)
	cfg.NumValidators = 3
	suite.Run(t, clienttestutil.NewIntegrationTestSuite(cfg))
}