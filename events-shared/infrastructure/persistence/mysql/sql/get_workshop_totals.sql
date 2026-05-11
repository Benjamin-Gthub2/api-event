SELECT workshops.id         AS workshop_id,
       workshops.total_pres AS workshop_total_pres
FROM workshops
WHERE workshops.id = ?
  AND workshops.deleted_at IS NULL;