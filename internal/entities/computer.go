package entities

type ComputerSpecs struct {
	ID        int64  `json:"id"`
	CPU       string `json:"cpu"`
	GPU       string `json:"gpu"`
	RAM       string `json:"ram"`
	SSD       string `json:"ssd"`
	HDD       string `json:"hdd"`
	Monitor   string `json:"monitor"`
	Keyboard  string `json:"keyboard"`
	Headset   string `json:"headset"`
	Mouse     string `json:"mouse"`
	CreatedAt string `json:"created_at"`
}

type ComputerStatus struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}
