package main

func main() {

	//1.1、显示声明变量类型
	//var userName string = "小袁同学"
	//UserName:="哈哈哈"

	//1.2、类型推断声明（Type Inference）
	//var userName = "小袁同学"

	//1.3、简短变量声明（只能在方法体内声明，使用，局部变量）
	//userName := "小袁同学"
	//println(userName,UserName)

	//2.1 集合类型赋值
	//var (
	//	userName = "小明"
	//	userAge  = 20
	//	isOk     = true
	//)

	//2.2 直接赋值
	//var userName, userAge, isOk = "小明", 20, true

	//2.3 简短多变量赋值
	//userName, userAge, isOk := "小明", 20, true
	//println(userName, userAge, isOk)

	//若只想获得 userName ，则函数调用语句可以用如下方式编写：
	_, _, userName := GetUserInfo()

	println(userName)
}

func GetUserInfo() (userAge int, isOk bool, userName string) {
	return 20, true, "追逐时光者"
}
