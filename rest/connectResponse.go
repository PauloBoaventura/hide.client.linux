package rest

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type ConnectResponse struct {
	PublicKey					[]byte			`json:"publicKey"`							// Server's public key
	Endpoint					net.UDPAddr		`json:"endpoint,omitempty"`					// Server's endpoint
	PresharedKey				[]byte			`json:"presharedKey,omitempty"`				// Preshared key ( if negotiated )
	PersistentKeepaliveInterval	time.Duration	`json:"persistentKeepalive,omitempty"`		// PersistentKeepAlive interval
	AllowedIps					[]net.IP		`json:"allowedIps,omitempty"`				// Server assigned IPs
	DNS							[]net.IP		`json:"DNS,omitempty"`						// Server assigned DNS server IPs
	Gateway						[]net.IP		`json:"gateway,omitempty"`					// Server assigned gateway IPs
	StaleAccessToken			bool			`json:"staleAccessToken,omitempty"`			// If true the Access-Token presented in the ConnectRequest is stale and should be updated
	SessionToken				[]byte			`json:"sessionToken,omitempty"`				// Session-Token uniquely identifies the VPN connection session
}

func ( w *ConnectResponse ) printSlice( ips []net.IP ) string {
	ipStrings := make( []string, len( ips ) )
	for i, ip := range ips { ipStrings[i] = ip.String() }
	return strings.Join( ipStrings, ", ")
}

func ( w *ConnectResponse ) Print() {
	fmt.Println( "Rest: Remote UDP endpoint is", w.Endpoint.String() )
	fmt.Println( "Rest: Keepalive is", w.PersistentKeepaliveInterval.Seconds(), "seconds" )
	fmt.Println( "Rest: Assigned IPs are", w.printSlice( w.AllowedIps ) )
	fmt.Println( "Rest: Gateway IPs are", w.printSlice( w.Gateway ) )
	fmt.Println( "Rest: DNS servers are", w.printSlice( w.DNS ) )
	if w.StaleAccessToken{ fmt.Println( "Conn: Access-Token is stale" ) }
	// fmt.Println( "Conn: Session-Token is", base64.StdEncoding.EncodeToString( w.SessionToken ) )
}