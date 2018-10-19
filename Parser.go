package main

import (
	"io/ioutil"
	"HuffmanCoding/Nodes"
	"fmt"
)

func contains(nodeList []Nodes.Node, node Nodes.Node) bool {
	for _, nodes := range nodeList {
		if node.Key == nodes.Key {
			return true
		}
	}
	return false
}

func getNode(nodeList []Nodes.Node, node Nodes.Node) *Nodes.Node{
	for index, nodes := range nodeList {
		if node.Key == nodes.Key {
			return &nodeList[index]
		}
	}
	return &node
}

func main(){
	data, _ := ioutil.ReadFile("coords.json")

	var nodeList []Nodes.Node

	for _,byt := range data {
		n := &Nodes.Node{
			Key:string(byt),
			Value: 1,
		}
		if contains(nodeList, *n){
			curentNode := getNode(nodeList, *n)
			curentNode.Value = curentNode.Value + 1
		}else {
			nodeList = append(nodeList, *n)
		}

	}

	for i := 0; i < len(nodeList); i++  {
		fmt.Println([]byte(nodeList[i].Key), nodeList[i].Value)

	}


	fmt.Print(string(34))

}