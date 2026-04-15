SELECT roles.id          AS role_id,
       roles.name        AS role_name,
       roles.description AS role_description,
       roles.enable      AS role_enable,
       roles.created_at  AS role_created_at
FROM core_users users
         INNER JOIN core_user_roles user_roles
                    ON users.id = user_roles.user_id AND user_roles.deleted_at IS NULL AND user_roles.enable = 1
         INNER JOIN core_roles roles ON user_roles.role_id = roles.id AND roles.deleted_at IS NULL AND roles.enable = 1
         INNER JOIN core_role_policies role_policies
                    ON roles.id = role_policies.role_id AND role_policies.deleted_at IS NULL AND
                       role_policies.enable = 1
         INNER JOIN core_policies policies ON role_policies.policy_id = policies.id AND policies.deleted_at IS NULL AND
                                              policies.enable = 1
         INNER JOIN core_policy_permissions policy_permissions ON policies.id = policy_permissions.policy_id AND
                                                                  policy_permissions.deleted_at IS NULL AND
                                                                  policy_permissions.enable = 1
         INNER JOIN core_permissions permissions ON policy_permissions.permission_id = permissions.id AND
                                                    permissions.deleted_at IS NULL
WHERE policies.merchant_id IS NULL
  AND policies.store_id IS NULL
  AND users.id = ?;
