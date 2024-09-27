class MyCalendarTwo {
  duplicateArr: number[][]
  arr: number[][]
  constructor() {
    this.duplicateArr = []
    this.arr = []
  }

  book(start: number, end: number): boolean {
    for (let i = 0; i < this.duplicateArr.length; i++) {
      if (start < this.duplicateArr[i][1] && this.duplicateArr[i][0] < end) {
        return false
      }
    }

    for (let i = 0; i < this.arr.length; i++) {
      if (start < this.arr[i][1] && this.arr[i][0] < end) {
        this.duplicateArr.push([Math.max(start, this.arr[i][0]), Math.min(end, this.arr[i][1])])
      }
    }

    this.arr.push([start, end])

    return true
  }
}

var obj = new MyCalendarTwo()
console.log(obj.book(10, 20))
console.log(obj.book(50, 60))
console.log(obj.book(10, 40))
console.log(obj.book(5, 15))
console.log(obj.book(5, 10))
console.log(obj.book(25, 55))
