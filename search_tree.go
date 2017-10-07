package main

import (
	"fmt"
	"log"
	"io"
	"os"
)

type RootNode struct {
	Root   *Node
	Logger *log.Logger
	LoggerConfDefault
}
type LoggerConfDefault struct {
	Flags  int
	Prefix string
	Out    io.Writer
}
type Node struct {
	ID    int
	PID   int
	Stock []*Node
	Obj   Comment
}
type Comment struct {
	ID   int
	PID  int
	Text string
}

func NewRootNode(logout io.Writer, prefix string) *RootNode {
	return &RootNode{
		Root:              &Node{ID: 0, PID: 0},
		Logger:            log.New(logout, prefix, log.Ldate|log.Ltime|log.Llongfile),
		LoggerConfDefault: LoggerConfDefault{
			Prefix:prefix,
			Flags:log.Ldate|log.Ltime|log.Llongfile,
			Out:logout,
		},
	}
}
func (r *RootNode) add(node *Node) {
	//проверка наличия рутовой ноды
	if node == nil {
		r.Logger.Printf("[ошибка] Нельзя добавить `nil`\n")
		return
	}
	//если нет корневой ноды, то связываем добавляему ноду как корневую
	if r.Root == nil {
		r.Root = node
		r.Logger.Printf("ДОбавлено корневая нода\n")
		return
	}
	//проверка корневой ноды на ID == node.PID если да то добавляем
	if node.PID == r.Root.ID {
		r.Root.Stock = append(r.Root.Stock, node)
		r.Logger.Printf("Нода успешно добавлена к корневой ноде\n")
		return
	}
	//PID добавляемой ноды не равен ID корневой, значит ищу в стоке рекурсивно при наличии элементов
	result := &Node{}
	if len(r.Root.Stock) > 0 {
		for _, x := range r.Root.Stock {
			result = r.search(node.PID, x)
			if result != nil {
				result.Stock = append(result.Stock, node)
				r.Logger.Printf("Новая нода `%v` успешно добавлена к своему `родителю`: `%v`\n", node, result)
				return
			}
		}
		r.Logger.Printf("[ошибка] ноды с `%d` в дереве не найдено\n", node.PID)
		return
	}
}
func (r *RootNode) search(id int, node *Node) *Node {
	result := &Node{}
	//если node == nil, ищем начиная с корневой ноды + списки что в Root.Stock[] при условии Root.Stock[] > 0
	if node == nil {
		//Root нода пустая, там ничего нет
		if r.Root == nil {
			r.Logger.Printf("Дерево пустое, нечего искать\n")
			return nil
		}
		//проверка корневого на ID
		if r.Root.ID == id {
			return r.Root
		}
		//рекурсивно ищу искомый объект с ID начиная с корневого
		if len(r.Root.Stock) > 0 {
			for _, x := range r.Root.Stock {
				result := r.search(id, x)
				if result != nil {
					return result
				}
			}
			r.Logger.Printf("[ошибка] ноды с `%d` в дереве не найдено\n", id)
			return nil
		}
	}
	//если node != nil
	if node != nil {
		if node.ID == id {
			return node
		}
		if len(node.Stock) > 0 {
			for _, x := range node.Stock {
				result = r.search(id, x)
				if result != nil {
					return result
				}
			}
		}
	}
	return nil
}
func (r *RootNode) show() {
	if r.Root == nil {
		r.Logger.Printf("[ошибка] дерева нет, нечего показывать\n")
		return
	}
	for _, x := range r.Root.Stock {
		r.Logger.Printf("`%v`\n", x)
	}
}
func main() {
	n := NewRootNode(os.Stdout, "[COMMENTSTREE]")

	commenstStock := []Comment{
		{1, 0, "ROOT 1"},
		{2, 0, "ROOT 2"},
		{3, 1, "Simpletext"},
		{4, 1, "Simpletext"},
		{5, 2, "Simpletext"},
		{6, 2, "Simpletext"},
		{7, 1, "Simpletext"},
		{8, 1, "Simpletext"},
		{9, 4, "Simpletext"},
		{10, 9, "Simpletext"},
		{11, 9, "Simpletext"},
		{12, 9, "Simpletext"},
	}
	for _, c := range commenstStock {
		nn := &Node{
			ID:  c.ID,
			PID: c.PID,
			Obj: c,
		}
		n.add(nn)
	}
	n.showtree(n.Root, os.Stdout)
	n.show()
}
func (r *RootNode) showtree(node *Node, out io.Writer) {
	if node == nil {
		r.Logger.Printf("дерево пустое, нечего показывать\n")
		return
	}

	fmt.Fprint(out, "<ul>\n")
	for _, x := range node.Stock {
		fmt.Fprintf(out, "<li id='%d_%d'> %s </li>\n", x.ID, x.PID, x.Obj.Text)
		if len(x.Stock) > 0 {
			r.showtree(x, out)
		}
	}
	fmt.Fprint(out, "</ul>\n")

}
