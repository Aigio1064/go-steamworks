package steamworks

import (
	"syscall"
	"unsafe"
)

type ISteamNetworkingSockets struct {
	VTable *uintptr
}

func (sdk *Steamworks) ISteamNetworkingSockets() *ISteamNetworkingSockets {
	r, _, _ := sdk.LIB.NewProc("SteamAPI_GetISteamNetworkingSockets").Call(uintptr(1), uintptr(0))
	return (*ISteamNetworkingSockets)(unsafe.Pointer(r))
}

func (sockets *ISteamNetworkingSockets) CreateListenSocketIP(nPort uint16, pOptions *SteamNetworkingConfig, pErrMsg *SteamNetworkingErrMsg) HSteamListenSocket {
	var opts *SteamNetworkingConfigValue_t
	var count int32
	if pOptions != nil {
		opts = pOptions.Data()
		count = pOptions.Count()
	}
	
	r, _, _ := sockets.callVTable(0, 
		uintptr(nPort),
		uintptr(unsafe.Pointer(opts)),
		uintptr(count),
		uintptr(unsafe.Pointer(pErrMsg)))
	return HSteamListenSocket(r)
}

func (sockets *ISteamNetworkingSockets) CreateListenSocketP2P(nVirtualPort int32, pOptions *SteamNetworkingConfig, pErrMsg *SteamNetworkingErrMsg) HSteamListenSocket {
	var opts *SteamNetworkingConfigValue_t
	var count int32
	if pOptions != nil {
		opts = pOptions.Data()
		count = pOptions.Count()
	}
	
	r, _, _ := sockets.callVTable(1, 
		uintptr(nVirtualPort),
		uintptr(unsafe.Pointer(opts)),
		uintptr(count),
		uintptr(unsafe.Pointer(pErrMsg)))
	return HSteamListenSocket(r)
}

func (sockets *ISteamNetworkingSockets) DestroyListenSocket(hSocket HSteamListenSocket) {
	sockets.callVTable(2, uintptr(hSocket))
}

func (sockets *ISteamNetworkingSockets) ConnectByIPAddress(pAddr *SteamNetworkingIdentity, pOptions *SteamNetworkingConfig, pErrMsg *SteamNetworkingErrMsg) HSteamNetConnection {
	var opts *SteamNetworkingConfigValue_t
	var count int32
	if pOptions != nil {
		opts = pOptions.Data()
		count = pOptions.Count()
	}
	
	r, _, _ := sockets.callVTable(3, 
		uintptr(unsafe.Pointer(pAddr)),
		uintptr(unsafe.Pointer(opts)),
		uintptr(count),
		uintptr(unsafe.Pointer(pErrMsg)))
	return HSteamNetConnection(r)
}

func (sockets *ISteamNetworkingSockets) ConnectBySteamID(pSteamID uint64, pOptions *SteamNetworkingConfig, pErrMsg *SteamNetworkingErrMsg) HSteamNetConnection {
	var opts *SteamNetworkingConfigValue_t
	var count int32
	if pOptions != nil {
		opts = pOptions.Data()
		count = pOptions.Count()
	}
	
	r, _, _ := sockets.callVTable(4, 
		uintptr(pSteamID),
		uintptr(unsafe.Pointer(opts)),
		uintptr(count),
		uintptr(unsafe.Pointer(pErrMsg)))
	return HSteamNetConnection(r)
}

func (sockets *ISteamNetworkingSockets) ConnectP2P(pIdentity *SteamNetworkingIdentity, nVirtualPort int32, pOptions *SteamNetworkingConfig, pErrMsg *SteamNetworkingErrMsg) HSteamNetConnection {
	var opts *SteamNetworkingConfigValue_t
	var count int32
	if pOptions != nil {
		opts = pOptions.Data()
		count = pOptions.Count()
	}
	
	r, _, _ := sockets.callVTable(5, 
		uintptr(unsafe.Pointer(pIdentity)),
		uintptr(nVirtualPort),
		uintptr(unsafe.Pointer(opts)),
		uintptr(count),
		uintptr(unsafe.Pointer(pErrMsg)))
	return HSteamNetConnection(r)
}

