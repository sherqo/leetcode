/*
 *we should have a map, with a key=key
 *and the value will be an array of struct{ timestamp, val }
 *when you set a key, you append the struct and you sure it's sorted
 *when you get a key, you binary search over the whole array for a
 *matching timestamp and return its value
 */

type TimeValPair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	mp map[string][]TimeValPair
}

func Constructor() TimeMap {
	return TimeMap{mp: make(map[string][]TimeValPair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.mp[key] = append(this.mp[key], TimeValPair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	var timeVals []TimeValPair = this.mp[key]

	n := len(timeVals)
	if n == 0 || timeVals[0].timestamp > timestamp {
		return ""
	}

	l, r := 0, n-1
	for l < r {
		mid := l + (r-l+1)/2

		if timeVals[mid].timestamp <= timestamp { // can work, keep it
			l = mid
		} else { // it's big, get rid of it
			r = mid - 1
		}
	}

	return timeVals[l].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
