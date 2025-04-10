package entities

type Booking struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	ComputerID int64  `json:"computer_id"`
	Package   string `json:"package"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	TotalPrice int    `json:"total_price"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`

}