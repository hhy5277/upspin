// Package proto contains the definitions shared between RPC store server and client,
// one pair for each remote call.
// TODO: Maybe move to gprc?
package proto

import "upspin.googlesource.com/upspin.git/upspin"

type GetRequest struct {
	Reference upspin.Reference
}

type GetResponse struct {
	Data      []byte
	Locations []upspin.Location
}

type PutRequest struct {
	Data []byte
}

type PutResponse struct {
	Reference upspin.Reference
}

type DeleteRequest struct {
	Reference upspin.Reference
}

type DeleteResponse struct {
}

// Types for the methods of upspin.Service.
// TODO: Put somewhere shared?

type ConfigureRequest struct {
	Options []string
}

type ConfigureResponse struct {
}

type EndpointRequest struct {
}

type EndpointResponse struct {
	Endpoint upspin.Endpoint
}

type ServerUserNameRequest struct {
}

type ServerUserNameResponse struct {
	UserName string
}