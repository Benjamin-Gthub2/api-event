SELECT sessions.id         AS session_id,
       sessions.total_reg  AS session_total_reg,
       sessions.total_pay  AS session_total_pay,
       sessions.total_pres AS session_total_pres
FROM sessions
WHERE sessions.deleted_at IS NULL;