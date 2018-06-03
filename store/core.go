package store

// Store gives simple functionality for operate persistance layer
type Store interface {
	Get(string) Device
	Save(Device) string
}

// Device represents abstract IoT item
type Device struct {
	ID   string `json:"id"`
	Name string `json:"name,omitemty"`
	Desc string `json:"desc,omitemty"`
	IP   string `json:"ip"`
}
