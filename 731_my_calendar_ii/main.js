var MyCalendarTwo = /** @class */ (function () {
    function MyCalendarTwo() {
        this.duplicateArr = [];
        this.arr = [];
    }
    MyCalendarTwo.prototype.book = function (start, end) {
        for (var i = 0; i < this.duplicateArr.length; i++) {
            if (start < this.duplicateArr[i][1] && this.duplicateArr[i][0] < end) {
                return false;
            }
        }
        for (var i = 0; i < this.arr.length; i++) {
            if (start < this.arr[i][1] && this.arr[i][0] < end) {
                this.duplicateArr.push([Math.max(start, this.arr[i][0]), Math.min(end, this.arr[i][1])]);
            }
        }
        this.arr.push([start, end]);
        return true;
    };
    return MyCalendarTwo;
}());
var obj = new MyCalendarTwo();
console.log(obj.book(10, 20));
console.log(obj.book(50, 60));
console.log(obj.book(10, 40));
console.log(obj.book(5, 15));
console.log(obj.book(5, 10));
console.log(obj.book(25, 55));
// console.log(obj.book(20, 29))
// console.log(obj.book(13, 22))
// console.log(obj.book(44, 50))
// console.log(obj.book(1, 7))
// console.log(obj.book(2, 10))
// console.log(obj.book(14, 20))
// console.log(obj.book(19, 25))
// console.log(obj.book(36, 42))
// console.log(obj.book(45, 50))
// console.log(obj.book(47, 50))
// console.log(obj.book(39, 45))
// console.log(obj.book(44, 50))
// console.log(obj.book(16, 25))
// console.log(obj.book(45, 50))
// console.log(obj.book(45, 50))
// console.log(obj.book(12, 20))
// console.log(obj.book(21, 29))
// console.log(obj.book(11, 20))
// console.log(obj.book(12, 17))
// console.log(obj.book(34, 40))
// console.log(obj.book(10, 18))
// console.log(obj.book(38, 44))
// console.log(obj.book(23, 32))
// console.log(obj.book(38, 44))
// console.log(obj.book(15, 20))
// console.log(obj.book(27, 33))
// console.log(obj.book(34, 42))
// console.log(obj.book(44, 50))
// console.log(obj.book(35, 40))
// console.log(obj.book(24, 31))
