#define MAX(x, y) (x > y ? x:y)
int maxProfit(int k, int* prices, int pricesSize) {
    int i, j = 0;
    int profit = 0;
    if( pricesSize == 0 || k < 1 ) return 0;
    if( k > pricesSize/2 ) {
        for( i = 1; i < pricesSize; i ++ ) {
            profit += MAX(prices[i] - prices[i-1],0);
        }
        return profit;
    }
    int dp[k+1][2];
    memset(dp, 0, sizeof(int) * (k+1) * 2);
    // init first day
	for( i = 0; i <= k; i++ ) {
		dp[i][1] = -prices[0];
	}
    for( i = 0; i < pricesSize; i++ ) {
        for( j = 1; j <= k; j++ ) {
            // i day, do not have any stock
			// i-1 day did not have or sold today
			dp[j][0] = MAX(dp[j][0], dp[j][1]+prices[i]);
			// i day, have stock
			// i-1 day had stock or bought today
			dp[j][1] = MAX(dp[j][1], dp[j-1][0]-prices[i]);
        }
    }
    return dp[k][0];
}