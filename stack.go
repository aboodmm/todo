package stack

type Stack struct {
	maxsize int
	top     int
	array   [maxsize]string
}

func Push(n string) {
	maxsize += 1
	array[maxsize] = n
}

func Pop() {

}
