package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input := fetch("../input.txt")

	disk := buildDisk(input)
	moveDisks(disk)
	res := calculateChecksum(disk)
	fmt.Println("Actual Result: ", res)

	testInput := `2333133121414131402`

	testDisk := buildDisk(testInput)
	moveDisks(testDisk)
	testRes := calculateChecksum(testDisk)
	fmt.Println("Test Result: ", testRes)
}

func fetch(fileName string) string {
	f, _ := os.Open(fileName)
	defer f.Close()
	data, _ := io.ReadAll(f)
	return string(data)
}

type Node struct {
	Size int
	ID   int // -1 for free
	Next *Node
	Prev *Node
}

func buildDisk(input string) *Node {
	var id int
	dummy := &Node{}
	cur := dummy

	for i, c := range input {
		num := int(c - '0')
		newNode := &Node{Size: num}
		if i%2 == 0 {
			newNode.ID = id
			id++
		} else {
			newNode.ID = -1
		}
		cur.Next = newNode
		newNode.Prev = cur
		cur = newNode
	}

	return dummy.Next
}


func moveDisks(start *Node) {
	right := start
	// Move to the rightmost node
	for right.Next != nil {
		right = right.Next
	}

	// Process from right to left
	for right != start {
		// Skip free nodes
		for right.ID == -1 && right.Prev != nil {
			right = right.Prev
		}

		if right.ID == -1 {
			break
		}

		neededSpace := right.Size
		left := start

		// Find the first free space that can fit the block
		for left != right && (left.ID != -1 || left.Size < neededSpace) {
			left = left.Next
		}

		// No suitable space found
		if left == right {
			right = right.Prev
			continue
		}

		// Replace the space
		if neededSpace == left.Size {
			left.ID = right.ID
			right.ID = -1
		} else {
			// Create a new free node with leftover space
			newNode := &Node{Size: left.Size - neededSpace, ID: -1}
			newNode.Next = left.Next
			if left.Next != nil {
				left.Next.Prev = newNode
			}
			left.Next = newNode
			newNode.Prev = left

			left.Size = neededSpace
			left.ID = right.ID
			right.ID = -1
		}

		right = right.Prev
	}
}

func calculateChecksum(start *Node) int {
	checksum := 0
	index := 0
	cur := start

	for cur != nil {
		if cur.ID == -1 {
			index += cur.Size
		} else {
			for i := 0; i < cur.Size; i++ {
				checksum += cur.ID * (index + i)
			}
			index += cur.Size
		}
		cur = cur.Next
	}

	return checksum
}
