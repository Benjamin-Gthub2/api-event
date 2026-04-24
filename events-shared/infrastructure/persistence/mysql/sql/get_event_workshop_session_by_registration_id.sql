SELECT registrations.id AS registration_id,
       sessions.id      AS session_id,
       workshops.id     AS workshop_id,
       events.id        AS event_id
FROM registrations registrations
         INNER JOIN sessions sessions ON registrations.session_id = sessions.id
         INNER JOIN workshops workshops ON sessions.workshop_id = workshops.id
         INNER JOIN events events ON workshops.event_id = events.id
WHERE registrations.id = ?
  AND registrations.deleted_at IS NULL
  AND sessions.deleted_at IS NULL
  AND workshops.deleted_at IS NULL
  AND events.deleted_at IS NULL