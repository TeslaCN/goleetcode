package p355_1

import (
	"reflect"
	"testing"
)

func TestTwitter_1_1(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	t.Log(reflect.DeepEqual(twitter.GetNewsFeed(1), []int{5}))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	t.Log(reflect.DeepEqual(twitter.GetNewsFeed(1), []int{6, 5}))
	twitter.Unfollow(1, 2)
	t.Log(reflect.DeepEqual(twitter.GetNewsFeed(1), []int{5}))
	twitter.Unfollow(1, 1)
	t.Log(reflect.DeepEqual(twitter.GetNewsFeed(1), []int{5}))
}
