--- 176. Second Highest Salary (Medium)
--- https://leetcode.com/problems/second-highest-salary/

-- Double MAX solution
SELECT MAX(Salary) AS SecondHighestSalary
FROM Employee
WHERE Salary < (SELECT MAX(Salary) FROM Employee)
