-- 临时表
-- 首先查出第二高的工资并将其作为临时表（否则当不存在第二高的工资时输出将为空而非 NULL）
-- 然后查询临时表中的内容
SELECT (
           SELECT DISTINCT salary
           FROM Employee
           ORDER BY salary DESC
           LIMIT 1 OFFSET 1
       ) AS SecondHighestSalary