-- 临时表
SELECT Email
FROM (
         SELECT Email, count(Email) AS cnt
         FROM Person
         GROUP BY Email
     ) a
WHERE cnt >= 2