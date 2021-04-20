package twitter

// Tweet is a type that represents a tweet.
// It contains the Content and Date (year) of the tweet.
type Tweet struct {
	Content string
	Date    int
}

// GetFirstAndLastTweets gets the first and lastest tweets of the user.
// It assumes that the tweets are ordered by time, with the oldest
// tweet as the first element of the slice and the most recent tweet
// as the last element of the slice.
// It has time complexity O(1) and space complexity O(1).
func GetFirstAndLastTweets(tweets []Tweet) (first, last Tweet) {
	first = tweets[0]
	last = tweets[len(tweets)-1]
	return
}

// GetTimeDiffBetweenFirstAndLastTweets1 returns the time difference
// between the user's first and latest tweets. It does not assume
// that the input slice is ordered by date.
// It has time complexity O(n^2) and space complexity O(1).
func GetTimeDiffBetweenFirstAndLastTweets1(tweets []Tweet) int {
	diff := 0

	for _, tweet1 := range tweets {
		for _, tweet2 := range tweets {
			tweetDiff := tweet1.Date - tweet2.Date
			if tweetDiff < 0 {
				tweetDiff = -tweetDiff
			}
			if tweetDiff > diff {
				diff = tweetDiff
			}
		}
	}

	return diff
}

// GetTimeDiffBetweenFirstAndLastTweets2 returns the time difference
// between the user's first and latest tweets. It does not assume
// that the input slice is ordered by date. It improves time complexity
// without adding space complexity.
// It has time complexity O(n) and space complexity O(1).
func GetTimeDiffBetweenFirstAndLastTweets2(tweets []Tweet) int {
	minDate := int(^uint(0) >> 1) // max int
	maxDate := 0

	for _, tweet := range tweets {
		if tweet.Date < minDate {
			minDate = tweet.Date
		}
		if tweet.Date > maxDate {
			maxDate = tweet.Date
		}
	}

	return maxDate - minDate
}
