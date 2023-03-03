# Duplicate Emails

## Problem

Write a SQL query to find all duplicate emails in a table named `Person`.
Table schema is as following:

```text
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| Id          | int     |
| Email       | varchar |
+-------------+---------+
```

## Example

```text
Input:

Person table:
+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+

Output:
+---------+
| Email   |
+---------+
| a@b.com |
+---------+
```

## Solution

### Join Approach

We join the table with itself with the condition that the email is the same and the id is different. We then select distinct emails.

#### Code

```sql
SELECT DISTINCT p1.email
FROM person AS p1, person AS p2
WHERE p1.id != p2.id
AND p1.email = p2.email
```

#### Complexity Analysis

- **Time complexity** : `O(n^2)` where `n` is the number of rows in the table.

- **Space complexity** : `O(n)` where `n` is the number of rows in the table.
