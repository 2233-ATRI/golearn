select
    user_id,
    count(*) AS log_nums,
    DENSE_RANK() OVER (ORDER BY COUNT(*) DESC) as rk
FROM
    login_tb
WHERE
    log_port = 'pc' AND log_time BETWEEN '2022-02-09' AND '2022-02-11'
GROUP BY
    user_id
ORDER BY
    log_nums DESC, user_id ASC ;