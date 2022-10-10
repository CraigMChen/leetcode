-- 使用 dense_rank() over 函数
SELECT score, dense_rank() OVER (ORDER BY score DESC) AS `rank`
FROM Scores