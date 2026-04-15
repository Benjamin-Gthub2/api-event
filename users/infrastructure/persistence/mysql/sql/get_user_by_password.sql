SELECT users.id          AS user_id,
       users.username    AS user_name,
       users.created_at  AS user_created_at,
       types.id          AS user_type_id,
       types.description AS user_type_description,
       types.code        AS user_type_code
FROM users users
         INNER JOIN user_types types ON users.type_id = types.id
         INNER JOIN people people ON users.id = people.user_id
WHERE users.deleted_at IS NULL
  AND people.deleted_at IS NULL
  AND users.username = ?
  AND users.password_hash = ?
  AND people.enable = 1;