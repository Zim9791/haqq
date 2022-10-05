package types

// constants
const (
	// module name
	ModuleName = "coinomics"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for message routing
	RouterKey = ModuleName
)

// prefix bytes for the inflation persistent store
const (
	prefixInflation = iota + 1
	prefixEra
	prefixEraStartedAtBlock
	prefixEraTargetMint
	prefixEraTargetSupply
	prefixTargetTotalSupply
)

// KVStore key prefixes
var (
	KeyPrefixInflation         = []byte{prefixInflation}
	KeyPrefixEra               = []byte{prefixEra}
	KeyPrefixEraStartedAtBlock = []byte{prefixEraStartedAtBlock}
	KetPrefixEraTargetMint     = []byte{prefixEraTargetMint}
	KeyPrefixEraTargetSupply   = []byte{prefixEraTargetSupply}
	KeyPrefixTargetTotalSupply = []byte{prefixTargetTotalSupply}
)
