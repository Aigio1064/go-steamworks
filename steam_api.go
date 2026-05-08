package steamworks

type ESteamAPIInitResult int

const (
	K_ESteamAPIInitResult_OK              ESteamAPIInitResult = 0
	K_ESteamAPIInitResult_FailedGeneric   ESteamAPIInitResult = 1 // Some other failure
	K_ESteamAPIInitResult_NoSteamClient   ESteamAPIInitResult = 2 // We cannot connect to Steam, steam probably isn't running
	K_ESteamAPIInitResult_VersionMismatch ESteamAPIInitResult = 3 // Steam client appears to be out of date
)

func (sdk *Steamworks) InitFlat(errMsg *string) ESteamAPIInitResult {
	var (
		r uintptr
	)
	proc := sdk.LIB.NewProc("SteamAPI_InitFlat")
	if errMsg != nil {
		r, _, _ = proc.Call(StringToCCharPtr(*errMsg))
	} else {
		r, _, _ = proc.Call()
	}
	return ESteamAPIInitResult(r)
}
func (sdk *Steamworks) Init() bool {
	return sdk.InitFlat(nil) == K_ESteamAPIInitResult_OK
}
func (sdk *Steamworks) InitSafe() bool {
	r, _, _ := sdk.LIB.NewProc("SteamAPI_InitSafe").Call()
	return int(r) == 1
}
func (sdk *Steamworks) Shutdown() {
	sdk.LIB.NewProc("SteamAPI_Shutdown").Call()
}
func (sdk *Steamworks) RestartAppIfNecessary() bool {
	r, _, _ := sdk.LIB.NewProc("SteamAPI_RestartAppIfNecessary").Call()
	return int(r) == 1
}
func (sdk *Steamworks) ReleaseCurrentThreadMemory() {
	sdk.LIB.NewProc("SteamAPI_ReleaseCurrentThreadMemory").Call()
}
func (sdk *Steamworks) IsSteamRunning() bool {
	r, _, _ := sdk.LIB.NewProc("SteamAPI_IsSteamRunning").Call()
	return int(r) == 1
}
func (sdk *Steamworks) GetSteamInstallPath() string {
	r, _, _ := sdk.LIB.NewProc("SteamAPI_GetSteamInstallPath").Call()
	return CStrToString(r)
}
