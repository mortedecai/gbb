package models

// GBBDeleteData represents the structure requesting a deletion from a bit burner server.
type GBBDeleteData struct {
	// Path is the filename to delete from the specified server.
	Path GBBFileName `json:"filename"`
	// Server is the bit burner server to delete the file from.
	//   * Must be the IP or name of a server in the bit burner game.
	Server string `json:"home,omitempty"`
}
