package util

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	astraApp "github.com/AstraProtocol/astra/v2/app"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	"github.com/evmos/ethermint/encoding"
)

var (
	EncodingConfig params.EncodingConfig
	Cdc            codec.Codec
)

func init() {
	EncodingConfig, Cdc = initEncodingConfigAndCdc()
}

func initEncodingConfigAndCdc() (params.EncodingConfig, codec.Codec) {
	encodingConfig := encoding.MakeConfig(astraApp.ModuleBasics)

	encodingConfig.InterfaceRegistry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&stakingtypes.MsgCreateValidator{},
	)
	encodingConfig.InterfaceRegistry.RegisterImplementations(
		(*cryptotypes.PubKey)(nil),
		&ethsecp256k1.PubKey{},
		&ed25519.PubKey{},
	)

	cdc := encodingConfig.Marshaler

	return encodingConfig, cdc
}
