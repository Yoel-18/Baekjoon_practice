package silver

import "fmt"

//스택에 사용 할 노드
type Node struct {
	Value int
}

//값 출력함수
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

//스택
func NewStack() *Stack {
	return &Stack{}
}

//스택 구조 본체 선언
type Stack struct {
	nodes []*Node
	count int
}

//데이터 삽입
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

//데이터 빼기
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
func main() {
	s := NewStack()

	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})
	s.Push(&Node{4})
	s.Push(&Node{5})

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
