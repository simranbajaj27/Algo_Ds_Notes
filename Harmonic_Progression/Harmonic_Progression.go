/*
        HARMONIC PROGRESSION

        Any sequence of elements where the difference between any
        two consecutive elements is equal is termed to be in A.P.
        Harmonic progression is a sequence of quantities whose
        reciprocals are in arithmetical progression.
        The nth term of an H.P. is defined as 1/(a + (n - 1)*d)
        where a is the first element of the A.P.
        d is the common difference of the A.P.
*/

package main

import "fmt"

func main() {
    fmt.Print("Enter the First Term of A.P.")
    var a int
    fmt.Scan(&a)
    fmt.Print("Enter the common difference")
    var d int
    fmt.Scan(&d)
    fmt.Print("Enter N (The index of term to find)")
    var n int
    fmt.Scan(&n)
    fmt.Print("The term at index ", n , " of H.P. is 1/", (a + (n - 1) * d))
}

/*
    INPUT : a = 5
            d = 2
            n = 7
    OUTPUT : The term at index 7 of H.P. is 1/17
    VERIFICATION : The A.P. would be 5,7,9,11,13,15,17...
                   so H.P. would be 1/5, 1/7, 1/9, 1/11, 1/13, 1/17....
*/
