package dynamic_programming

/**
疯子坐飞机找座问题

数学推导后得出，当n>=3时， f(n) = f(n-1)
由于已知 f(1)=1 和 f(2)=0.5，因此当 n≥3 时，根据 f(n) = f(n-1)可知，
f(n)={ 
1.0,n=1
0.5,n≥2
​}
*/
func nthPersonGetsNthSeat(n int) float64 {
    if n == 1{
        return 1
    }
    
    return 0.5
}