func (sockets *ISteamNetworkingSockets) CloseConnection(hConn HSteamNetConnection, nReason int32, pDebugMsg string, bEnableLinger bool) {
	sockets.callVTable(6, 
		uintptr(hConn),
		uintptr(nReason),
		StringToCCharPtr(pDebugMsg),
		boolToUintptr(bEnableLinger))
}

func (sockets *ISteamNetworkingSockets) SendMessageOnConnection(hConn HSteamNetConnection, pubData []byte, nDataSize int32, nFlags int32, pErrMsg *SteamNetworkingErrMsg) bool {
	r, _, _ := sockets.callVTable(7, 
		uintptr(hConn),
		uintptr(unsafe.Pointer(&pubData[0])),
		uintptr(nDataSize),
		uintptr(nFlags),
		uintptr(unsafe.Pointer(pErrMsg)))
	return int(r) == 1
}

func (sockets *ISteamNetworkingSockets) ReceiveMessagesOnConnection(hConn HSteamNetConnection, ppOutMessages *[]*SteamNetworkingMessage, nMaxMessages int32) int32 {
	var ppMessages uintptr
	r, _, _ := sockets.callVTable(8, uintptr(hConn), uintptr(unsafe.Pointer(&ppMessages)), uintptr(nMaxMessages))
	
	if r > 0 && ppMessages != 0 {
		*ppOutMessages = make([]*SteamNetworkingMessage, int(r))
		for i := 0; i < int(r); i++ {
			msgPtr := *(*uintptr)(unsafe.Pointer(ppMessages + uintptr(i)*unsafe.Sizeof(uintptr(0))))
			(*ppOutMessages)[i] = (*SteamNetworkingMessage)(unsafe.Pointer(msgPtr))
		}
	}
	
	return int32(r)
}

func (sockets *ISteamNetworkingSockets) CreatePollGroup() HSteamNetPollGroup {
	r, _, _ := sockets.callVTable(9)
	return HSteamNetPollGroup(r)
}

func (sockets *ISteamNetworkingSockets) DestroyPollGroup(hPollGroup HSteamNetPollGroup) {
	sockets.callVTable(10, uintptr(hPollGroup))
}

func (sockets *ISteamNetworkingSockets) SetConnectionPollGroup(hConn HSteamNetConnection, hPollGroup HSteamNetPollGroup) bool {
	r, _, _ := sockets.callVTable(11, uintptr(hConn), uintptr(hPollGroup))
	return int(r) == 1
}

func (sockets *ISteamNetworkingSockets) RunMessageLoop() {
	sockets.callVTable(12)
}

func (sockets *ISteamNetworkingSockets) GetConnectionInfo(hConn HSteamNetConnection, pConnectionInfo *SteamNetworkingConnectionState) bool {
	r, _, _ := sockets.callVTable(13, uintptr(hConn), uintptr(unsafe.Pointer(pConnectionInfo)))
	return int(r) == 1
}

func (sockets *ISteamNetworkingSockets) GetListenSocketInfo(hSocket HSteamListenSocket, pInfo *SteamNetworkingFakeIPResult) bool {
	r, _, _ := sockets.callVTable(14, uintptr(hSocket), uintptr(unsafe.Pointer(pInfo)))
	return int(r) == 1
}

func (sockets *ISteamNetworkingSockets) callVTable(index int, args ...uintptr) (uintptr, uintptr, error) {
	vtable := *sockets.VTable
	proc := *(*uintptr)(unsafe.Pointer(vtable + uintptr(index)*unsafe.Sizeof(uintptr(0))))
	return syscall.SyscallN(proc, args...)
}

func boolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}