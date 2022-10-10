-- HAVING
SELECT Email
FROM Person
GROUP BY Email
HAVING count(Email) >= 2