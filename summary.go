package main

// 硬币 数组中和为 target的组合数 背包解决

//最小路径和，动态规划

//catchRain 先求左边最大值 再求右边最大值 之后遍历找左边右边较小的值 减去当前的高度就是这个地方能接多少水加和就是

// 最大矩形面积，枚举高度 找到左右第一个小于当前高度的index area = right-left-1 * 优化方法是单调栈和数组
// 先从左到右遍历并将[index,height] 压入栈中 条件是height > 大于栈顶的height arr 对应位置填栈顶的index，如果小于则pop
// 知道栈顶的height 小于 当前的heihght arr 中0 对应的是-1 最高对应len 之后在从右到左遍历记录arr 根据 left arr 和 right arr
// 计算面积 (rightarr[i]-leftarr[i]) arr[i]height

// 全排列，depth 控制选择次数递归深度，arr[] 记录当前 排列，map 用来记录哪些用过了哪些没用过，*[][]记录结果 递归结束depth==
// len(s)

// 买股票1，只能买一次，遍历一次同时计算第i天之前的minPrice 和假设今天卖出的收益并比较获得最大收益
// 买股票2 能买多次 买前要卖出能买多次但买之前手里的股票必须卖出 二维数组dp[i][j] i 第i天 j 0，1 是否持有股票 值为收益
////没有持有股票 要不是前一天就没持有，要么昨天今天卖掉
//		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+arr[i])
//		//没有持有股票 要不是前一天就持有，要么是昨天没有今天刚买
//		dp[i][1] = max(dp[i-1][1],dp[i-1][0]-arr[i])


// 子集得看通俗的做法

//岛屿面积 和岛屿数量 广度优先遍历 第一次碰到1 将 节点放入 队列中并且将紧邻的节点也放入 继续遍历队列遇到1 再放入到对列，数量就是
// 初始化队列的次数 每个岛的面积就是 每次遍历完队列的次数

// 旋转矩阵 先上下翻转再对角线翻转

// 数组字典序下一个排列，先从最后一个元素向前比遍历，找到第一个不符合递增序列的index，之后在 [index+1:] 范围内寻找第一个大于
//index 元素的位置两个数交换，之后将之前的递减序列调整为递增序列

// 最长连续序列的长度（任意组合）不管之前的顺序，做法是先遍历数组将数组放入一个map中去重，之后 遍历之后检查map[nums[i]-1]是否存在


// 如果不存在则以自己为起点 检查map[nums[i]+1] 如果存在当前长度+1 之后和max 比较，如果存在map[nums[i]-1] 则不管后续遍历到nums[i]-1时
//会把自己算进去

// 打家劫舍 动态规划 不能打劫相邻的房间 则准备一个dp 数组 表示如果第i间房存在能强多少 因为每家都有钱所以肯定抢的越多越好
// dp[0] = nums[0]  dp[1] = max(dp[0],nums[1]+0) dp[i] = max(dp[i-1],dp[i-2]+nums[i])

// 最长子序列长度
// 1. 只求最长自序列的长度 动态规划 dp[i] 用来标识如果使用nums[i]之后最长递增序列的长度，dp[0]=1, 遍历nums 到位置i的时候，从
// 0 开始遍历nums if nums[j] < nums[i] 则判断dp[j]+1 是否 > maxLen (下一次遍历清空) max = dp[j]+1 最后将dp[j] = maxLen
// 最后遍历dp 找到最大值
// 2.求最长子序列的数量除了维护dp 还需要维护一个counts 数组 p[j]+1 是否 > maxLen count = count[j] maxlen = p[j]+1；
// p[j]+1 == maxLen 则count += count[j]  count[i] = count
// 最后遍历dp 找到最长长度 以及 最长长度 对应count 中的数量相加
// 3. 除了所有递增序列 维护dp 还维护[][][]int{} 用来存储考虑nums[i]时最长的序列

// 背包问题 零钱组成目标的最小组合数，和为某个值的组合数
// 先遍历物品 在遍历背包 求组合数 先遍历背包 在遍历物品是排列数
// 递增遍历背包 可重复选择物品 递减遍历背包不能重复使用物品
// 符合 能够通过 dp[i-1][j] 和dp[i][j-coins[i]] => dp[i][j] 其实只需要一个数组即可 dp[j] old 和 dp[j-conis[i]]= >
//dp[j]  当前遍历到哪里就是在考虑是否使用某个物品

// 最优解要考虑动态规划


// todo 扔鸡蛋问题，分隔数组最大值


