#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int mySqrt(int x) {
    double guess = (double)x;
    double x0 = (double)x, xi;
	if (x == 0) {
		return 0;
	}

	for (;;) {
		// 2xix0-(x0^2+guess) = 0
		// xi = 1/2(x0+guess/x0)
		xi = 0.5 * (x0 + guess/x0);
		if ((int)x0==(int)xi) {
			break;
		}
		x0 = xi;
	}
	return (int)x0;
}

int bulbSwitch(int n){
    return mySqrt(n);
}

int main() {
    printf("Output: %d.",bulbSwitch(3));
    return 0;
}