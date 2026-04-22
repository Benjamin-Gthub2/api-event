SELECT modules.id            AS module_id,
       modules.name          AS module_name,
       modules.description   AS module_description,
       modules.code          AS module_code,
       modules.icon          AS module_icon,
       MAX(modules.position) AS module_position,
       modules.created_at    AS module_created_at,
       views.id              AS view_id,
       views.name            AS view_name,
       views.description     AS view_description,
       views.url             AS view_url,
       views.icon            AS view_icon,
       views.position        AS view_position,
       views.created_at      AS view_created_at
FROM users users
         INNER JOIN user_roles user_roles ON users.id = user_roles.user_id
         INNER JOIN roles roles ON user_roles.role_id = roles.id
         INNER JOIN role_views role_views ON roles.id = role_views.role_id
         INNER JOIN views views ON role_views.view_id = views.id
         INNER JOIN modules modules ON views.module_id = modules.id
WHERE users.id = ?
  AND users.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL
  AND roles.deleted_at IS NULL
  AND modules.deleted_at IS NULL
  AND views.deleted_at IS NULL
  AND role_views.deleted_at IS NULL
GROUP BY modules.id, views.id, views.position
ORDER BY module_position, views.position;
