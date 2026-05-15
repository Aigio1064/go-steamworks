package steamworks

import "unsafe"

type HSteamNetConnection uint64
type HSteamListenSocket uint64
type HSteamNetPollGroup uint64

type SteamNetConnectionStatus int32

const (
	SteamNetConnectionStatus_None                   SteamNetConnectionStatus = 0
	SteamNetConnectionStatus_Connecting             SteamNetConnectionStatus = 1
	SteamNetConnectionStatus_FindingRoute           SteamNetConnectionStatus = 2
	SteamNetConnectionStatus_Connected              SteamNetConnectionStatus = 3
	SteamNetConnectionStatus_ClosedByPeer           SteamNetConnectionStatus = 4
	SteamNetConnectionStatus_ProblemDetectedLocally SteamNetConnectionStatus = 5
	SteamNetConnectionStatus_ClosedByUser           SteamNetConnectionStatus = 6
)

type SteamNetworkingConnectionState struct {
	EConnectionState       SteamNetConnectionStatus
	EEndReason            int32
	SzEndDebugString      [256]byte
	NPing                int32
	NPacketsLost         int32
	FlConnectionQuality  float32
	NOutgoingPacketsQueued int32
}

type SteamNetworkingIdentity struct {
	EType int32
	URaw  [64]byte
}

type SteamNetworkingMessage struct {
	NDataSize int32
	PData     uintptr
	Conn      HSteamNetConnection
	Identity  SteamNetworkingIdentity
	NRemoteTimestamp int32
}

type SteamNetworkingConfigValue_t struct {
	EValueType int32
	ValInt64   int64
	ValFloat   float64
	PchString  uintptr
}

type SteamNetworkingErrMsg struct {
	SzMsg [512]byte
}

type SteamNetworkingConfig struct {
	config []SteamNetworkingConfigValue_t
}

func (cfg *SteamNetworkingConfig) SetInt(key int32, value int64) {
	cfg.config = append(cfg.config, SteamNetworkingConfigValue_t{
		EValueType: 1,
		ValInt64:   value,
	})
}

func (cfg *SteamNetworkingConfig) SetFloat(key int32, value float64) {
	cfg.config = append(cfg.config, SteamNetworkingConfigValue_t{
		EValueType: 2,
		ValFloat:   value,
	})
}

func (cfg *SteamNetworkingConfig) SetString(key int32, value string) {
	bytes := append([]byte(value), 0)
	cfg.config = append(cfg.config, SteamNetworkingConfigValue_t{
		EValueType: 3,
		PchString:  uintptr(unsafe.Pointer(&bytes[0])),
	})
}

func (cfg *SteamNetworkingConfig) Data() *SteamNetworkingConfigValue_t {
	if len(cfg.config) == 0 {
		return nil
	}
	return &cfg.config[0]
}

func (cfg *SteamNetworkingConfig) Count() int32 {
	return int32(len(cfg.config))
}

type ESteamNetworkingAvailability int32

const (
	ESteamNetworkingAvailability_None       ESteamNetworkingAvailability = 0
	ESteamNetworkingAvailability_Partial    ESteamNetworkingAvailability = 1
	ESteamNetworkingAvailability_Connected  ESteamNetworkingAvailability = 2
)

type SteamNetworkingFakeIPResult struct {
	BLocal    bool
	NIP       uint32
	NPort     uint16
	NScopeID  uint32
}

type SteamDatagramRelayAuthTicket struct {
	HAuthTicket uint64
}