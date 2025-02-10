func flipAndInvertImage(image [][]int) [][]int {
   n := len(image)

   for i:=0; i<n; i++ {
    for j:=0; j<n/2; j++ {
        image[i][j], image[i][n-1-j] = image[i][n-1-j], image[i][j]
    }
   }
   for i:=0; i<n; i++ {
    for j:=0; j<n; j++ {
        if image[i][j] == 0 {
            image[i][j] =1
        } else {
            image[i][j] =0
        }
    }
   }
   return image
}