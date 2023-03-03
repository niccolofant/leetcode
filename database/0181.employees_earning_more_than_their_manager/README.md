# Employees Earning More Than Their Managers

## Problem

The `Employee` table holds all employees including their managers. Every employee has an Id, and there is also a column for the manager Id.
Write a SQL query to find out employees who earn more than their managers.
The `Employee` table is as follows:

```text
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| Id          | int     |
| Name        | varchar |
| Salary      | int     |
| ManagerId   | int     |
+-------------+---------+
```

## Example

```text
Input:
Employee table:
+-----+-------+--------+-----------+
| Id  | Name  | Salary | ManagerId |
+-----+-------+--------+-----------+
| 1   | Joe   | 70000  | 3         |
| 2   | Henry | 80000  | 4         |
| 3   | Sam   | 60000  | NULL      |
| 4   | Max   | 90000  | NULL      |
+-----+-------+--------+-----------+

Output:
+----------+
| Employee |
+----------+
| Joe      |
+----------+
```

## Solution

### Join Approach

We join the table with itself with the condition that the manager id is the same as the employee id and the employee salary is greater than the manager salary.

#### Code

```sql
SELECT e.Name AS Employee
FROM Employee AS e
JOIN Employee AS m
ON e.ManagerId = m.Id
WHERE e1.ManagerId = e2.Id
AND e1.Salary > e2.Salary
```

#### Complexity Analysis

- **Time complexity** : `O(n^2)` where `n` is the number of rows in the table.

- **Space complexity** : `O(n)` where `n` is the number of rows in the table.
