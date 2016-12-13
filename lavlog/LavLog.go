package lavlog

import "fmt"

func Error(msg interface{}){
	fmt.Println("[ERROR] : ",msg)
}
func Info(msg interface{}){
	fmt.Println("[Info] : ",msg)
}
