SELECT sessions.id            AS session_id,
       sessions.start_date    AS session_start_date,
       sessions.end_date      AS session_end_date,
       sessions.created_at    AS session_created_at,
       workshops.id           AS workshop_id,
       workshops.name         AS workshop_name,
       workshops.shortname    AS workshop_shortname,
       workshops.code         AS workshop_code,
       workshops.capacity     AS workshop_capacity,
       creator_users.id       AS created_by_id,
       creator_users.username AS created_by_username
FROM sessions sessions
         INNER JOIN workshops workshops ON sessions.workshop_id = workshops.id
         INNER JOIN users creator_users ON sessions.created_by = creator_users.id
WHERE IF(? IS NULL, TRUE, sessions.workshop_id = TRIM(?))
  AND IF(? IS NULL, TRUE, DATE(sessions.start_date) >= DATE(?))
  AND IF(? IS NULL, TRUE, DATE(sessions.end_date) <= DATE(?))
  AND sessions.deleted_at IS NULL
LIMIT ? OFFSET ?;
