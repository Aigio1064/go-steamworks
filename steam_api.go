package steamworks

import "unsafe"

type ESteamAPIInitResult int

const (
	K_ESteamAPIInitResult_OK              ESteamAPIInitResult = 0
	K_ESteamAPIInitResult_FailedGeneric   ESteamAPIInitResult = 1 // Some other failure
	K_ESteamAPIInitResult_NoSteamClient   ESteamAPIInitResult = 2 // We cannot connect to Steam, steam probably isn't running
	K_ESteamAPIInitResult_VersionMismatch ESteamAPIInitResult = 3 // Steam client appears to be out of date
)

func (sdk *Steamworks) InitFlat() ESteamAPIInitResult {
	r, _, err := sdk.dll.NewProc("SteamAPI_InitFlat").Call()
	Panic(err)
	return *(*ESteamAPIInitResult)(unsafe.Pointer(r))
}
