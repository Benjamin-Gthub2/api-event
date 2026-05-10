SELECT workshops.id     AS workshop_id,
       events.id        AS event_id
FROM workshops workshops
         INNER JOIN events events ON workshops.event_id = events.id
WHERE workshops.id = ?
  AND workshops.deleted_at IS NULL
  AND events.deleted_at IS NULL