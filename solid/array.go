package solid

type Notify struct {
	List [][]string `json:"list"`
}

// {"remoteServer":[],
//"notifyInfoType":[["device_alarm","cpu_usage"],["device_alarm","cpu_temp"],["device_alarm","mem_usage"],
//["device_alarm","disk_usage"],["device_alarm","disk_status"],["device_alarm","network_card_status"],["device_alarm","network_card_flow"],["device_alarm","fan_status"],["device_alarm","power_status"]]}

func Marshal() Notify {

	str1 := []string{"device_alarm", "cpu_usage"}
	str2 := []string{"device_alarm", "network_card_status"}

	myList := make([][]string, 2)

	myList[0] = str1
	myList[1] = str2

	return Notify{List: myList}

}
