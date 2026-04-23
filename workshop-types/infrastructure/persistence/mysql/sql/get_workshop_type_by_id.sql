SELECT workshop_types.id          AS workshop_type_id,
       workshop_types.code        AS workshop_type_code,
       workshop_types.description AS workshop_type_description,
       workshop_types.enable      AS workshop_type_enable,
       workshop_types.created_at  AS workshop_type_created_at,
       creator_users.id           AS created_by_id,
       creator_users.username     AS created_by_username
FROM workshop_types workshop_types
         INNER JOIN users creator_users ON workshop_types.created_by = creator_users.id
WHERE workshop_types.id = ?
  AND workshop_types.deleted_at IS NULL;
