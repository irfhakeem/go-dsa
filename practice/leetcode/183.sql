SELECT c.name as Customers
FROM Customers c
WHERE c.id NOT IN (
    SELECT c.id
    FROM Customers c
    JOIN Orders o ON o.customerId = c.id
)
