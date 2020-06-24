package p355_1

type Twitter struct {
	user   map[int]map[int]struct{}
	latest *Tweet
}

type Tweet struct {
	id     int
	userId int
	pre    *Tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		user: make(map[int]map[int]struct{}),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	t := &Tweet{
		id:     tweetId,
		userId: userId,
		pre:    this.latest,
	}
	this.latest = t
}

/**
Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent.
*/
func (this *Twitter) GetNewsFeed(userId int) []int {
	tweetIds := make([]int, 0, 10)
	p := this.latest
	if _, exists := this.user[userId]; !exists {
		this.user[userId] = map[int]struct{}{userId: {}}
	}
	followees := this.user[userId]
	for len(tweetIds) < 10 && p != nil {
		if _, exists := followees[p.userId]; exists {
			tweetIds = append(tweetIds, p.id)
		}
		p = p.pre
	}
	return tweetIds
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, exists := this.user[followerId]; !exists {
		this.user[followerId] = map[int]struct{}{followerId: {}}
	}
	this.user[followerId][followeeId] = struct{}{}
}

// 不能取关自己！！！
/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, exists := this.user[followerId]; followeeId != followerId && exists {
		delete(this.user[followerId], followeeId)
	}
}

/*
执行用时 : 56 ms , 在所有 Go 提交中击败了 30.77% 的用户
内存消耗 : 10.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
