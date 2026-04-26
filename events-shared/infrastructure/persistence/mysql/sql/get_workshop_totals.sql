SELECT workshops.id         AS workshop_id,
       workshops.total_reg  AS workshop_total_reg,
       workshops.total_pay  AS workshop_total_pay,
       workshops.total_pres AS workshop_total_pres
FROM workshops
WHERE workshops.id = ?
  AND workshops.deleted_at IS NULL;