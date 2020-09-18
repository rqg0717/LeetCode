#include <stdio.h>
#include <stdlib.h>


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    int i = 0, j = 0;
    int *result = (int *)malloc(sizeof(int) * 2);
    for (i = 0; i < numsSize - 1; i++)
    {
        for (j = i + 1; j < numsSize; j++)
        {
            if (nums[i] + nums[j] == target)
            {
                result[0] = i;
                result[1] = j;
                *returnSize = 2;
                return result;
            }
        }
    }
    return result;
}

int main() {
    int nums[4]={2, 7, 11, 15};
    int size = 0;
    int *results = twoSum(nums, 4, 9, &size);
    if (size > 0) {
        printf("Output: [%d, %d].", results[0], results[1]);
    } else {
        printf("Output: NO RESULT.");
    }
    free(results);
    results = NULL;
    return 0;
}