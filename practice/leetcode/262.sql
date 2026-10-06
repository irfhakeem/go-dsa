WITH canceled_trips AS (
    SELECT id
    FROM Trips
    WHERE SUBSTRING_INDEX(status, '_', 1) = 'cancelled'
), banned_acc AS (
    SELECT users_id AS id
    FROM Users
    WHERE banned = 'Yes'
)
SELECT t.request_at AS `Day`, ROUND(SUM(c.id IS NOT NULL)/COUNT(*), 2) AS `Cancellation Rate`
FROM Trips t
LEFT JOIN canceled_trips c ON c.id = t.id
WHERE t.client_id NOT IN (
    SELECT id
    FROM banned_acc
)
AND t.driver_id NOT IN (
    SELECT id
    FROM banned_acc
)
AND t.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY t.request_at
