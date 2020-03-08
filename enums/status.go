package enums

type Status string

const (
	Pending            Status = "pending"
	Approved           Status = "approved"
	Declined           Status = "declined"
	ApprovedButStalled Status = "approved_but_stalled"
)

type StatusID int

const (
	New StatusID = iota
	InWork
	Waiting
	Closed
)

type StatusSales string

const (
	NewSale                  StatusSales = "new"
	InvalidQR                StatusSales = "invalid_qr"
	ReceiptInfoFound         StatusSales = "receipt_info_found"
	ReceiptInfoNotFound      StatusSales = "receipt_info_not_found"
	ReceiptInfoIsNotYetReady StatusSales = "receipt_info_is_not_yet_ready"
	FailSale                 StatusSales = "fail"
	PaymentsCreated          StatusSales = "payments_created"
	PaymentsNotFound         StatusSales = "payments_not_found"
)
