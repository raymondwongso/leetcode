function checkInclusion(s1: string, s2: string): boolean {
  if ( s1.length > s2.length ) {
    return false
  }
  
  let s1arr = new Array<number>(26)
  s1arr.fill(0)
  let s2arr = new Array<number>(26)
  s2arr.fill(0)
  let offset = 97

  for (let i = 0; i < s1.length; i++) {
    s1arr[s1.charCodeAt(i)-offset]++
    s2arr[s2.charCodeAt(i)-offset]++
  }

  if (JSON.stringify(s1arr) == JSON.stringify(s2arr)) {
    return true
  }

  for (let i = s1.length; i < s2.length; i++) {
    let j = i - s1.length

    s2arr[s2.charCodeAt(j)-offset]--
    s2arr[s2.charCodeAt(i)-offset]++

    if (JSON.stringify(s1arr) == JSON.stringify(s2arr)) {
      return true
    }
  }

  return false
};

checkInclusion("abc", "aadbcasext")
