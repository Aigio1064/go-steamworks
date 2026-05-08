package steamworks

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/ying32/dylib"
)

//go:embed bin/libsteam_api_windows_amd64.dll
var libData []byte

const libName = "steam_api64.dll"

func InitSteamworks(appid string) *Steamworks {
	dir, err := os.Getwd()
	Panic(err)
	dir = filepath.Join(dir, libName)
	_, err = os.Open(libName)
	if os.IsNotExist(err) {
		err = os.WriteFile(dir, libData, 0755)
		Panic(err)
	}
	if appid != "" {
		err = os.Setenv("SteamAppId", appid)
		Panic(err)
		err = os.Setenv("SteamGameId", appid)
		Panic(err)
	}
	dll := dylib.NewLazyDLL(libName)
	err = dll.Load()
	Panic(err)
	return &Steamworks{LIB: dll}
}
