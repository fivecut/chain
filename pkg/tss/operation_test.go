package tss_test

import (
	"encoding/hex"

	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func (suite *TSSTestSuite) TestComputeKeySym() {
	keySym, err := tss.ComputeKeySym(suite.kpI.PrivateKey, suite.kpJ.PublicKey)
	suite.Require().NoError(err)
	suite.Require().
		Equal("03bc213e4251592d29c070e4c31b980d150e755ec848afa4c06730ec1dcd09c482", hex.EncodeToString(keySym))
}

func (suite *TSSTestSuite) TestComputeNonceSymOddCase() {
	nonceSym, err := tss.ComputeNonceSym(tss.Scalar(suite.kpI.PrivateKey), suite.kpJ.PublicKey)
	suite.Require().NoError(err)
	suite.Require().
		Equal("03bc213e4251592d29c070e4c31b980d150e755ec848afa4c06730ec1dcd09c482", hex.EncodeToString(nonceSym))
}

func (suite *TSSTestSuite) TestComputeNonceSymEvenCase() {
	nonceSym, err := tss.ComputeNonceSym(tss.ParseScalar(new(secp256k1.ModNScalar).SetInt(1)), suite.kpJ.PublicKey)
	suite.Require().NoError(err)
	suite.Require().
		Equal("03f70e80bac0b32b2599fa54d83b5471e90fac27bb09528f0337b49d464d64426f", hex.EncodeToString(nonceSym))
}

func (suite *TSSTestSuite) TestSumPoints() {
	// Prepare
	var p1, p2, expectedPoint secp256k1.JacobianPoint

	s1 := new(secp256k1.ModNScalar).SetInt(1)
	secp256k1.ScalarBaseMultNonConst(s1, &p1)

	s2 := new(secp256k1.ModNScalar).SetInt(2)
	secp256k1.ScalarBaseMultNonConst(s2, &p2)

	secp256k1.ScalarBaseMultNonConst(s1.Add(s2), &expectedPoint)

	// Try sum with function
	total, err := tss.SumPoints(tss.ParsePoint(&p1), tss.ParsePoint(&p2))
	suite.Require().NoError(err)
	suite.Require().Equal(tss.ParsePoint(&expectedPoint), total)
}

func (suite *TSSTestSuite) TestSumScalars() {
	total, err := tss.SumScalars(
		tss.ParseScalar(new(secp256k1.ModNScalar).SetInt(1)),
		tss.ParseScalar(new(secp256k1.ModNScalar).SetInt(2)),
		tss.ParseScalar(new(secp256k1.ModNScalar).SetInt(3)),
	)
	suite.Require().NoError(err)
	suite.Require().Equal(tss.ParseScalar(new(secp256k1.ModNScalar).SetInt(6)), total)
}
