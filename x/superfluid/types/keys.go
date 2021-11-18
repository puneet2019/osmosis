package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// ModuleName defines the module name
	ModuleName = "superfluid"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// KeyPrefixSuperfluidAsset defines prefix key for superfluid asset
	KeyPrefixSuperfluidAsset = []byte{0x01}

	// KeyPrefixSuperfluidAssetInfo defines prefix key for superfluid asset info
	KeyPrefixSuperfluidAssetInfo = []byte{0x02}

	// KeyPrefixTokenPriceTwap defines prefix key per epoch
	KeyPrefixTokenPriceTwap = []byte{0x03}

	// KeyPrefixIntermediaryAccount defines prefix to set intermediary account struct to its address
	KeyPrefixIntermediaryAccount = []byte{0x04}

	// KeyPrefixLockIntermediaryAccAddr defines prefix to connect lockId and intermediary account address
	KeyPrefixLockIntermediaryAccAddr = []byte{0x05}
)

func TokenPriceTwapEpochPrefix(epoch int64) []byte {
	return append(KeyPrefixTokenPriceTwap, sdk.Uint64ToBigEndian(uint64(epoch))...)
}
