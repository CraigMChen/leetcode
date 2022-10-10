-- 思路与 lc176 类似，注意 OFFSET 后面不能跟表达式
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET N = N - 1;
    RETURN (
        SELECT (
                   SELECT DISTINCT salary
                   FROM Employee
                   ORDER BY salary DESC
                   LIMIT 1 OFFSET N
               )
    );
END