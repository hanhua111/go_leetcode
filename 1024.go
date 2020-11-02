//# class Solution:
//#     def videoStitching(self, clips: List[List[int]], T: int) -> int:
//#         res = 0
//#         l = len(clips)
//#         clips = sorted(clips)
//#         dp = [[0,0]]
//#         for ind, element in enumerate(clips):
//#             if dp[-1][-1] >= T:
//#                 return res
//#             elif ind == 0:
//#                 if element[0] != 0:
//#                     return -1
//#                 else:
//#                     dp.append(element)
//#                     res += 1
//#             else:
//#                 if element[0] > dp[-1][-1]:
//#                     return -1
//#                 elif element[0] == dp[-1][-1]:
//#                     res += 1
//#                     dp.append(element)
//#                 elif element[0] < dp[-1][1]:
//#                     if element[0] > dp[-1][0]:
//#                         if
//#                     elif element[0]
//#         if dp[-1][-1] < T:
//#             return -1
//#         return res
//class Solution:
//def videoStitching(self, clips: List[List[int]], T: int) -> int:
//maxn = [0] * T
//last = ret = pre = 0
//for a, b in clips:
//if a < T:
//maxn[a] = max(maxn[a], b)
//
//for i in range(T):
//last = max(last, maxn[i])
//if i == last:j
//return -1
//if i == pre:
//ret += 1
//pre = last

return ret