SELECT attendances.id         AS attendance_id,
       attendances.created_at AS attendance_created_at,
       creator_users.id       AS created_by_id,
       creator_users.username AS created_by_username
FROM attendances
         INNER JOIN users creator_users ON attendances.created_by = creator_users.id
WHERE attendances.id = ?
  AND attendances.deleted_at IS NULL;
