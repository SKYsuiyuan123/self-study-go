package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"todo-cli/storage"
)

// parseID 解析命令里的 ID 参数
func parseID(args []string) (int, error) {
	if len(args) < 3 {
		return 0, errors.New("请提供 ID, 例如: todo done 1")
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, fmt.Errorf("ID 必须是数字: %s", args[2])
	}

	return id, nil
}

func main() {
	if len(os.Args) < 2 {
		storage.Usage()
		return
	}

	storeVal, err := storage.Load()
	if err != nil {
		fmt.Println("读取数据失败: ", err)
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("用法: todo add <内容>")
			return
		}
		title := os.Args[2]
		t := storeVal.Add(title)
		if err := storeVal.Save(); err != nil {
			fmt.Println("保存失败: ", err)
			return
		}

		fmt.Printf("已添加 #%d: %s\n", t.Id, t.Title)

	case "list", "ls":
		storeVal.PrintList()

	case "done":
		id, err := parseID(os.Args)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := storeVal.Done(id); err != nil {
			fmt.Println(err)
			return
		}
		if err := storeVal.Save(); err != nil {
			fmt.Println("保存失败: ", err)
			return
		}

		fmt.Printf("已完成 #%d \n", id)

	case "rm", "remove", "del":
		id, err := parseID(os.Args)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := storeVal.Remove(id); err != nil {
			fmt.Println(err)
			return
		}
		if err := storeVal.Save(); err != nil {
			fmt.Println("保存失败: ", err)
			return
		}

		fmt.Printf("已删除 %#d \n", id)

	case "help", "-h", "--help":
		storage.Usage()

	default:
		fmt.Printf("未知命令: %s\n\n", cmd)
		storage.Usage()
	}
}
