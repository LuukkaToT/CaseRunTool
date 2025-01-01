package casetable

import (
	"fmt"
	//"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	//"os"
)

type CaseTableInfo struct {
	CaseName string
	CasePath string
	IP       string
}

// ReadCaseTable 读取用例执行表
func ReadCaseTable(localCaseTablePath string) ([]CaseTableInfo, string) {
	//打开用例执行表
	localCaseTable, err := excelize.OpenFile(localCaseTablePath)
	if err != nil {
		return nil, "打开用例执行表失败"
	}
	defer func(localCaseTable *excelize.File) {
		err = localCaseTable.Close()
		if err != nil {
			return
		}
	}(localCaseTable)
	//获取用例执行表第一页
	sheet1 := localCaseTable.GetSheetName(0)
	if sheet1 == "" {
		return nil, "打开用例执行表第一页失败！"
	}
	//定义数组存储用例执行表的信息
	var allCaseInfo []CaseTableInfo
	rows, err := localCaseTable.GetRows(sheet1)
	if err != nil {
		return nil, "获取用例执行表信息失败！"
	}
	//提取用例执行表信息，存入allCaseInfo数组中,第一行除外
	for _, row := range rows[1:] {
		allCaseInfo = append(allCaseInfo, CaseTableInfo{
			CaseName: row[0],
			CasePath: row[1],
			IP:       row[2],
		})
	}
	return allCaseInfo, "读取用例执行表成功"
}

// 按照执行机ip对这部分ip进行排序
func sortCaseTable(allCaseInfo []CaseTableInfo) {
	sort.Slice(allCaseInfo, func(i, j int) bool {
		return allCaseInfo[i].CaseName > allCaseInfo[j].CaseName
	})
}

// 写入执行机用例执行表
func writeToCaseTable(c *CaseTableInfo, targetFile *excelize.File) string {
	sheetName := "sheet1"
	rowIndex := 2
	err := targetFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), c.CaseName)
	if err != nil {
		return fmt.Sprintf("写入用IP:%s例执行表A%d失败", c.IP, rowIndex)
	}
	err = targetFile.SetCellValue(sheetName, fmt.Sprintf("H%d", rowIndex), c.CasePath)
	if err != nil {
		return fmt.Sprintf("写入用IP:%s例执行表H%d失败", c.IP, rowIndex)
	}
	err = targetFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), ":YES")
	if err != nil {
		return fmt.Sprintf("写入用IP:%s例执行表C%d失败", c.IP, rowIndex)
	}
	return "写入成功"
}

// Process 执行用例执行表写入过程
func Process(localCaseTablePath string) string {
	//读取本地用例执行表信息
	allCaseInfo, errorMessage := ReadCaseTable(localCaseTablePath)
	if errorMessage != "读取用例执行表成功" {
		return errorMessage
	}
	//根据IP对用例数据进行排序
	sortCaseTable(allCaseInfo)
	for i := 0; i < len(allCaseInfo); i++ {
		c := allCaseInfo[i]
		targetFile, err := excelize.OpenFile(c.IP)
		if err != nil {
			return fmt.Sprintf("打开IP:%s用例执行表失败", c.IP)
		}
		message := writeToCaseTable(&c, targetFile)
		if message != "写入成功" {
			return message
		}
		for j := i + 1; j < len(allCaseInfo) && allCaseInfo[j].IP == c.IP; j++ {
			i = j
			message = writeToCaseTable(&allCaseInfo[j], targetFile)
			if message != "写入成功" {
				return message
			}
		}
		err = targetFile.SetCellValue("sheet1", fmt.Sprintf("A%d", 2), "end")
		if err != nil {
			return fmt.Sprintf("写入end失败")
		}

		err = targetFile.SaveAs(c.IP)
		if err != nil {
			return fmt.Sprintf("保存IP:%s用例执行表失败", c.IP)
		}

	}
	return "用例执行表写入成功"
}
