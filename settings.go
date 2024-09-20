package main

import (
	
	"fmt"
	"os"
	"encoding/json"

)

type Setting struct {

	FileSizeLimit int64 `json:"FileSizeLimitMB"` //maximum allowed upload file size
	TextSizeLimit int64 `json:"TextSizeLimitMB"` //maximum allowed upload text size
	StreamSizeLimit int64 `json:"StreamSizeLimitKB"` //streaming buffer size
	StreamThrottle int64 `json:"StreamThrottleMS"` //streaming Sleep() timer to not use too much cpu
	Pbkdf2Iteraions int `json:"Pbkdf2Iteraions"` //key derviation function iterations
	CmdUploadDefaultDurationMinute int64 `json:"CmdUploadDefaultDurationMinute"` //default file duration when uploaded through curl / other cmdline

}
 
var Global Setting;
func InitSettings(){
	
	file, err := os.Open("./data/settings.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()


	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Global)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	
}
