SELECT modules.id              AS module_id,
       modules.name            AS module_name,
       modules.description     AS module_description,
       modules.code            AS module_code,
       modules.icon            AS module_icon,
       modules.position        AS module_position,
       permissions.id          AS permission_id,
       permissions.name        AS permission_name,
       permissions.description AS permission_description,
       permissions.created_at  AS permission_created_at
FROM (SELECT MAX(modules.id)          AS module_id,
             MAX(modules.name)        AS module_name,
             MAX(modules.description) AS module_description,
             modules.code             AS module_code,
             policies.merchant_id     AS policie_merchant_id,
             policies.store_id        AS policie_store_id,
             roles.id                 AS role_id,
             permissions.id           AS permission_id
      FROM core_users users
               INNER JOIN core_user_roles user_roles ON users.id = user_roles.user_id AND user_roles.deleted_at IS NULL
               INNER JOIN core_roles roles ON user_roles.role_id = roles.id AND roles.deleted_at IS NULL
               INNER JOIN core_role_policies role_policies
                          ON roles.id = role_policies.role_id AND role_policies.deleted_at IS NULL
               INNER JOIN core_policies policies
                          ON role_policies.policy_id = policies.id AND policies.deleted_at IS NULL
               INNER JOIN core_policy_permissions policy_permissions
                          ON policies.id = policy_permissions.policy_id AND policy_permissions.deleted_at IS NULL
               INNER JOIN core_permissions permissions
                          ON policy_permissions.permission_id = permissions.id AND permissions.deleted_at IS NULL
               INNER JOIN core_view_permissions view_permissions
                          ON permissions.id = view_permissions.permission_id AND view_permissions.deleted_at IS NULL
               INNER JOIN core_modules modules ON permissions.module_id = modules.id AND modules.deleted_at IS NULL
               INNER JOIN core_views views ON view_permissions.view_id = views.id AND views.deleted_at IS NULL
      WHERE users.id = ?
        AND role_policies.role_id = ?
      GROUP BY module_code, policie_merchant_id, policie_store_id, role_id, permission_id
      ORDER BY module_code, policie_merchant_id, policie_store_id) merchant_stores
         INNER JOIN core_modules modules ON merchant_stores.module_code = modules.code
         INNER JOIN core_permissions permissions ON merchant_stores.permission_id = permissions.id
         LEFT JOIN (SELECT * FROM core_merchants WHERE deleted_at IS NULL) AS merchants ON
    ((merchant_stores.policie_merchant_id IS NULL) OR
     (merchant_stores.policie_merchant_id = merchants.id))
         LEFT JOIN (SELECT * FROM core_stores WHERE deleted_at IS NULL) AS stores ON
    ((merchant_stores.policie_store_id IS NULL AND merchants.id = stores.merchant_id) OR
     (merchant_stores.policie_store_id = stores.id))
WHERE stores.id IS NOT NULL;
