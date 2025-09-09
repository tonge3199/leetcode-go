package hashMap

/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.

Note:

The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.
题目大意 #
编写一个程序，通过已填充的空格来解决数独问题。一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 ‘.’ 表示。

解题思路 #
给出一个数独谜题，要求解出这个数独
解题思路 DFS 暴力回溯枚举。数独要求每横行，每竖行，每九宫格内，1-9 的数字不能重复，每次放下一个数字的时候，在这 3 个地方都需要判断一次。
另外找到一组解以后就不需要再继续回溯了，直接返回即可。
*/

// TODO: 完成
