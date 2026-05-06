SELECT MAX(merchant_stores.module_id)          AS module_id,
       MAX(merchant_stores.module_name)        AS module_name,
       MAX(merchant_stores.module_description) AS module_description,
       merchant_stores.module_code             AS module_code_middle,
       merchants.id                            AS merchant_id_middle,
       stores.id                               AS store_id_middle
FROM (SELECT MAX(modules.id)          AS module_id,
             MAX(modules.name)        AS module_name,
             MAX(modules.description) AS module_description,
             modules.code             AS module_code,
             policies.merchant_id     AS policie_merchant_id,
             policies.store_id        AS policie_store_id
      FROM core_users users
               INNER JOIN core_user_roles user_roles ON users.id = user_roles.user_id
               INNER JOIN core_roles roles ON user_roles.role_id = roles.id
               INNER JOIN core_role_policies role_policies ON roles.id = role_policies.role_id
               INNER JOIN core_policies policies ON role_policies.policy_id = policies.id
               INNER JOIN core_policy_permissions policy_permissions ON policies.id = policy_permissions.policy_id
               INNER JOIN core_permissions permissions ON policy_permissions.permission_id = permissions.id
               INNER JOIN core_view_permissions ON permissions.id = core_view_permissions.permission_id
               INNER JOIN core_modules modules ON permissions.module_id = modules.id
               INNER JOIN core_views views ON core_view_permissions.view_id = views.id
      WHERE users.id = ?
        AND users.deleted_at IS NULL
        AND user_roles.deleted_at IS NULL
        AND roles.deleted_at IS NULL
        AND role_policies.deleted_at IS NULL
        AND policies.deleted_at IS NULL
        AND policy_permissions.deleted_at IS NULL
        AND permissions.deleted_at IS NULL
        AND modules.deleted_at IS NULL
        AND views.deleted_at IS NULL
        AND core_view_permissions.deleted_at IS NULL
      GROUP BY module_code, policie_merchant_id, policie_store_id
      ORDER BY module_code, policie_merchant_id, policie_store_id) AS merchant_stores
         LEFT JOIN core_merchants AS merchants ON
    ((merchant_stores.policie_merchant_id IS NULL) OR
     (merchant_stores.policie_merchant_id = merchants.id))
         LEFT JOIN core_stores AS stores ON
    ((merchant_stores.policie_store_id IS NULL AND merchants.id = stores.merchant_id) OR
     (merchant_stores.policie_store_id = stores.id))
WHERE stores.id IS NOT NULL
GROUP BY module_code_middle, merchant_id_middle, store_id_middle
