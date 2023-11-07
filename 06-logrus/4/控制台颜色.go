package main

import "fmt"

// 自建函数打印对应颜色的字符串

const (
	Black int = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
)

// _ColorString  返回带颜色的格式化字符串，color : 对应的颜色代码, mode : true=>背景色 false=> 前景色 text :变色文本
func _ColorString(color int, mode bool, text string) string {
	var modeCode int
	if modeCode = 3; mode {
		modeCode = 4
	}
	return fmt.Sprintf("\033[%d%dm%s\033[0m", modeCode, color, text)
}

func main() {

	fmt.Println(_ColorString(Blue, false, "南城真帅"))

	//fmt.Println("\033[30m This is 黑  \033[0m")
	//fmt.Println("\033[31m This is 红  \033[0m")
	//fmt.Println("\033[32m This is 绿  \033[0m")
	//fmt.Println("\033[33m This is 黄  \033[0m")
	//fmt.Println("\033[34m This is 蓝  \033[0m")
	//fmt.Println("\033[35m This is 紫  \033[0m")
	//fmt.Println("\033[36m This is 蓝绿  \033[0m")
	//fmt.Println("\033[37m This is 灰  \033[0m")
	//
	//fmt.Println("\033[40m This is 黑  \033[0m")
	//fmt.Println("\033[41m This is 红  \033[0m")
	//fmt.Println("\033[42m This is 绿  \033[0m")
	//fmt.Println("\033[43m This is 黄  \033[0m")
	//fmt.Println("\033[44m This is 蓝  \033[0m")
	//fmt.Println("\033[45m This is 紫  \033[0m")
	//fmt.Println("\033[46m This is 蓝绿\033[0m")
	//fmt.Println("\033[47m This is 灰  \033[0m")
}
