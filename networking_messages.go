package steamworks

import (
	"syscall"
	"unsafe"
)

type ISteamNetworkingMessages struct {
	VTable *uintptr
}

func (sdk *Steamworks) ISteamNetworkingMessages() *ISteamNetworkingMessages {
	r, _, _ := sdk.LIB.NewProc("SteamAPI_GetISteamNetworkingMessages").Call(uintptr(1), uintptr(0))
	return (*ISteamNetworkingMessages)(unsafe.Pointer(r))
}

func (msgs *ISteamNetworkingMessages) SendMessageToUser(pIdentity *SteamNetworkingIdentity, pubData []byte, nDataSize int32, nChannel int32, pErrMsg *SteamNetworkingErrMsg) bool {
	r, _, _ := msgs.callVTable(0, 
		uintptr(unsafe.Pointer(pIdentity)),
		uintptr(unsafe.Pointer(&pubData[0])),
		uintptr(nDataSize),
		uintptr(nChannel),
		uintptr(unsafe.Pointer(pErrMsg)))
	return int(r) == 1
}

func (msgs *ISteamNetworkingMessages) ReceiveMessages(pOutMessages *[]*SteamNetworkingMessage, nMaxMessages int32) int32 {
	var ppMessages uintptr
	r, _, _ := msgs.callVTable(1, uintptr(unsafe.Pointer(&ppMessages)), uintptr(nMaxMessages))
	
	if r > 0 && ppMessages != 0 {
		*pOutMessages = make([]*SteamNetworkingMessage, int(r))
		for i := 0; i < int(r); i++ {
			msgPtr := *(*uintptr)(unsafe.Pointer(ppMessages + uintptr(i)*unsafe.Sizeof(uintptr(0))))
			(*pOutMessages)[i] = (*SteamNetworkingMessage)(unsafe.Pointer(msgPtr))
		}
	}
	
	return int32(r)
}

func (msgs *ISteamNetworkingMessages) FreeMessage(pMessage *SteamNetworkingMessage) {
	msgs.callVTable(2, uintptr(unsafe.Pointer(pMessage)))
}

func (msgs *ISteamNetworkingMessages) AcceptSessionWithUser(pIdentity *SteamNetworkingIdentity) bool {
	r, _, _ := msgs.callVTable(3, uintptr(unsafe.Pointer(pIdentity)))
	return int(r) == 1
}

func (msgs *ISteamNetworkingMessages) CloseSessionWithUser(pIdentity *SteamNetworkingIdentity) {
	msgs.callVTable(4, uintptr(unsafe.Pointer(pIdentity)))
}

func (msgs *ISteamNetworkingMessages) CloseChannelWithUser(pIdentity *SteamNetworkingIdentity, nChannel int32) {
	msgs.callVTable(5, uintptr(unsafe.Pointer(pIdentity)), uintptr(nChannel))
}

func (msgs *ISteamNetworkingMessages) GetSessionConnectionInfo(pIdentity *SteamNetworkingIdentity, pConnectionInfo *SteamNetworkingConnectionState) bool {
	r, _, _ := msgs.callVTable(6, uintptr(unsafe.Pointer(pIdentity)), uintptr(unsafe.Pointer(pConnectionInfo)))
	return int(r) == 1
}

func (msgs *ISteamNetworkingMessages) GetSessionUserCount() int32 {
	r, _, _ := msgs.callVTable(7)
	return int32(r)
}

func (msgs *ISteamNetworkingMessages) GetSessionUserByIndex(nIndex int32, pIdentity *SteamNetworkingIdentity) bool {
	r, _, _ := msgs.callVTable(8, uintptr(nIndex), uintptr(unsafe.Pointer(pIdentity)))
	return int(r) == 1
}

func (msgs *ISteamNetworkingMessages) callVTable(index int, args ...uintptr) (uintptr, uintptr, error) {
	vtable := *msgs.VTable
	proc := *(*uintptr)(unsafe.Pointer(vtable + uintptr(index)*unsafe.Sizeof(uintptr(0))))
	return syscall.SyscallN(proc, args...)
}