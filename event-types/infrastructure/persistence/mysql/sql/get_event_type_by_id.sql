SELECT event_types.id          AS event_type_id,
       event_types.code        AS event_type_code,
       event_types.description AS event_type_description,
       event_types.enable      AS event_type_enable,
       event_types.created_at  AS event_type_created_at,
       creator_users.id        AS created_by_id,
       creator_users.username  AS created_by_username
FROM event_types event_types
         INNER JOIN users creator_users ON event_types.created_by = creator_users.id
WHERE event_types.id = ?
  AND event_types.deleted_at IS NULL;
