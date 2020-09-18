#include <stdio.h>
#include <stdlib.h>
#include "uthash.h"

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    int i = 0, j = 0;
    int *results = (int *)malloc(sizeof(int) * 2);
    for (i = 0; i < numsSize - 1; i++)
    {
        for (j = i + 1; j < numsSize; j++)
        {
            if (nums[i] + nums[j] == target)
            {
                results[0] = i;
                results[1] = j;
                *returnSize = 2;
                return results;
            }
        }
    }
    return results;
}


struct HASH_TABLE {
    int key;
    int value;
    UT_hash_handle hh;
};

typedef struct HASH_TABLE *TABLE;

int* twoSum1(int* nums, int numsSize, int target, int* returnSize) {
    int i = 0, key = 0;
    TABLE head = NULL, add = NULL, out = NULL;
    int *results = NULL;

    for( i = 0; i < numsSize; i++ ) {
        key = target - nums[i];
        HASH_FIND_INT(head, &key, out);
        if( out != NULL ) {
            results = (int*)malloc(sizeof(int) * 2);
            *returnSize = 2;
            results[1] = i;
            results[0] = out->value;
            return results;
        }
        add = (TABLE)malloc(sizeof(struct HASH_TABLE));
        add->key = nums[i];
        add->value = i;
        HASH_ADD_INT(head, key, add);
    }

    return results;
}

int main() {
    int nums[4]={2, 7, 11, 15};
    int size = 0;
    int *results = twoSum1(nums, 4, 9, &size);
    if (size > 0) {
        printf("Output: [%d, %d].", results[0], results[1]);
    } else {
        printf("Output: NO RESULT.");
    }
    free(results);
    results = NULL;
    return 0;
}