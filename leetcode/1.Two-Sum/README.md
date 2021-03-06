# [1. Two Sum](https://leetcode.com/problems/two-sum/)

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- 2 <= nums.length <= 105
- -109 <= nums[i] <= 109
- -109 <= target <= 109
- Only one valid answer exists.



**个人体会：**

>这道题的思路是：将数组中的元素放在map中，索引为该元素，通过target-i=j，得到期望的j，再判断j是否存在。
>
>第一次实现参考handleOne，考虑的比较简单，代码很简单，就不多说了。
>
>参考了其他人的代码，即handleTwo，发现别人的代码更为优雅一些，时间复杂度也更优些。

