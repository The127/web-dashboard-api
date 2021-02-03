package models

type DashboardItem struct {
	Id          string `json:"Id"`
	GroupId     string `json:"GroupId"`
	ImageUrl    string `json:"ImageUrl"`
	IconUrl     string `json:"IconUrl"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	WebsiteUrl  string `json:"WebisteUrl"`
}
