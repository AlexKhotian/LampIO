package Versioning

// Version of the client
type Version struct {
	Minor int32 `json:"Minor"`
	Major int32 `json:"Major"`
}

// Current versions
const (
	ClientVersionMinor = 1
	ClientVersionMajor = 0
)
