SELECT workshops.id         AS workshop_id,
       workshops.name       AS workshop_name,
       workshops.capacity   AS workshop_capacity,
       workshops.total_reg  AS workshop_total_reg,
       workshops.total_pay  AS workshop_total_pay,
       workshops.total_pres AS workshop_total_pres,
       sessions.id          AS session_id,
       sessions.start_date  AS session_start_date,
       sessions.end_date    AS session_end_date,
       sessions.total_reg   AS session_total_reg,
       sessions.total_pay   AS session_total_pay,
       sessions.total_pres  AS session_total_pres
FROM workshops workshops
         LEFT JOIN sessions sessions ON workshops.id = sessions.workshop_id
WHERE IF(? IS NULL, TRUE, workshops.id = ?)
  AND workshops.deleted_at IS NULL
  AND sessions.deleted_at IS NULL;