package model

type OperationLog struct {
	Universal
	OptModule     string `json:"optModule"`
	OptType       string `json:"optType"`
	OptUrl        string `json:"optUrl"`
	OptMethod     string `json:"optMethod"`
	OptDesc       string `json:"optDesc"`
	RequestParam  string `json:"requestParam"`
	RequestMethod string `json:"requestMethod"`
	ResponseData  string `json:"responseData"`
	UserId        int    `json:"userId"`
	Nickname      string `json:"nickname"`
	IpAddress     string `json:"ipAddress"`
	IpSource      string `json:"ipSource"`
}
