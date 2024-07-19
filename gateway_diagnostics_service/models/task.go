package models

type Task struct {
	ServiceName string `json:"serviceName"`
	Data        string `json:"data"`
}

type CreateTaskRequest struct {
	ServiceName string `json:"serviceName" binding:"required"`
	Data        string `json:"data" binding:"required"`
}

type CreateMessageRequest struct {
	Version     int    `json:"version" binding:"required"`
	MsgType     string `json:"msgType" binding:"required"`
	RequestId   string `json:"requestId" binding:"required"`
	ServiceType int    `json:"serviceType" binding:"required"`
	StatusCode  int    `json:"statusCode" binding:"required"`
}

type AckEvent struct {
	Version     int    `json:"version" binding:"required"`
	MsgType     string `json:"msgType" binding:"required"`
	RequestId   string `json:"requestId" binding:"required"`
	ServiceType int    `json:"serviceType" binding:"required"`
	StatusCode  int    `json:"statusCode" binding:"required"`
}

type DcmEvent struct {
	MessageVersion int            `json:"messageVersion"`
	MessageData    DcmMessageData `json:"messageData"`
}

type DcmMessageData struct {
	RequestId   string  `json:"requestId"`
	MsgType     string  `json:"msgType"`
	DataVersion int     `json:"dataVersion"`
	Data        DcmData `json:"data"`
}

type DcmData struct {
	SerialNumber string `json:"serialNumber"`
	EventTime    int    `json:"eventTime"`
}
