package node

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/cometbft/cometbft/crypto/secp256k1"

	gogogrpc "github.com/cosmos/gogoproto/grpc"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterNodeService registers the node gRPC service on the provided gRPC router.
func RegisterNodeService(clientCtx client.Context, server gogogrpc.Server, cfg config.Config) {
	RegisterServiceServer(server, NewQueryServer(clientCtx, cfg))
}

// RegisterGRPCGatewayRoutes mounts the node gRPC service's GRPC-gateway routes
// on the given mux object.
func RegisterGRPCGatewayRoutes(clientConn gogogrpc.ClientConn, mux *runtime.ServeMux) {
	_ = RegisterServiceHandlerClient(context.Background(), mux, NewServiceClient(clientConn))
}

// to check queryServer implements ServiceServer
var _ ServiceServer = queryServer{}

// queryServer implements ServiceServer
type queryServer struct {
	clientCtx client.Context
	cfg       config.Config
}

// NewQueryServer returns new queryServer from provided client.Context
func NewQueryServer(clientCtx client.Context, cfg config.Config) ServiceServer {
	return queryServer{
		clientCtx: clientCtx,
		cfg:       cfg,
	}
}

// ChainID returns QueryChainIDResponse that has chain id from ctx
func (s queryServer) ChainID(ctx context.Context, _ *ChainIDRequest) (*ChainIDResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	return &ChainIDResponse{
		ChainID: sdkCtx.ChainID(),
	}, nil
}

// EVMValidators returns top 100 validators' address and voting power order by voting power
func (s queryServer) EVMValidators(
	ctx context.Context,
	_ *EVMValidatorsRequest,
) (*EVMValidatorsResponse, error) {
	node, err := s.clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	// Get top 100 validators for now
	page := 1
	perPage := 100
	validators, err := node.Validators(context.Background(), nil, &page, &perPage)
	if err != nil {
		return nil, err
	}

	evmValidatorsResponse := EVMValidatorsResponse{}
	evmValidatorsResponse.BlockHeight = validators.BlockHeight
	evmValidatorsResponse.Validators = []ValidatorMinimal{}

	// put each validator's address and voting power to the response
	for _, validator := range validators.Validators {
		pubKeyBytes, ok := validator.PubKey.(secp256k1.PubKey)
		if !ok {
			return nil, fmt.Errorf("can't get validator public key")
		}

		if pubkey, err := crypto.DecompressPubkey(pubKeyBytes[:]); err != nil {
			return nil, err
		} else {
			evmValidatorsResponse.Validators = append(
				evmValidatorsResponse.Validators,
				ValidatorMinimal{
					Address:     crypto.PubkeyToAddress(*pubkey).String(),
					VotingPower: validator.VotingPower,
				},
			)
		}
	}

	return &evmValidatorsResponse, nil
}
