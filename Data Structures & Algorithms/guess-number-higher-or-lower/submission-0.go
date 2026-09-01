/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi - lo) / 2
		ans := guess(mid)
		if ans <= 0 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
