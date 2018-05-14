package main

import(
	"fmt"
	"dingapi/common"
)



// SendByDingTalkRobot 通过钉钉发送消息通知
func SendByDingTalkRobot(messageType, message, title, robotURL string) (bool, error) {
	var url string
	if robotURL == "" {
		url = "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxx"
	} else {
		url = robotURL
	}

	dingtalk := &common.DingTalkClient{
		RobotURL: url,
		Message: &common.DingTalkMessage{
			Type:    messageType,
			Message: message,
			Title:   title,
		},
	}
	ok, err := common.SendMessage(dingtalk)
	if err != nil {
		dingFields := map[string]interface{}{
			"entryType":     "DingTalkRobot",
			"dingTalkRobot": url,
		}
		fmt.Println(dingFields)
//		Logger.Error(dingFields, fmt.Sprintf("发送钉钉通知失败了: %s", err))
		return false, err
	}
	return ok, nil
}


func main(){

	msgType := "markdown"
	msg := "## golang测试msg \n\n #### 这是一条测试消息 \n\n #### 来自于golang接口 \n"
	title := "golang测试 title"
	robotURL := ""
	_, err := SendByDingTalkRobot(msgType, msg, title, robotURL)
	fmt.Println(err)



}