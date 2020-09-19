#include <stdio.h>

int missingNumber(int* nums, int numsSize){
	int result = 0;
	int i = 0;
	for (i = 0; i < numsSize; i++) {
		result ^= nums[i] ^ i;
	}
	return result ^ numsSize;
}

int main() {
    int nums[] = {9, 6, 4, 2, 3, 5, 7, 0, 1};
    printf("Output: %d.", missingNumber(nums, 9));
    return 0;
}