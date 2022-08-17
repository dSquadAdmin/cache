package cache

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Key            string
	Value          interface{}
	next, previous *Node
}

func NewNode() *Node {
	node := new(Node)
	return node
}

func (node *Node) Clone() *Node {
	newNode := new(Node)
	newNode.Key = node.Key
	newNode.Value = node.Value
	newNode.next = nil
	newNode.previous = nil
	return newNode
}

func (node *Node) SetNext(next *Node) *Node {
	node.next = next
	return node
}

func (node *Node) Next(next *Node) *Node {
	return node.next
}

func (node *Node) SetPrevious(previous *Node) *Node {
	node.previous = previous
	return node
}

func (node *Node) Previous() *Node {
	return node.previous
}

func (node *Node) Set(key string, value interface{}) *Node {
	node.Key = key
	node.Value = value
	return node
}

func (node *Node) IsHead() bool {
	val, ok := node.Value.(bool)
	return node.Key == "head" && ok && val
}

func (node *Node) IsTail() bool {
	val, ok := node.Value.(bool)
	return node.Key == "tail" && ok && val
}

func (node *Node) Node() string {
	data, _ := json.Marshal(node.Value)
	return fmt.Sprintf("{ \"%s\": \"%s\" }", node.Key, string(data))
}
