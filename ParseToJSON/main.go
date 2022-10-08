package main

import (
	"encoding/json"
	"fmt"
)

//含有多个嵌套，需要层层解析，使用递归更为高效
func ParseToJSON(jsonVar map[string]interface{}) interface{} {
	//_, ok := map1[key1] 如果key1存在则ok == true，否则ok为false
	var index = make(map[string]interface{})
	v1, ok := jsonVar["fields"]
	//开始一层一层剥开map
	//存在interface，ok=true
	if ok {
		//找[]interface{}
		v2, _ := v1.([]interface{}) //v2 保存[]interface{}
		for _, temp := range v2 {
			//在[]interface中寻找map,key为string，value为interface{}
			i, ok := temp.(map[string]interface{})
			//存在map,key为string，value为interface{}
			if ok {
				//查找key及model，此处的key不是map中的key
				v3, _ := i["key"].(string)
				v4, _ := i["model"].(map[string]interface{})
				//将key于model对应
				index[v3] = ParseToJSON(v4)
			}
		}
	} else {
		//当map的key不再有fields
		//判断值类型
		value, ok := jsonVar["type"].(string)
		if ok {
			switch value {
			case "number":
				return 111
			case "string":
				return "Sample Text"
			case "object":
				return nil
			case "boolclean":
				return true
			case "array":
				return "null"
			}
		}
	}
	return index
}

func main() {
	originStr := "{\"fields\":[{\"key\":\"code\",\"model\":{\"type\":\"number\"}},{\"key\":\"data\",\"model\":{\"fields\":[{\"key\":\"event\",\"model\":{\"fields\":[{\"key\":\"attendee_ability\",\"model\":{\"type\":\"string\"}},{\"key\":\"color\",\"model\":{\"type\":\"number\"}},{\"key\":\"description\",\"model\":{\"type\":\"string\"}},{\"key\":\"end_time\",\"model\":{\"fields\":[{\"key\":\"date\",\"model\":{\"type\":\"string\"}},{\"key\":\"timestamp\",\"model\":{\"type\":\"string\"}},{\"key\":\"timezone\",\"model\":{\"type\":\"string\"}}],\"type\":\"object\"}},{\"key\":\"event_id\",\"model\":{\"type\":\"string\"}},{\"key\":\"free_busy_status\",\"model\":{\"type\":\"string\"}},{\"key\":\"is_exception\",\"model\":{\"type\":\"boolean\"}},{\"key\":\"location\",\"model\":{\"fields\":[{\"key\":\"address\",\"model\":{\"type\":\"string\"}},{\"key\":\"latitude\",\"model\":{\"type\":\"number\"}},{\"key\":\"longitude\",\"model\":{\"type\":\"number\"}},{\"key\":\"name\",\"model\":{\"type\":\"string\"}}],\"type\":\"object\"}},{\"key\":\"need_notification\",\"model\":{\"type\":\"boolean\"}},{\"key\":\"recurrence\",\"model\":{\"type\":\"string\"}},{\"key\":\"recurring_event_id\",\"model\":{\"type\":\"string\"}},{\"key\":\"reminders\",\"model\":{\"items\":{\"fields\":[{\"key\":\"minutes\",\"model\":{\"type\":\"number\"}}],\"type\":\"object\"},\"type\":\"array\"}},{\"key\":\"schemas\",\"model\":{\"items\":{\"fields\":[{\"key\":\"app_link\",\"model\":{\"type\":\"string\"}},{\"key\":\"ui_name\",\"model\":{\"type\":\"string\"}},{\"key\":\"ui_status\",\"model\":{\"type\":\"string\"}}],\"type\":\"object\"},\"type\":\"array\"}},{\"key\":\"start_time\",\"model\":{\"fields\":[{\"key\":\"date\",\"model\":{\"type\":\"string\"}},{\"key\":\"timestamp\",\"model\":{\"type\":\"string\"}},{\"key\":\"timezone\",\"model\":{\"type\":\"string\"}}],\"type\":\"object\"}},{\"key\":\"status\",\"model\":{\"type\":\"string\"}},{\"key\":\"summary\",\"model\":{\"type\":\"string\"}},{\"key\":\"vchat\",\"model\":{\"fields\":[{\"key\":\"description\",\"model\":{\"type\":\"string\"}},{\"key\":\"icon_type\",\"model\":{\"type\":\"string\"}},{\"key\":\"meeting_url\",\"model\":{\"type\":\"string\"}},{\"key\":\"vc_type\",\"model\":{\"type\":\"string\"}}],\"type\":\"object\"}},{\"key\":\"visibility\",\"model\":{\"type\":\"string\"}}],\"type\":\"object\"}}],\"type\":\"object\"}},{\"key\":\"msg\",\"model\":{\"type\":\"string\"}}],\"type\":\"object\"}"
	jsonVar := make(map[string]interface{})
	json.Unmarshal([]byte(originStr), &jsonVar)
	jsonStr, _ := json.Marshal(ParseToJSON(jsonVar))
	fmt.Println("JSON example is: ", string(jsonStr))
}

//code，data,msg三个key，data对应的value是一大堆map
// {"code":111,"data":{"event":{"attendee_ability":"Sample text","color":111,"description":"Sample text","end_time":{"date":"Sample text","timestamp":"Sample text","timezone":"Sample text"},"event_id":"Sample text","free_busy_status":"Sample text","is_exception":true,"location":{"address":"Sample text","latitude":111,"longitude":111,"name":"Sample text"},"need_notification":true,"recurrence":"Sample text","recurring_event_id":"Sample text","reminders":null,"schemas":null,"start_time":{"date":"Sample text","timestamp":"Sample text","timezone":"Sample text"},"status":"Sample text","summary":"Sample text","vchat":{"description":"Sample text","icon_type":"Sample text","meeting_url":"Sample text","vc_type":"Sample text"},"visibility":"Sample text"}},"msg":"Sample text"}
