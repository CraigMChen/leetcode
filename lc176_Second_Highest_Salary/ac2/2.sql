-- IFNULL 字句
-- IFNULL(expression, alt_value)
SELECT IFNULL(
               (SELECT DISTINCT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET 1),
               NULL
           ) AS SecondHighestSalary