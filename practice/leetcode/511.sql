WITH first_logins AS (
    SELECT a.player_id AS player_id, MIN(a.event_date) AS first_login
    FROM Activity a
    GROUP BY a.player_id
)
SELECT player_id, first_login
FROM first_logins
