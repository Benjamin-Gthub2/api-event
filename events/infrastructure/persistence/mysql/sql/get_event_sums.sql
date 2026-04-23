SELECT events.id            AS event_id,
       events.name          AS event_name,
       events.total_reg     AS event_total_reg,
       events.total_pay     AS event_total_pay,
       events.total_pres    AS event_total_pres,
       workshops.id         AS workshop_id,
       workshops.name       AS workshop_name,
       workshops.total_reg  AS workshop_total_reg,
       workshops.total_pay  AS workshop_total_pay,
       workshops.total_pres AS workshop_total_pres,
       sessions.id          AS session_id,
       sessions.start_date  AS session_start_date,
       sessions.end_date    AS session_end_date,
       sessions.total_reg   AS session_total_reg,
       sessions.total_pay   AS session_total_pay,
       sessions.total_pres  AS session_total_pres
FROM events events
         LEFT JOIN workshops workshops ON events.id = workshops.event_id
         LEFT JOIN sessions sessions ON workshops.id = sessions.workshop_id
WHERE events.deleted_at IS NULL
  AND workshops.deleted_at IS NULL
  AND sessions.deleted_at IS NULL;