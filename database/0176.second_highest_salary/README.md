# Second Highest Salary

## Problem

Write an SQL query to report the second highest salary from the `Employee` table.
If there is no second highest salary, the query should report `null`.

The `Employee` table is as follows:

```text
+-------------+------+
| Column Name | Type |
+-------------+------+
| Id          | int  |
| Salary      | int  |
+-------------+------+
```

##  Example

```text
Input:
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+

Output:
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
```

```text
Input:
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+

Output:
+---------------------+
| SecondHighestSalary |
+---------------------+
| null                |
+---------------------+
```

## Solution

### Double MAX Approach

We can use the `MAX` function twice to find the second highest salary. The first `MAX` function will find the highest salary, and the second `MAX` function will find the second highest salary. We need to use a `WHERE` clause to exclude the highest salary from the second `MAX` function.

#### Code

```sql
SELECT MAX(Salary) AS SecondHighestSalary
FROM Employee
WHERE Salary < (SELECT MAX(Salary) FROM Employee)
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of rows in the table.

- **Space Complexity** : `O(1)`.
