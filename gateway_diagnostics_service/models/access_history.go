package models

type AccessHistory struct {
	OrgId         int    `json:"orgId"`
	MsgType       string `json:"msgType"`
	AccessPointId int    `json:"accessPointId"`
	UserId        int    `json:"userId"`
}
