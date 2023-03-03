--- 175. Combine Two Tables (Easy)
--- https://leetcode.com/problems/combine-two-tables/

--- Join solution
SELECT p.FirstName, p.LastName, a.City, a.State
FROM Person AS p
LEFT JOIN Address AS a
ON p.PersonId = a.PersonId