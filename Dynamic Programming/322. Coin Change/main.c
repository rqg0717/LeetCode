#include <stdio.h>
#include <stdlib.h>

int getMIN(int x, int y)
{
	int ret = x < y ? x : y;
	return ret;
}

int coinChange(int *coins, int coinsSize, int amount)
{
	int i = 0, result = -1;
	int *dp = (int *)malloc(sizeof(int) * (amount + 1));
	dp[0] = 0;
	for (i = 1; i <= amount; i++)
	{
		dp[i] = amount + 1;
		for (int j = 0; j < coinsSize; j++)
		{
			if (i >= coins[j])
			{
				dp[i] = getMIN(dp[i - coins[j]] + 1, dp[i]);
			}
		}
	}
	result = dp[amount] > amount ? -1 : dp[amount];
	free(dp);
	return result;
}

int main()
{
	int coins[3] = {1, 2, 5};
	int coinsSize = 3;
	int amount = 11;
	int result = coinChange(coins, coinsSize, amount);
	if (result > 0)
	{
		printf("Output: [%d].", result);
	}
	else
	{
		printf("Output: NO RESULT.");
	}
	return 0;
}