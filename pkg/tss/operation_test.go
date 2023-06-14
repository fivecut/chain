package tss_test

import (
	"encoding/hex"

	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/pkg/tss/testutil"
)

func (suite *TSSTestSuite) TestComputeKeySym() {
	suite.RunOnPairMembers(
		suite.testCases,
		func(tc testutil.TestCase, memberI testutil.Member, memberJ testutil.Member) {
			keySym, err := tss.ComputeKeySym(
				memberI.OneTimePrivKey,
				memberJ.OneTimePubKey(),
			)

			suite.Require().NoError(err)
			suite.Require().Equal(memberI.KeySyms[testutil.GetSlot(memberI.ID, memberJ.ID)], keySym)
		},
	)
}

func (suite *TSSTestSuite) TestComputeNonceSym() {
	nonce := testutil.HexDecode("1111111111111111111111111111111111111111111111111111111111111111")
	pubKey := testutil.HexDecode("03c820245f18671206e752122953397786af3444f3a8da8098e594a8f612d94059")

	nonceSym, err := tss.ComputeNonceSym(nonce, pubKey)
	suite.Require().NoError(err)
	suite.Require().
		Equal("0360bf3f69810cc3472702c1f76ec76cdbefd85b1537db870e91b382a2e6e2bf6c", hex.EncodeToString(nonceSym))
}

func (suite *TSSTestSuite) TestSumScalars() {
	tests := []struct {
		name     string
		scalars  tss.Scalars
		expTotal tss.Scalar
		expError error
	}{
		{
			"zero element",
			tss.Scalars{},
			testutil.HexDecode("0000000000000000000000000000000000000000000000000000000000000000"),
			nil,
		},
		{
			"one element",
			tss.Scalars{testutil.HexDecode("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")},
			testutil.HexDecode("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"),
			nil,
		},
		{
			"three element",
			tss.Scalars{
				testutil.HexDecode("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"),
				testutil.HexDecode("c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"),
			},
			testutil.HexDecode("3fc2e6133bca391985e5a304644787e0a464ae400b74c54545cc2c87a332753c"),
			nil,
		},
		{
			"big values",
			tss.Scalars{
				testutil.HexDecode("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
				testutil.HexDecode("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
			},
			testutil.HexDecode("000000000000000000000000000000028aa24632a16ebf88805b42e45f9375de"),
			nil,
		},
		{
			"length is too short",
			tss.Scalars{
				testutil.HexDecode("fffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
				testutil.HexDecode("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
			},
			nil,
			tss.ErrInvalidLength,
		},
		{
			"length is too short",
			tss.Scalars{
				testutil.HexDecode("02fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
				testutil.HexDecode("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
			},
			nil,
			tss.ErrInvalidLength,
		},
	}
	for _, t := range tests {
		suite.Run(t.name, func() {
			total, err := tss.SumScalars(t.scalars...)
			suite.Require().ErrorIs(err, t.expError)
			suite.Require().Equal(t.expTotal, total)
		})
	}
}

func (suite *TSSTestSuite) TestSolveScalarPolynomial() {
	tests := []struct {
		name         string
		coefficients tss.Scalars
		x            tss.Scalar
		expTotal     tss.Scalar
		expError     error
	}{
		{
			"case 1",
			tss.Scalars{
				testutil.HexDecode("b07024fb8035d29ad2bdcf422bb460e83ac816a9319b4566aca8031be6502169"),
				testutil.HexDecode("2611e629e7043dbd32682e91c73292b087bb9f0747cd4bdd94e92093b6716504"),
			},
			testutil.HexDecode("0000000000000000000000000000000000000000000000000000000000000002"),
			testutil.HexDecode("fc93f14f4e3e4e15378e2c65ba1986494a3f54b7c135dd21d67a44435332eb71"),
			nil,
		},
		{
			"case 2",
			tss.Scalars{
				testutil.HexDecode("98ddb2a9f4a9fc49d06ea78a92a6059b8c6fe05a4ed4b4a6d0a002549829b9b5"),
				testutil.HexDecode("42e8ead39b0d57a943cf5d7fba99da80a96eac0599bebfea0cfc5a775a6bae09"),
			},
			testutil.HexDecode("0000000000000000000000000000000000000000000000000000000000000001"),
			testutil.HexDecode("dbc69d7d8fb753f3143e050a4d3fe01c35de8c5fe8937490dd9c5ccbf29567be"),
			nil,
		},
	}
	for _, t := range tests {
		suite.Run(t.name, func() {
			result, err := tss.SolveScalarPolynomial(t.coefficients, t.x)
			suite.Require().ErrorIs(err, t.expError)
			suite.Require().Equal(t.expTotal, result)
		})
	}
}

func (suite *TSSTestSuite) TestSumPoints() {
	tests := []struct {
		name     string
		points   tss.Points
		expTotal tss.Point
		expError error
	}{
		{
			"zero element",
			tss.Points{},
			testutil.HexDecode("020000000000000000000000000000000000000000000000000000000000000000"),
			nil,
		},
		{
			"one element",
			tss.Points{testutil.HexDecode("0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")},
			testutil.HexDecode("0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"),
			nil,
		},
		{
			"three element",
			tss.Points{
				testutil.HexDecode("0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"),
				testutil.HexDecode("02c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"),
			},
			testutil.HexDecode("02f9308a019258c31049344f85f89d5229b531c845836f99b08601f113bce036f9"),
			nil,
		},
		{
			"value is too big",
			tss.Points{
				testutil.HexDecode("02fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
				testutil.HexDecode("02fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc30"),
			},
			nil,
			tss.ErrParseError,
		},
	}
	for _, t := range tests {
		suite.Run(t.name, func() {
			total, err := tss.SumPoints(t.points...)
			suite.Require().ErrorIs(err, t.expError)
			suite.Require().Equal(t.expTotal, total)
		})
	}
}

func (suite *TSSTestSuite) TestSolvePointPolynomial() {
	tests := []struct {
		name         string
		coefficients tss.Points
		x            tss.Scalar
		expTotal     tss.Point
		expError     error
	}{
		{
			"case 1",
			tss.Points{
				testutil.HexDecode("039cff182d3b5653c215207c5b141983d6e784e51cc8088b7edfef6cba504573e3"),
				testutil.HexDecode("035b8a99ebc56c07b88404407046e6f0d5e5318a87b888ea25d6d12d8175b2d70c"),
			},
			testutil.HexDecode("0000000000000000000000000000000000000000000000000000000000000002"),
			testutil.HexDecode("021e230158a714d176058051be2389b5b60c1700b56cc7a0f9387911aa92f2963a"),
			nil,
		},
		{
			"case 2",
			tss.Points{
				testutil.HexDecode("02786741d28ca0a66b628d6401d975da448fc08c15a1228eb7b65203c6bac5cedb"),
				testutil.HexDecode("023d61b24c8785efe8c7459dc706d95b197c0acb31697feb49fec2d3446dc36de4"),
			},
			testutil.HexDecode("0000000000000000000000000000000000000000000000000000000000000001"),
			testutil.HexDecode("02e9cceb096f012fc7170b85c961036e9baf0c3e9b58cb94f80b4376bb38876806"),
			nil,
		},
	}
	for _, t := range tests {
		suite.Run(t.name, func() {
			result, err := tss.SolvePointPolynomial(t.coefficients, t.x)
			suite.Require().ErrorIs(err, t.expError)
			suite.Require().Equal(t.expTotal, result)
		})
	}
}
