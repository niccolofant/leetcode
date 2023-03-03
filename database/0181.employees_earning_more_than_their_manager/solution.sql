--- 181. Employees Earning More Than Their Managers (Easy)
--- https://leetcode.com/problems/employees-earning-more-than-their-managers/

--- Join solution
SELECT e.Name AS Employee
FROM Employee e
JOIN Employee m
ON e.ManagerId = m.Id
WHERE e.Salary > m.Salary