package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"

	"todo-cli/util"
)

const dataFile = "./task.json"

type Store struct {
	List   []util.Task `json:"list"`
	NextID int         `json:"next_id"`
}

// Load 从文件新增
func Load() (*Store, error) {
	s := &Store{
		NextID: 1,
	}

	data, err := os.ReadFile(dataFile)

	if errors.Is(err, os.ErrNotExist) {
		return s, nil // 文件不存在，返回空 store
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, s); err != nil {
		return nil, err
	}
	if s.NextID == 0 {
		s.NextID = 1
	}

	return s, nil
}

// Save 写回文件
func (s *Store) Save() error {
	data, err := json.MarshalIndent(s, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}

// Add 新增
func (s *Store) Add(title string) util.Task {
	t := util.Task{
		Id:        s.NextID,
		Title:     title,
		CreatedAt: time.Now(),
	}

	s.NextID++
	s.List = append(s.List, t)
	return t
}

// findIndex 按 ID 找下标
func (s *Store) findIndex(id int) int {
	for i, t := range s.List {
		if t.Id == id {
			return i
		}
	}

	return -1
}

// Done 标记完成
func (s *Store) Done(id int) error {
	i := s.findIndex(id)
	if i < 0 {
		return fmt.Errorf("找不到 ID 为 %d 的代办", id)
	}

	s.List[i].Done = true
	s.List[i].CompletedAt = time.Now()
	return nil
}

// Remove 删除
func (s *Store) Remove(id int) error {
	i := s.findIndex(id)
	if i < 0 {
		return fmt.Errorf("找不到 ID 为 %d 的代办", id)
	}

	s.List = append(s.List[:i], s.List[i+1:]...)
	return nil
}

// List 打印
func (s *Store) PrintList() {
	if len(s.List) == 0 {
		// fmt.Println("暂无待办，用 `add 内容` 添加一条吧")
		color.Yellow("暂无待办，用 `add 内容` 添加一条吧")
		return
	}

	// fmt.Printf("%-4s %-6s %s\n", "Id", "状态", "内容")
	// fmt.Println("---------------------------------------")

	header := color.New(color.Bold, color.FgCyan)
	header.Printf("%-4s %-6s %s\n", "ID", "状态", "内容")
	fmt.Println("--------------------------------")

	for _, t := range s.List {
		// status := "[ ]"
		if t.Done {
			// status = "[x]"

			gray := color.New(color.FgHiBlack, color.CrossedOut)
			gray.Printf("%-4d [x]    %s\n", t.Id, t.Title)
		} else {
			color.Green("%-4d [ ]    %s\n", t.Id, t.Title)
		}

		// fmt.Printf("%-4d %-6s %s\n", t.Id, status, t.Title)
	}
}

// usage 帮助
func Usage() {
	fmt.Println(`TODO CLI 使用说明：

  add <内容>     添加一条待办
  list           列出所有待办
  done <ID>      标记完成
  rm <ID>        删除一条
  help           显示帮助

示例：
  add 学习 Go 的 select
  list
  done 1
  rm 2`)
}
