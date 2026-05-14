SELECT views.id          AS view_id,
       views.name        AS view_name,
       views.description AS view_description,
       views.url         AS view_url,
       views.icon        AS view_icon,
       views.position    AS view_position,
       views.created_at  AS view_created_at
FROM users users
         INNER JOIN user_roles user_roles ON users.id = user_roles.user_id
         INNER JOIN roles roles ON user_roles.role_id = roles.id
         INNER JOIN role_views role_views ON roles.id = role_views.role_id
         INNER JOIN views views ON role_views.view_id = views.id
WHERE users.id = ?
  AND users.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL
  AND roles.deleted_at IS NULL
  AND views.deleted_at IS NULL
  AND role_views.deleted_at IS NULL
ORDER BY views.position;
