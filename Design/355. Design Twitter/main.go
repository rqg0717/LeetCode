package main

import (
	"fmt"
	"sort"
)

// User ...
type User struct {
	id   int
	time int
}

// Twitter ...
type Twitter struct {
	time      int
	followers map[int]map[int]int //user ID, follower ID
	tweets    map[int]map[int]int //user ID, tweet ID, tweet time
}

// Constructor ...
func Constructor() Twitter {
	time := 0
	followers := make(map[int]map[int]int)
	tweets := make(map[int]map[int]int)
	return Twitter{
		time:      time,
		followers: followers,
		tweets:    tweets}
}

// PostTweet posts a new tweet.
func (twitter *Twitter) PostTweet(userID int, tweetID int) {
	tweet, ok := twitter.tweets[userID]
	if !ok {
		tweet = make(map[int]int)
		twitter.tweets[userID] = tweet
	}
	twitter.time++
	tweet[tweetID] = twitter.time
}

// Follow follows a user.
func (twitter *Twitter) Follow(userID int, followeeID int) {
	if userID == followeeID {
		return
	}
	follower, ok := twitter.followers[userID]
	if !ok {
		follower = make(map[int]int)
		twitter.followers[userID] = follower
	}
	follower[followeeID] = 1
}

// Unfollow unfollows a user.
func (twitter *Twitter) Unfollow(userID int, followeeID int) {
	if follower, ok := twitter.followers[userID]; ok {
		delete(follower, followeeID)
	}
}

// GetNewsFeed retrieve the 10 most recent tweet ids in the user's news feed.
func (twitter *Twitter) GetNewsFeed(userID int) []int {
	ids := []int{userID}
	if follower, ok := twitter.followers[userID]; ok {
		for followerID := range follower {
			ids = append(ids, followerID)
		}
	}

	users := make([]*User, 0)
	for _, id := range ids {
		if tweet, ok := twitter.tweets[id]; ok {
			for uID, uTime := range tweet {
				users = append(users, &User{uID, uTime})
			}
		}
	}

	tweets := []int{}
	sort.Slice(users, func(i, j int) bool {
		return users[i].time > users[j].time
	})
	for i := 0; i < len(users) && i < 10; i++ {
		tweets = append(tweets, users[i].id)
	}
	return tweets
}

func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Unfollow(1, 2)
	fmt.Println(twitter.GetNewsFeed(1))
}
