package entities

type Computer struct {
	ID        int64  `json:"id,omitempty"`
	Status    string `json:"status"`
	CPU       string `json:"cpu"`
	GPU       string `json:"gpu"`
	RAM       string `json:"ram"`
	SSD       string `json:"ssd"`
	HDD       string `json:"hdd"`
	Monitor   string `json:"monitor"`
	Keyboard  string `json:"keyboard"`
	Headset   string `json:"headset"`
	Mouse     string `json:"mouse"`
	Mousepad  string `json:"mousepad"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type ComputerStatus struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

type Status string

const (
	Busy        Status = "busy"
	Available   Status = "available"
	Pending     Status = "pending"
	NotWorking  Status = "not working"
	UnderRepair Status = "under repair"
)

func (s Status) IsValid() bool {
	switch s {
	case "busy", "available", "pending":
		return true
	case "not working", "under repair":
		return true
	default:
		return false
	}
}
