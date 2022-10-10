-- JOIN 和 IN
SELECT d.name AS Department, e.name AS Employee, e.salary AS Salary
FROM Employee e
         JOIN Department d ON e.departmentId = d.id
WHERE (
              (e.departmentId, e.salary) IN
              (
                  SELECT departmentId, max(salary)
                  FROM Employee
                  GROUP BY departmentId
              )
    )