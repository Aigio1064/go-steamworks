package steamworks

import (
	"syscall"
	"unsafe"
)

type ISteamNetworkingUtils struct {
	VTable *uintptr
}

func (sdk *Steamworks) ISteamNetworkingUtils() *ISteamNetworkingUtils {
	r, _, _ := sdk.LIB.NewProc("SteamAPI_GetISteamNetworkingUtils").Call(uintptr(1), uintptr(0))
	return (*ISteamNetworkingUtils)(unsafe.Pointer(r))
}

func (utils *ISteamNetworkingUtils) InitRelayNetworkAccess() ESteamNetworkingAvailability {
	r, _, _ := utils.callVTable(0)
	return ESteamNetworkingAvailability(r)
}

func (utils *ISteamNetworkingUtils) GetRelayNetworkStatus() ESteamNetworkingAvailability {
	r, _, _ := utils.callVTable(1)
	return ESteamNetworkingAvailability(r)
}

func (utils *ISteamNetworkingUtils) GetLocalPeerIdentity(pIdentity *SteamNetworkingIdentity) bool {
	r, _, _ := utils.callVTable(2, uintptr(unsafe.Pointer(pIdentity)))
	return int(r) == 1
}

func (utils *ISteamNetworkingUtils) GetIdentityAsString(pIdentity *SteamNetworkingIdentity) string {
	r, _, _ := utils.callVTable(3, uintptr(unsafe.Pointer(pIdentity)))
	return CStrToString(r)
}

func (utils *ISteamNetworkingUtils) ParseIdentityString(pchIdentity string, pIdentity *SteamNetworkingIdentity) bool {
	r, _, _ := utils.callVTable(4, StringToCCharPtr(pchIdentity), uintptr(unsafe.Pointer(pIdentity)))
	return int(r) == 1
}

func (utils *ISteamNetworkingUtils) GetConfigValueInfo(nValue int32, pOutInfo *[256]byte) bool {
	r, _, _ := utils.callVTable(5, uintptr(nValue), uintptr(unsafe.Pointer(pOutInfo)))
	return int(r) == 1
}

func (utils *ISteamNetworkingUtils) SetConfigValue(pIdentity *SteamNetworkingIdentity, nValue int32, pArg unsafe.Pointer) bool {
	r, _, _ := utils.callVTable(6, uintptr(unsafe.Pointer(pIdentity)), uintptr(nValue), uintptr(pArg))
	return int(r) == 1
}

func (utils *ISteamNetworkingUtils) GetConfigValue(pIdentity *SteamNetworkingIdentity, nValue int32, pArg unsafe.Pointer) bool {
	r, _, _ := utils.callVTable(7, uintptr(unsafe.Pointer(pIdentity)), uintptr(nValue), uintptr(pArg))
	return int(r) == 1
}

func (utils *ISteamNetworkingUtils) GetSessionConnectionInfo(hConn HSteamNetConnection, pConnectionInfo *SteamNetworkingConnectionState) bool {
	r, _, _ := utils.callVTable(8, uintptr(hConn), uintptr(unsafe.Pointer(pConnectionInfo)))
	return int(r) == 1
}

func (utils *ISteamNetworkingUtils) GetQuickConnectionStatus(hConn HSteamNetConnection) SteamNetConnectionStatus {
	r, _, _ := utils.callVTable(9, uintptr(hConn))
	return SteamNetConnectionStatus(r)
}

func (utils *ISteamNetworkingUtils) SendMessages(nMessages int32, pMessages *SteamNetworkingMessage, pOutResults *int32) {
	utils.callVTable(10, uintptr(nMessages), uintptr(unsafe.Pointer(pMessages)), uintptr(unsafe.Pointer(pOutResults)))
}

func (utils *ISteamNetworkingUtils) ReceiveMessagesOnChannel(nChannel int32, ppOutMessages *uintptr, nMaxMessages int32) int32 {
	r, _, _ := utils.callVTable(11, uintptr(nChannel), uintptr(unsafe.Pointer(ppOutMessages)), uintptr(nMaxMessages))
	return int32(r)
}

func (utils *ISteamNetworkingUtils) FreeMessage(pMessage *SteamNetworkingMessage) {
	utils.callVTable(12, uintptr(unsafe.Pointer(pMessage)))
}

func (utils *ISteamNetworkingUtils) GetFakeIPForIdentity(pIdentity *SteamNetworkingIdentity, pResult *SteamNetworkingFakeIPResult) bool {
	r, _, _ := utils.callVTable(13, uintptr(unsafe.Pointer(pIdentity)), uintptr(unsafe.Pointer(pResult)))
	return int(r) == 1
}

func (utils *ISteamNetworkingUtils) callVTable(index int, args ...uintptr) (uintptr, uintptr, error) {
	vtable := *utils.VTable
	proc := *(*uintptr)(unsafe.Pointer(vtable + uintptr(index)*unsafe.Sizeof(uintptr(0))))
	return callProc(proc, args...)
}

func callProc(proc uintptr, args ...uintptr) (uintptr, uintptr, error) {
	return syscall.SyscallN(proc, args...)
}