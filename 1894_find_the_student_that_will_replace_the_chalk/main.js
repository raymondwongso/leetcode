/**
 * @param {number[]} chalk
 * @param {number} k
 * @return {number}
 */
var chalkReplacer = function(chalk, k) {
  let sum = 0
  for (let i = 0; i < chalk.length; i++) {
    sum += chalk[i]
  }

  if (k > sum) {
    k = k % sum
  }

  for (let i = 0; i < chalk.length; i++) {
    if (k >= chalk[i]) {
      k -= chalk[i]
    } else {
      return i
    }
  }

  return 0
};

chalk = [1, 1, 1]
k = 2

console.log(chalkReplacer(chalk, k))
