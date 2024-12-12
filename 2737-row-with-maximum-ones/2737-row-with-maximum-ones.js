/**
 * @param {number[][]} mat
 * @return {number[]}
 */
var rowAndMaximumOnes = function(mat) {
    let maxOnes = 0;
    let arr = []
    for(let i=0; i<mat.length; i++){
        let temp = 0;
        for(let j = 0; j < mat[i].length; j++){
            if(mat[i][j] == 1){
                temp++
                console.log(temp)
            }
        }
        if(maxOnes < temp){
            maxOnes = temp;
           arr[0] = i;
           arr[1] = maxOnes
        }
    }
    if(maxOnes == 0){
        let y = [0,0]
        return y
    }
    return arr
};