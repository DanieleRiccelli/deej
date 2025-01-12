package deej

type DeviceInfo struct {
	Name string
	ID   string
}

// SessionFinder represents an entity that can find all current audio sessions
type SessionFinder interface {
	GetAllSessions() ([]Session, error)
	Release() error
}
