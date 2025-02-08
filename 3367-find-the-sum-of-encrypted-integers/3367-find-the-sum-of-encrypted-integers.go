func sumOfEncryptedInt(nums []int) int {
   totalSum := 0
   for _, num := range nums {
    totalSum += encrypt(num)
   }
   return totalSum
}

func encrypt(num int) int {
    maxdigit := 0
    temp := num
    for temp > 0 {
    digit := temp % 10 
    if digit > maxdigit {
        maxdigit = digit
    }
    temp /= 10
}
newNum := 0
multiplier :=1
temp = num
for temp > 0 {
    newNum += maxdigit * multiplier
    multiplier *= 10
    temp /= 10
}
return newNum
}