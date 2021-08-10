package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	queue := CreateQueue(10)

	queue.Insert("C:\\temp1.txt")
	queue.Insert("C:\\temp2.txt")
	queue.Insert("C:\\temp3.txt")
	queue.Insert("C:\\temp4.txt")
	queue.Insert("C:\\temp5.txt")
	queue.Insert("C:\\temp6.txt")
	queue.Insert("C:\\temp6.txt")
	queue.Insert("C:\\temp7.txt")
	queue.Insert("C:\\temp9.txt")
	queue.Insert("C:\\temp10.txt")
	queue.Insert("C:\\temp11.txt")

	//insert do a "shift up" of all elements
	assert.Equal(t,"C:\\temp2.txt", queue.q[0])
	assert.Equal(t,"C:\\temp11.txt", queue.q[len(queue.q) -1])
}
