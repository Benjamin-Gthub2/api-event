SELECT COUNT(DISTINCT event_types.id) AS total
FROM event_types event_types
         INNER JOIN users creator_users ON event_types.created_by = creator_users.id
WHERE event_types.deleted_at IS NULL;
