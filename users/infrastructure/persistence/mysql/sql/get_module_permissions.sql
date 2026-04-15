SELECT permissions.id   AS permission_id,
       permissions.code AS permission_code
FROM core_users users
         INNER JOIN core_user_roles user_roles ON users.id = user_roles.user_id
         INNER JOIN core_roles roles ON user_roles.role_id = roles.id
         INNER JOIN core_role_policies role_policies ON roles.id = role_policies.role_id
         INNER JOIN core_policies policies ON role_policies.policy_id = policies.id
         INNER JOIN core_policy_permissions policy_permissions ON policies.id = policy_permissions.policy_id
         INNER JOIN core_permissions permissions ON policy_permissions.permission_id = permissions.id
         INNER JOIN core_modules modules ON permissions.module_id = modules.id
WHERE modules.code = ?
  AND users.id = ?
  AND users.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL
  AND roles.deleted_at IS NULL
  AND role_policies.deleted_at IS NULL
  AND policies.deleted_at IS NULL
  AND policy_permissions.deleted_at IS NULL
  AND permissions.deleted_at IS NULL
  AND modules.deleted_at IS NULL
GROUP BY permission_id;
