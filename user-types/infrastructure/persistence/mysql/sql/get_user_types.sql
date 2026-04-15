SELECT id,
       description,
       code,
       enable,
       created_at
FROM user_types
WHERE deleted_at IS NULL
ORDER BY created_at
LIMIT ? OFFSET ?;
