// This problem is so stupid and weird -- not recommended at all
type MyStack struct {
    q []int
}

func Constructor() MyStack {
    return MyStack{q: make([]int, 0)}
}

// Push appends and rotates older elements so the newest item is at the front
func (this *MyStack) Push(x int) {
    this.q = append(this.q, x)
    for i := 0; i < len(this.q)-1; i++ {
        this.q = append(this.q, this.q[0])
        this.q = this.q[1:]
    }
}

func (this *MyStack) Pop() int {
    val := this.q[0]
    this.q = this.q[1:]
    return val
}

func (this *MyStack) Top() int {
    return this.q[0]
}

func (this *MyStack) Empty() bool {
    return len(this.q) == 0
}
