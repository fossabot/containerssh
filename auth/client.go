package auth

import "containerssh/protocol"

type Client interface {
	Password(
		//Username provided
		username string,
		//Password provided
		password []byte,
		//Opaque session ID to identify the login attempt
		sessionId []byte,
		//Remote address in IP:port format
		remoteAddr string,
	) (*protocol.AuthResponse, error)
	PubKey(
		//Username provided
		username string,
		//Serialized key data in SSH wire format
		pubKey []byte,
		//Opaque session ID to identify the login attempt
		sessionId []byte,
		//Remote address in IP:port format
		remoteAddr string,
	) (*protocol.AuthResponse, error)
}
