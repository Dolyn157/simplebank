-- Tx: transfer 10$ from account1 to account2
BEGIN;

UPDATE accounts SET balance = balance - 15 WHERE id = 1 RETURNING *;
UPDATE accounts SET balance = balance + 15 WHERE id = 2 RETURNING *;

ROLLBACK;

-- Tx2: transfer 10$ from account2 to account1
BEGIN;

UPDATE accounts SET balance = balance - 10 WHERE id =2 RETURNING *;
UPDATE accounts SET balance = balance + 10 WHERE id =1 RETURNING *;

ROLLBACK;


SELECT a.datname,
        a.application_name,
        l.relation::regclass,
        l.mode,
        l.locktype,
        l.GRANTED,
        a.usename,
        a.query,
        a.pid
FROM pg_stat_activity a
JOIN pg_locks l ON l.pid = a.pid
WHERE a.application_name = 'psql'
ORDER BY a.pid;