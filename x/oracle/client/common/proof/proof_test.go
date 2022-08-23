package proof

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/bandprotocol/chain/v2/x/oracle/types"
)

func hexToBytes(hexstr string) []byte {
	b, err := hex.DecodeString(hexstr)
	if err != nil {
		panic(err)
	}
	return b
}

func leafHash(item []byte) []byte {
	// leaf prefix is 0
	return tmhash.Sum(append([]byte{0}, item...))
}

func innerHash(left, right []byte) []byte {
	// branch prefix is 1
	return tmhash.Sum(append([]byte{1}, append(left, right...)...))
}

func TestEncodeRelay(t *testing.T) {
	block := BlockRelayProof{
		MultiStoreProof: MultiStoreProof{
			OracleIAVLStateHash:               hexToBytes("5248463E932D16F7D092E268C0DED87B23D3B0E71856F1C6AE2AA91F6C713320"),
			ParamsStoreMerkleHash:             hexToBytes("7FD5F5C7C2920C187618542901CDC5717BE8204F24BE856E80902A1BB04737E4"),
			SlashingToStakingStoresMerkleHash: hexToBytes("F981716562A49DE06E3DCAFBFB6388C294BAA4FA9D45777E25740A92F81CF65E"),
			GovToMintStoresMerkleHash:         hexToBytes("E8E27CBB44BB654F64EEEF4667868AD48667CEB28E3DB5C4DF7A4B4B87F0C04B"),
			AuthToFeegrantStoresMerkleHash:    hexToBytes("FC96CFFD30E5B8979EA66F9D0DA1CBAB16F69669E8B2A1FB2E1BEB457C9726E8"),
			TransferToUpgradeStoresMerkleHash: hexToBytes("C9C8849ED125CC7681329C4D27B83B1FC8ACF7A865C9D1D1DF575CCA56F48DBE"),
		},
		BlockHeaderMerkleParts: BlockHeaderMerkleParts{
			VersionAndChainIdHash:             hexToBytes("3F02642D9E70D5C1C493A4F732BFE9C9B95A4A42651703B816EDCFC8FADA5312"),
			Height:                            25000,
			TimeSecond:                        1629849931,
			TimeNanoSecond:                    290650376,
			LastBlockIdAndOther:               hexToBytes("9B4825C99C3E739E1DC171FFB0E2BF34E99EEE41B34E407E40CF594834427B09"),
			NextValidatorHashAndConsensusHash: hexToBytes("BF23413F237906B07202B3355E7311651ACE6BD2A34FD6FC3BD98EFE4FB78755"),
			LastResultsHash:                   hexToBytes("9FB9C7533CAF1D218DA3AF6D277F6B101C42E3C3B75D784242DA663604DD53C2"),
			EvidenceAndProposerHash:           hexToBytes("7D11A74E40884411901BD7A70631734990B1FDBF5DE9E4C92C63B7650A6A6659"),
		},
		Signatures: []TMSignature{
			{
				R:                hexToBytes("84B8585B71240FEE0E674952B79ED25D793F1B31B42DD37B80F75B98510B5754"),
				S:                hexToBytes("1EC44DD7C5389474DF8E5C25CC6ED8B573CCA2E009AA824EE825BDC693935927"),
				V:                28,
				SignedDataPrefix: hexToBytes("6D080211A86100000000000022480A20"),
				SignedDataSuffix: hexToBytes("1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD9296890610EAE9963D320962616E64636861696E"),
			},
			{
				R:                hexToBytes("394365193F819CF539381366D31B6C5849AAA31AE8BA6F95C62C5C80656BFB5C"),
				S:                hexToBytes("6A07E4A3C0ABCEAE5F854D492DF699438FB84762F152F739DDEAC48DDCFCB5CC"),
				V:                28,
				SignedDataPrefix: hexToBytes("6D080211A86100000000000022480A20"),
				SignedDataSuffix: hexToBytes("1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD9296890610EA928633320962616E64636861696E"),
			},
			{
				R:                hexToBytes("5D7B4BE7B21B00D08AD7DBE48CF2761CECCB599E64AAB10B2901A0DD58F00325"),
				S:                hexToBytes("7160EF689A533C1E983707507FC8466DAEA1D0DC7A889E3A27D1BB1D09CEC030"),
				V:                28,
				SignedDataPrefix: hexToBytes("6D080211A86100000000000022480A20"),
				SignedDataSuffix: hexToBytes("1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD929689061086FAB239320962616E64636861696E"),
			},
			{
				R:                hexToBytes("5654A44FB89330C34CF2D862F940763194A145B72ED3BB0ADD5759E1E68FD145"),
				S:                hexToBytes("2AC795D02A9C574CF12343FDFC67FDCED8A24F88EC8138C7F8230F6EB442B726"),
				V:                28,
				SignedDataPrefix: hexToBytes("6D080211A86100000000000022480A20"),
				SignedDataSuffix: hexToBytes("1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD9296890610F0E1F733320962616E64636861696E"),
			},
		},
	}
	result, err := block.encodeToEthData()
	require.Nil(t, err)
	expect := hexToBytes("5248463E932D16F7D092E268C0DED87B23D3B0E71856F1C6AE2AA91F6C7133207FD5F5C7C2920C187618542901CDC5717BE8204F24BE856E80902A1BB04737E4F981716562A49DE06E3DCAFBFB6388C294BAA4FA9D45777E25740A92F81CF65EE8E27CBB44BB654F64EEEF4667868AD48667CEB28E3DB5C4DF7A4B4B87F0C04BFC96CFFD30E5B8979EA66F9D0DA1CBAB16F69669E8B2A1FB2E1BEB457C9726E8C9C8849ED125CC7681329C4D27B83B1FC8ACF7A865C9D1D1DF575CCA56F48DBE3F02642D9E70D5C1C493A4F732BFE9C9B95A4A42651703B816EDCFC8FADA531200000000000000000000000000000000000000000000000000000000000061A8000000000000000000000000000000000000000000000000000000006125894B000000000000000000000000000000000000000000000000000000001152F9089B4825C99C3E739E1DC171FFB0E2BF34E99EEE41B34E407E40CF594834427B09BF23413F237906B07202B3355E7311651ACE6BD2A34FD6FC3BD98EFE4FB787559FB9C7533CAF1D218DA3AF6D277F6B101C42E3C3B75D784242DA663604DD53C27D11A74E40884411901BD7A70631734990B1FDBF5DE9E4C92C63B7650A6A665900000000000000000000000000000000000000000000000000000000000001E00000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000001C00000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000044084B8585B71240FEE0E674952B79ED25D793F1B31B42DD37B80F75B98510B57541EC44DD7C5389474DF8E5C25CC6ED8B573CCA2E009AA824EE825BDC693935927000000000000000000000000000000000000000000000000000000000000001C00000000000000000000000000000000000000000000000000000000000000A000000000000000000000000000000000000000000000000000000000000000E000000000000000000000000000000000000000000000000000000000000000106D080211A86100000000000022480A2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003E1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD9296890610EAE9963D320962616E64636861696E0000394365193F819CF539381366D31B6C5849AAA31AE8BA6F95C62C5C80656BFB5C6A07E4A3C0ABCEAE5F854D492DF699438FB84762F152F739DDEAC48DDCFCB5CC000000000000000000000000000000000000000000000000000000000000001C00000000000000000000000000000000000000000000000000000000000000A000000000000000000000000000000000000000000000000000000000000000E000000000000000000000000000000000000000000000000000000000000000106D080211A86100000000000022480A2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003E1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD9296890610EA928633320962616E64636861696E00005D7B4BE7B21B00D08AD7DBE48CF2761CECCB599E64AAB10B2901A0DD58F003257160EF689A533C1E983707507FC8466DAEA1D0DC7A889E3A27D1BB1D09CEC030000000000000000000000000000000000000000000000000000000000000001C00000000000000000000000000000000000000000000000000000000000000A000000000000000000000000000000000000000000000000000000000000000E000000000000000000000000000000000000000000000000000000000000000106D080211A86100000000000022480A2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003E1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD929689061086FAB239320962616E64636861696E00005654A44FB89330C34CF2D862F940763194A145B72ED3BB0ADD5759E1E68FD1452AC795D02A9C574CF12343FDFC67FDCED8A24F88EC8138C7F8230F6EB442B726000000000000000000000000000000000000000000000000000000000000001C00000000000000000000000000000000000000000000000000000000000000A000000000000000000000000000000000000000000000000000000000000000E000000000000000000000000000000000000000000000000000000000000000106D080211A86100000000000022480A2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003E1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A63012A0B08CD9296890610F0E1F733320962616E64636861696E0000")
	require.Equal(t, expect, result)
}

