WITH highest_salary AS (
    SELECT departmentId, MAX(salary) AS max_salary
    FROM Employee
    GROUP BY departmentId
)
SELECT d.name AS Department, e.name AS Employee, e.salary AS Salary
FROM Employee e
JOIN Department d ON d.id = e.departmentId
JOIN highest_salary h ON h.departmentId = e.departmentId AND h.max_salary = e.salary;
