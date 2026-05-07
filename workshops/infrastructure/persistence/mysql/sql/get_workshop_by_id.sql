SELECT workshops.id               AS workshop_id,
       workshops.name             AS workshop_name,
       workshops.shortname        AS workshop_shortname,
       workshops.code             AS workshop_code,
       workshops.capacity         AS workshop_capacity,
       workshops.created_at       AS workshop_created_at,
       workshop_types.id          AS workshop_type_id,
       workshop_types.code        AS workshop_type_code,
       workshop_types.description AS workshop_type_description,
       events.id                  AS event_id,
       events.name                AS event_name,
       events.code                AS event_code,
       creator_users.id           AS created_by_id,
       creator_users.username     AS created_by_username
FROM workshops workshops
         INNER JOIN workshop_types workshop_types ON workshops.type_id = workshop_types.id
         INNER JOIN events events ON workshops.event_id = events.id
         INNER JOIN users creator_users ON workshops.created_by = creator_users.id
WHERE workshops.id = ?
  AND workshops.deleted_at IS NULL;
