#include <stdio.h>
#include <stdlib.h>

int singleNumber(int* nums, int numsSize){
    int a = 0, b = 0, c = 0;
    int i = 0;
	for (i = 0; i < numsSize; i++) {
		a = (b & ~nums[i]) | (c & nums[i]);
		c = (~b & ~c & nums[i]) | (c & ~nums[i]);
		b = a;
	}
	return c;
}

int main() {
    int nums[] = {0, 1, 0, 1, 0, 1, 99};
    printf("Output: %d.", singleNumber(nums, 7));
    return 0;
}