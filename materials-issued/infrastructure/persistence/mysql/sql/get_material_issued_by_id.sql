SELECT materials_issued.id          AS material_issued_id,
       materials_issued.description AS material_issued_description,
       materials_issued.created_at  AS material_issued_created_at,
       creator_users.id             AS created_by_id,
       creator_users.username       AS created_by_username
FROM materials_issued
         INNER JOIN users creator_users ON materials_issued.created_by = creator_users.id
WHERE materials_issued.id = ?
  AND materials_issued.deleted_at IS NULL;
