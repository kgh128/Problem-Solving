/*
1. 비트 연산으로 2의 거듭제곱 구하기
-> 2 << i는 2를 (i+1)번 제곱한 것과 같다. (i >= 0)
ex) 2:       0010
    2 << 1:  0100  => 4
-> 단, 이 방법으로는 2의 0 제곱, 즉 1을 구할 수는 없으므로 1은 if문으로 따로 검사해야 한다.

2. 슬라이스의 뒤에 0을 추가하기
- make를 이용하여 원하는 길이의 슬라이스를 만든다.
=> 이 슬라이스는 0으로 초기화되어 있음.
- append를 이용하여 기존 슬라이스에 새로 만든 슬라이스를 붙인다.
=> 다른 값이 아닌 0을 추가하는 것이기 때문에 가능한 방법
*/

func solution(arr []int) []int {
    addLength := 0

    // 1. 2의 0제곱 -> 1
    if (len(arr) == 1) {
        return arr
    }

    // 2. 2의 1제곱부터 계산
    for i := 0; ; i++ {
        if (2 << i) >= len(arr) {
            addLength = (2 << i) - len(arr)
            break
        }
    }

    // 3. 0 추가
    return append(arr, make([]int, addLength)...)
}


/*
3. len(arr) & (len(arr)-1)이 0이면 len(arr)은 2의 거듭제곱
- 위의 시프트 연산자를 보면 알 수 있듯이, 2의 거듭제곱들은 비트로 표현하면 1이 하나뿐이다.
  (2(0010)에서 비트를 왼쪽으로 이동만 시키면 2의 거듭제곱이므로)
- 따라서 (2의 거듭제곱)-1은 원래 1이었던 비트는 0이 되고, 그 아래 비트들은 모두 1이 된다.
ex) 16: 0001 0000
    15: 0000 1111
=> 결국 둘을 & 연산하면 모든 비트가 0이 되므로 결과가 0이 나온다.
*/

func solution(arr []int) []int {
    for len(arr) & (len(arr)-1) != 0 {
        arr = append(arr, 0)
    }
    return arr
}
