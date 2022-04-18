package array

/**
	在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄。他的攻击可以让敌方英雄艾希（编者注：寒冰射手）进入中毒状态。

当提莫攻击艾希，艾希的中毒状态正好持续duration 秒。

正式地讲，提莫在t发起发起攻击意味着艾希在时间区间 [t, t + duration - 1]（含 t 和 t + duration - 1）处于中毒状态。

如果提莫在中毒影响结束前再次攻击，中毒状态计时器将会重置，在新的攻击之后，中毒影响将会在duration秒后结束。

给你一个非递减的整数数组timeSeries，其中timeSeries[i]表示提莫在timeSeries[i]秒时对艾希发起攻击，以及一个表示中毒持续时间的整数duration 。

返回艾希处于中毒状态的总秒数。
*/

func findPoisonedDuration(timeSeries []int, duration int) int {
	var sum int
	for i := 1; i < len(timeSeries); i++ {
		end := timeSeries[i-1] + duration
		if end < timeSeries[i] {
			sum += duration
		} else {
			sum += timeSeries[i] - timeSeries[i-1]
		}
	}
	sum += duration
	return sum
}
