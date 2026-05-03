SELECT registration_statuses.id          AS registration_status_id,
       registration_statuses.code        AS registration_status_code,
       registration_statuses.description AS registration_status_description,
       registration_statuses.position    AS registration_status_position,
       registration_statuses.enable      AS registration_status_enable,
       registration_statuses.created_at  AS registration_status_created_at
FROM registration_statuses
WHERE registration_statuses.deleted_at IS NULL
ORDER BY registration_statuses.position ASC
LIMIT ? OFFSET ?;
