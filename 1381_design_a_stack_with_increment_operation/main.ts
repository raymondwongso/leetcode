class CustomStack {
  arr: Array<number>
  index: number
  maxSize: number

  constructor(maxSize: number) {
    this.arr = new Array()
    this.index = 0
    this.maxSize = maxSize
  }

  push(x: number): void {
    if (this.index == this.maxSize) {
      return
    }

    this.arr[this.index] = x
    this.index += 1
  }

  pop(): number {
    if (this.index == 0) {
      return -1
    }

    let res = this.arr[this.index-1]
    this.arr[this.index-1] = -1
    this.index -= 1

    return res
  }

  increment(k: number, val: number): void {
    for (let i = 0; i < k; i++) {
      if (i >= this.index) {
        break
      }

      this.arr[i] += val
    }
  }
}

var obj = new CustomStack(3)
obj.push(1)
obj.push(2)
obj.pop()
obj.push(2)
obj.push(3)
obj.push(4)
obj.increment(5, 100)
obj.increment(2, 100)
obj.pop()
obj.pop()
obj.pop()
obj.pop()
