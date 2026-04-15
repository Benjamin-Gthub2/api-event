SELECT COUNT(*)
FROM user_types
WHERE deleted_at IS NULL
ORDER BY created_at;