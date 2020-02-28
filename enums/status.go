package enums

type Status string

const (
	Pending            Status = "pending"
	Approved           Status = "approved"
	Declined           Status = "declined"
	ApprovedButStalled Status = "approved_but_stalled"
)
