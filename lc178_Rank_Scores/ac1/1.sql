-- 用子句查询查出大于等于当前分数的去重个数
SELECT a.score AS score,
       (
           SELECT COUNT(DISTINCT b.score)
           FROM Scores b
           WHERE b.score >= a.score
       )       AS `rank`
FROM Scores a
ORDER BY a.score DESC