var CustomStack = /** @class */ (function () {
    function CustomStack(maxSize) {
        this.arr = new Array();
        this.index = 0;
        this.maxSize = maxSize;
    }
    CustomStack.prototype.push = function (x) {
        if (this.index == this.maxSize) {
            return;
        }
        this.arr[this.index] = x;
        this.index += 1;
    };
    CustomStack.prototype.pop = function () {
        if (this.index == 0) {
            return -1;
        }
        var res = this.arr[this.index];
        this.arr[this.index] = -1;
        this.index -= 1;
        return res;
    };
    CustomStack.prototype.increment = function (k, val) {
        for (var i = 0; i < k; i++) {
            if (i > this.index) {
                break;
            }
            this.arr[i] += val;
        }
    };
    return CustomStack;
}());
var obj = new CustomStack(3);
obj.push(1);
obj.push(2);
obj.pop();
obj.push(2);
obj.push(3);
obj.push(4);
obj.increment(5, 100);
obj.increment(2, 100);
obj.pop();
obj.pop();
obj.pop();
obj.pop();
