SELECT sessions.id         AS session_id,
       sessions.start_date AS session_start_date,
       sessions.end_date   AS session_end_date,
       sessions.total_reg  AS session_total_reg,
       sessions.total_pay  AS session_total_pay,
       sessions.total_pres AS session_total_pres
FROM sessions sessions
WHERE IF(? IS NULL, TRUE, sessions.id = ?)
  AND sessions.deleted_at IS NULL;