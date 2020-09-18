#include <iostream>
#include <vector>
#include <map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) 
    {
        map<int,int> mNums;
        vector<int> results(2, -1);
        for(int i = 0; i < nums.size(); i++)
        {
            if(mNums.count(target-nums[i]) > 0)
            {
                results[0] = mNums[target-nums[i]];
                results[1] = i;
                break;
            }
            mNums[nums[i]] = i;
        }
        return results;
    };
};

int main()
{
    Solution s;
    int nums[] = {2, 7, 11, 15};
    vector<int> vNums (nums, nums + sizeof(nums) / sizeof(int) );
    vector<int> results = s.twoSum(vNums, 9);
    printf("Output: [%d, %d].", results[0], results[1]);
    return 0;
}