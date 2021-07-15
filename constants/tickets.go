package constants

type draftTicketStatus struct {
	Name  string
	Order int64
	Type  TicketType
}

var DefaultTicketStatuses = []draftTicketStatus{
	{
		Name:  "Draft",
		Order: 0,
		Type:  TicketPending,
	},
	{
		Name:  "Need to Order Part",
		Order: 1,
		Type:  TicketPending,
	},
	{
		Name:  "In Queue",
		Order: 2,
		Type:  TicketPending,
	},
	{
		Name:  "On Hold for Part",
		Order: 3,
		Type:  TicketPending,
	},
	{
		Name:  "Need to Approve",
		Order: 4,
		Type:  TicketPending,
	},

	{
		Name:  "Ready for Pickup",
		Order: 5,
		Type:  TicketCompleted,
	},
	{
		Name:  "Archived",
		Order: 6,
		Type:  TicketCompleted,
	},
	{
		Name:  "Completed",
		Order: 7,
		Type:  TicketCompleted,
	},
	{
		Name:  "Warranty Repair",
		Order: 8,
		Type:  TicketCompleted,
	},
}

const (
	TicketUnpaid int64 = iota + 1
	TicketPaid
)

const (
	ServiceDefault = "1"
)

type TicketType int64

const (
	TicketPending TicketType = iota
	TicketCompleted
)

func (t TicketType) Int64() int64 {
	return int64(t)
}
