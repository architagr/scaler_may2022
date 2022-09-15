package count_of_pairs_with_the_given_sum

func Solve(A []int , B int )  (int) {
    
    count := 0
    i,j:= 0, len(A)-1
    for i<j{
        sum := A[i]+A[j]
        if  sum == B{
            count++
            i++
            j--
        }else if sum < B{
            i++
        }else{
            j--
        }
        
    }
    return count
}
