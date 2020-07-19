int coinChange(int* coins, int coinsSize, int amount)
{
	int i = 0, ret = -1;
    int * opt = (int *)malloc(sizeof(int) * (amount + 1));
	opt[0] = 0;
	for (i = 1; i <= amount; i++) {
        opt[i] = amount + 1;
		for (int j = 0; j < coinsSize; j++) {
			if (i >= coins[j]) {
				opt[i] = min(opt[i - coins[j]] + 1, opt[i]);
			}
		}
	}
    ret = opt[amount] > amount ? -1 : opt[amount]
    free(opt);
	return ret;
}