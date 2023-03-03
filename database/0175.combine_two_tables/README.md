# Combine Two Tables

## Problem

Write a SQL query for a report that provides the following information for each person in the `Person` table, regardless if there is an address for each of those people.
If the address of a person is not known, the report should still contain that person with `NULL` in the corresponding columns.
Table schema is as following:

```text
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
```

## Example

```text
Input:
Person table:
+---------+----------+-----------+
| PersonId| FirstName| LastName  |
+---------+----------+-----------+
| 1       | Jon      | Snow      |
| 2       | Rob      | Stark     |
| 3       | Rob      | Baratheon |
+---------+----------+-----------+

Address table:
+-----------+-----------+-----------+--------------+
| AddressId | PersonId  | City      | State        |
+-----------+-----------+-----------+--------------+
| 1         | 1         | New York  | New York     |
| 2         | 2         | New York  | New York     |
+-----------+-----------+-----------+--------------+

Output:
+----------+----------+-----------+--------------+
| FirstName| LastName | City      | State        |
+----------+----------+-----------+--------------+
| Jon      | Snow     | New York  | New York     |
| Rob      | Stark    | New York  | New York     |
| Rob      | Baratheon| NULL      | NULL         |
+----------+----------+-----------+--------------+
```

## Solution

### Join Approach

We join the two tables with the condition that the person id is the same. We then select the first name, last name, city, and state.

#### Code

```sql
SELECT p.FirstName, p.LastName, a.City, a.State
FROM Person AS p
LEFT JOIN Address AS a
ON p.PersonId = a.PersonId
```

#### Complexity Analysis

- **Time complexity** : `O(n)` where `n` is the number of rows in the table.

- **Space complexity** : `O(n)` where `n` is the number of rows in the table.
