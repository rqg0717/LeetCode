#include <stdio.h>
#include <stdlib.h>

int singleNumber(int* nums, int numsSize){
    int result = 0;
    int i = 0;
    for (i = 0; i < numsSize; i++) {
        result = result ^ nums[i];
    }
    return result;
}

int main() {
    int nums[] = {3, 1, 2, 2, 1};
    printf("Output: %d.", singleNumber(nums, 5));
    return 0;
}