func TestEncodeVerify(t *testing.T) {
	data := OracleDataProof{
		Version: 217,
		Result: types.Result{
			ClientID:       "",
			OracleScriptID: 1,
			Calldata:       hexToBytes("000000010000000342544300000000000186a0"),
			AskCount:       1,
			MinCount:       1,
			RequestID:      1,
			AnsCount:       1,
			RequestTime:    1629803667,
			ResolveTime:    1629803671,
			ResolveStatus:  1,
			Result:         hexToBytes("000000010000000124ec078c"),
		},
		MerklePaths: []IAVLMerklePath{
			{
				IsDataOnRight:  true,
				SubtreeHeight:  1,
				SubtreeSize:    2,
				SubtreeVersion: 217,
				SiblingHash:    hexToBytes("EB739BB22F48B7F3053A90BA2BA4FE07FAB262CADF8664489565C50FF505B8BD"),
			},
			{
				false,
				2,
				4,
				217,
				hexToBytes("1847107507D5E7B4CD9941EB6FFE1694264AF34C685C19DC478BEADDA265A578"),
			},
			{
				true,
				3,
				6,
				217,
				hexToBytes("E80AAE581EC004239854C4D90D8148E85F1F90D0704A73668FD2DA44DC0CEA53"),
			},
			{
				false,
				5,
				16,
				24998,
				hexToBytes("3174174561E430E12C1EAD3DC5B08BAE7D4315F417D1202E7277879C1433E658"),
			},
		},
	}

	result, err := data.encodeToEthData(25000)
	require.Nil(t, err)
	expect := hexToBytes("00000000000000000000000000000000000000000000000000000000000061a8000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000d900000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000006124d493000000000000000000000000000000000000000000000000000000006124d497000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013000000010000000342544300000000000186a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000010000000124ec078c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000d9eb739bb22f48b7f3053a90ba2ba4fe07fab262cadf8664489565c50ff505b8bd00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000d91847107507d5e7b4cd9941eb6ffe1694264af34c685c19dc478beadda265a57800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000d9e80aae581ec004239854c4d90d8148e85f1f90d0704a73668fd2da44dc0cea5300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000061a63174174561e430e12c1ead3dc5b08bae7d4315f417d1202e7277879c1433e658")
	require.Equal(t, expect, result)
}

