package status

type Status int

const (
	FREE Status = iota
	OCCUPIED
	RESERVED
)
