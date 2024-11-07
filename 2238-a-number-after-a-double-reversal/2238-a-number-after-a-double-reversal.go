func isSameAfterReversals(num int) bool {
   reversed := 0
   if reversed == num {
    return true
   }
   if num % 10 == 0 {
    return false
   }
   return true
}