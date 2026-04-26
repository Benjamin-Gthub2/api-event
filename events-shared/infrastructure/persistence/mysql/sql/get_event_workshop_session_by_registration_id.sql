SELECT sessions.id      AS session_id,
       workshops.id     AS workshop_id,
       events.id        AS event_id
FROM sessions sessions
         INNER JOIN workshops workshops ON sessions.workshop_id = workshops.id
         INNER JOIN events events ON workshops.event_id = events.id
WHERE sessions.id = ?
  AND sessions.deleted_at IS NULL
  AND workshops.deleted_at IS NULL
  AND events.deleted_at IS NULL