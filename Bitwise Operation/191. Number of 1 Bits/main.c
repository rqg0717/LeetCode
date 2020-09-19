#include <stdio.h>
#include <stdint.h>

int hammingWeight(uint32_t n) {
    int count = 0;
	while(n > 0) {
		n = n & (n - 1);
		count++;
	}
	return count;
}

int main() {
    uint32_t n = 0b11111111111111111111111111111101;
    printf("Output: %d.", hammingWeight(n));
    return 0;
}