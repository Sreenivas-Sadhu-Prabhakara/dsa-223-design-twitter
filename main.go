package main

import (
	"fmt"
	"sort"
)

type tweet struct {
	id, time int
}

// Twitter is the classic feed-design exercise: post, follow, unfollow, and
// a news feed of the 10 most recent tweets across self and followees.
type Twitter struct {
	clock   int
	tweets  map[int][]tweet      // userId -> tweets in posting order
	follows map[int]map[int]bool // userId -> set of followees
}

// NewTwitter returns an empty service.
func NewTwitter() *Twitter {
	return &Twitter{tweets: map[int][]tweet{}, follows: map[int]map[int]bool{}}
}

// PostTweet records a tweet by userId.
func (t *Twitter) PostTweet(userId, tweetId int) {
	t.clock++
	t.tweets[userId] = append(t.tweets[userId], tweet{tweetId, t.clock})
}

// Follow makes follower see followee's tweets. Self-follows are ignored.
func (t *Twitter) Follow(follower, followee int) {
	if follower == followee {
		return
	}
	if t.follows[follower] == nil {
		t.follows[follower] = map[int]bool{}
	}
	t.follows[follower][followee] = true
}

// Unfollow reverses Follow; unfollowing a non-followee is a no-op.
func (t *Twitter) Unfollow(follower, followee int) {
	delete(t.follows[follower], followee)
}

// GetNewsFeed returns the ids of the 10 most recent tweets by userId or
// anyone they follow, newest first.
func (t *Twitter) GetNewsFeed(userId int) []int {
	candidates := []tweet{}
	// Own last 10 plus each followee's last 10 are the only possible answers.
	take := func(list []tweet) {
		start := 0
		if len(list) > 10 {
			start = len(list) - 10
		}
		candidates = append(candidates, list[start:]...)
	}
	take(t.tweets[userId])
	for followee := range t.follows[userId] {
		take(t.tweets[followee])
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i].time > candidates[j].time })
	if len(candidates) > 10 {
		candidates = candidates[:10]
	}
	out := make([]int, len(candidates))
	for i, c := range candidates {
		out[i] = c.id
	}
	return out
}

func main() {
	tw := NewTwitter()
	tw.PostTweet(1, 5)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Unfollow(1, 2)
	fmt.Println(tw.GetNewsFeed(1))
}
