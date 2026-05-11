// n = 97
// f(n) = f(n - 1) + f(n - 2)

func climbStairs(n int) int {
   cache := make([]int, n + 3)
   cache [1] = 1
   cache [2] = 2

   for i := 3; i <= n; i++ {
    cache[i] = cache[i - 1] + cache [i - 2]
   }

   return cache[n]
}
