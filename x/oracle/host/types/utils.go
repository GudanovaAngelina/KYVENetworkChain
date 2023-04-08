package types

import (
	"encoding/json"

	"github.com/cosmos/gogoproto/jsonpb"

	// Auth
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	// Global
	globalTypes "github.com/KYVENetwork/chain/x/global/types"
	// IBC Core
	channelTypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	// IBC Transfer
	transferTypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
)

func ParseOraclePacket(packet channelTypes.Packet) (
	data *transferTypes.FungibleTokenPacketData, req *OracleQuery, valid bool, err error,
) {
	if err := json.Unmarshal(packet.Data, &data); err != nil {
		return nil, nil, false, nil
	}
	if data.Receiver != authTypes.NewModuleAddress(ModuleName).String() {
		return data, nil, false, nil
	}

	// Check -- Token Denom
	trace := transferTypes.ParseDenomTrace(data.Denom)
	isNativeToken := transferTypes.ReceiverChainIsSource(packet.SourcePort, packet.SourceChannel, data.Denom)
	isNativeKYVE := isNativeToken && trace.BaseDenom == globalTypes.Denom

	if !isNativeKYVE {
		return data, nil, false, ErrInvalidOracleToken
	}

	// Check -- Oracle Instructions
	var memo OracleMemo
	if err := jsonpb.UnmarshalString(data.GetMemo(), &memo); err != nil {
		return nil, nil, false, ErrInvalidOracleMemo
	}

	return data, memo.Query, true, nil
}