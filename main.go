package main

import (
	"CaseRunTool/casetable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"os"
	//"path/filepath"
)

func main() {
	// 创建一个应用
	myApp := app.New()

	// 创建一个窗口
	myWindow := myApp.NewWindow("SSH Command Executor")

	// 输入框：目标 IP 地址
	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("Enter server IP")

	// 输入框：命令
	cmdEntry := widget.NewEntry()
	cmdEntry.SetPlaceHolder("Enter command (e.g., ls -a)")

	// 输出框：显示命令执行结果
	outputLabel := widget.NewLabel("Output will be shown here")

	//获取本地路径，调试代码请注释，然后自定义文件路径

	//exePath, err := os.Executable()
	//if err != nil {
	//	outputLabel.SetText(fmt.Sprintf("Failed to get executable path: %v", err))
	//	return
	//}
	//dirPath := filepath.Dir(exePath)
	//localTablePath := dirPath + "\\用例执行表.xlsx"

	//test
	localTablePath := "D:\\Project\\GoProject\\CaseRunTool\\用例执行表.xlsx"

	// 执行按钮
	executeButton := widget.NewButton("Execute", func() {
		str := casetable.Process(localTablePath)
		outputLabel.SetText(str)
		//ip := ipEntry.Text
		//cmd := cmdEntry.Text
		//if ip == "" || cmd == "" {
		//	outputLabel.SetText("Please enter both IP and command.")
		//	return
		//}
		//
		//// 执行 SSH 连接并获取输出
		//output, err := SSH.Ssh_connect_execute(ip, cmd)
		//if err != nil {
		//	outputLabel.SetText(fmt.Sprintf("Error: %v", err))
		//} else {
		//	outputLabel.SetText(fmt.Sprintf("Output:\n%s", output))
		//}
	})

	// 布局：将所有组件放入垂直布局中
	content := container.NewVBox(
		widget.NewLabel("Command:"),
		executeButton,
		outputLabel,
	)

	// 设置窗口内容并展示
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
