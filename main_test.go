package main

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	tw := NewTwitter()
	tw.PostTweet(1, 5)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("own tweet feed = %v", got)
	}
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{6, 5}) {
		t.Errorf("followee feed = %v", got)
	}
	tw.Unfollow(1, 2)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("after unfollow = %v", got)
	}
	// Feed caps at 10, newest first.
	for i := 100; i < 115; i++ {
		tw.PostTweet(3, i)
	}
	feed := tw.GetNewsFeed(3)
	if len(feed) != 10 || feed[0] != 114 || feed[9] != 105 {
		t.Errorf("cap/order = %v", feed)
	}
	// Followers don't see tweets backwards: user 2's feed lacks user 1's.
	if got := tw.GetNewsFeed(2); len(got) != 1 || got[0] != 6 {
		t.Errorf("one-way follow = %v", got)
	}
	// Unfollow of a non-followee and self-follow must not panic or change state.
	tw.Unfollow(9, 8)
	tw.Follow(4, 4)
	if got := tw.GetNewsFeed(4); len(got) != 0 {
		t.Errorf("self-follow produced a feed: %v", got)
	}
}
