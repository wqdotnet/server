package netwrok

// import ()

//NetworkInterface network
type NetworkInterface interface {
	Start()
	Stop()
	Send(msg *NetworkMsg)
}

//NetworkMsg   tcp udp send/receive msg
type NetworkMsg struct {
	// recvbuf []byte
	// bufptr  []byte
	Module int
	method int
	buf    []byte
}