func TestEncodeVerifyCount(t *testing.T) {
	data := RequestsCountProof{
		Version: 215,
		Count:   1,
		MerklePaths: []IAVLMerklePath{
			{
				IsDataOnRight:  true,
				SubtreeHeight:  1,
				SubtreeSize:    2,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("7DB3EEB5BDBEBCDE9EC49489CED0BD8A1ABAB7E653C720F18DF8149572699F1F"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  2,
				SubtreeSize:    4,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("24427D84B35482BBDD44DAF9C13A2C573C8B9DD8FBD4E1BC3B0F9201D707EC6B"),
			},
			{
				IsDataOnRight:  false,
				SubtreeHeight:  4,
				SubtreeSize:    10,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("E7A2CA3EC55627E7D6B06189FF7EF46C13ECC3B1900D127C7D437A7BA3EE3FFA"),
			},
			{
				IsDataOnRight:  false,
				SubtreeHeight:  5,
				SubtreeSize:    16,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("CE859DBB78E1B401CF0C91864E6A94E20A647B180804D36982FA2297796F0908"),
			},
		},
	}

	result, err := data.encodeToEthData(25000)
	require.Nil(t, err)
	expect := hexToBytes("00000000000000000000000000000000000000000000000000000000000061a8000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000d70000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000061a67db3eeb5bdbebcde9ec49489ced0bd8a1abab7e653c720f18df8149572699f1f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000061a624427d84b35482bbdd44daf9c13a2c573c8b9dd8fbd4e1bc3b0f9201d707ec6b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000061a6e7a2ca3ec55627e7d6b06189ff7ef46c13ecc3b1900d127c7d437a7ba3ee3ffa00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000061a6ce859dbb78e1b401cf0c91864e6a94e20a647b180804d36982fa2297796f0908")
	require.Equal(t, expect, result)
}
