SELECT COUNT(DISTINCT registration_statuses.id) AS total
FROM registration_statuses
WHERE registration_statuses.deleted_at IS NULL;
