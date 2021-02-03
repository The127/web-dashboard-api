package models

type DashboardSubItem struct {
	Id          string `json:"Id"`
	ItemId      string `json:"ItemId"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	WebsiteUrl  string `json:"WebsiteUrl"`
}
