package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var DefaultMintDenom = "aISLM"

// Parameter store keys
var (
	ParamStoreKeyMintDenom        = []byte("ParamStoreKeyMintDenom")
	ParamStoreKeyBlockPerEra      = []byte("ParamStoreKeyBlockPerEra")
	ParamStoreKeyEnableCoinomics  = []byte("ParamStoreKeyEnableCoinomics")
	ParamStoreKeyMintDistribution = []byte("ParamStoreKeyMintDistribution")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	mintDenom string,
	blockPerEra uint64,
	enableCoinomics bool,
	mintDistribution MintDistribution,
) Params {
	return Params{
		MintDenom:        mintDenom,
		BlocksPerEra:     blockPerEra,
		EnableCoinomics:  enableCoinomics,
		MintDistribution: mintDistribution,
	}
}

// 20346674890910137847030794376

// 346674890 91 0137 8470 3079 4376
// 6066810 59 0927 4123 2303 8901.580000000000000000

// 5259600 * 2, // 2 years

func DefaultParams() Params {
	return Params{
		MintDenom:       DefaultMintDenom,
		BlocksPerEra:    100,
		EnableCoinomics: true,
		MintDistribution: MintDistribution{
			EvergreenDao:     sdk.NewDecWithPrec(1000, 2),
			BlockProposerMin: sdk.NewDecWithPrec(100, 2),
			BlockProposerMax: sdk.NewDecWithPrec(500, 2),
		},
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(ParamStoreKeyBlockPerEra, &p.BlocksPerEra, validateBlockPerEra),
		paramtypes.NewParamSetPair(ParamStoreKeyEnableCoinomics, &p.EnableCoinomics, validateBool),
		paramtypes.NewParamSetPair(ParamStoreKeyMintDistribution, &p.MintDistribution, validateMintDistribution),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateBlockPerEra(i interface{}) error {
	v, ok := i.(uint64)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return errors.New("block per era must not be zero")
	}

	return nil
}

func validateMintDistribution(i interface{}) error {
	fmt.Printf("validateMintDistribution: %T", i)

	v, ok := i.(MintDistribution)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.EvergreenDao.IsNegative() {
		return fmt.Errorf("MintDistribution: EvergreenDao value cannot be negative")
	}

	if v.BlockProposerMin.IsNegative() {
		return fmt.Errorf("MintDistribution: BlockProposerMin value cannot be negative")
	}

	if v.BlockProposerMax.IsNegative() {
		return fmt.Errorf("MintDistribution: BlockProposerMax value cannot be negative")
	}

	return nil
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateBlockPerEra(p.BlocksPerEra); err != nil {
		return err
	}

	if err := validateMintDistribution(p.MintDistribution); err != nil {
		return err
	}

	return validateBool(p.EnableCoinomics)
}
