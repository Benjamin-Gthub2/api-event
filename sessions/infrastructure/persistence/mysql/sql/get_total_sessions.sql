SELECT COUNT(DISTINCT sessions.id) AS total
FROM sessions sessions
         INNER JOIN workshops workshops ON sessions.workshop_id = workshops.id
         INNER JOIN users creator_users ON sessions.created_by = creator_users.id
WHERE IF(? IS NULL, TRUE, sessions.workshop_id = TRIM(?))
  AND IF(? IS NULL, TRUE, DATE(sessions.start_date) >= DATE(?))
  AND IF(? IS NULL, TRUE, DATE(sessions.end_date) <= DATE(?))
  AND sessions.deleted_at IS NULL;
