package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// module name
	ModuleName = "bandtss"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName

	// RouterKey is the message route for the tss module
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the tss module
	QuerierRoute = ModuleName
)

var (

	// GlobalStoreKeyPrefix is the prefix for global primitive state variables.
	GlobalStoreKeyPrefix = []byte{0x00}

	// ParamsKeyPrefix is a prefix for keys that store TSSMember's parameters
	ParamsKeyPrefix = []byte{0x01}

	// StatusStoreKeyPrefix is the prefix for status store.
	StatusStoreKeyPrefix = []byte{0x02}
)

func StatusStoreKey(address sdk.AccAddress) []byte {
	return append(StatusStoreKeyPrefix, address...)
}
