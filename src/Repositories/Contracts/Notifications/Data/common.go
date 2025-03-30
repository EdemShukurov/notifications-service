package data

// Storage for errors
type ErrorDbo struct {
	Type    string
	Message string
}

// Storage for reference info
type ReferenceDbo struct {
	Id   string
	Type string
}
