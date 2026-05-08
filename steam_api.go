package steamworks

func (sdk *Steamworks) InitFlat() {
	sdk.dll.NewProc("SteamAPI_InitFlat").Call()
}
