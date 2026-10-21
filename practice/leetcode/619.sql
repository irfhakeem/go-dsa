WITH single_num AS (
    SELECT num
    FROM MyNumbers
    GROUP BY num
    HAVING COUNT(num) = 1
)
SELECT (
    SELECT num
    FROM single_num
    ORDER BY num DESC
    LIMIT 1
) AS num
