SELECT merchants.id          AS merchant_id,
       merchants.name        AS merchant_name,
       merchants.description AS merchant_description,
       stores.id             AS store_id,
       stores.name           AS store_name,
       roles.id              AS role_id,
       roles.name            AS role_name,
       roles.description     AS role_description,
       roles.enable          AS role_enable,
       roles.created_at      AS role_created_at
FROM (SELECT policies.merchant_id AS policie_merchant_id,
             policies.store_id    AS policie_store_id,
             roles.id             AS role_id
      FROM core_users users
               INNER JOIN core_user_roles user_roles ON users.id = user_roles.user_id AND user_roles.deleted_at IS NULL
               INNER JOIN core_roles roles ON user_roles.role_id = roles.id AND roles.deleted_at IS NULL
               INNER JOIN core_role_policies role_policies
                          ON roles.id = role_policies.role_id AND role_policies.deleted_at IS NULL
               INNER JOIN core_policies policies
                          ON role_policies.policy_id = policies.id AND policies.deleted_at IS NULL
      WHERE users.id = ?
      ORDER BY policie_merchant_id, policie_store_id) merchant_stores
         LEFT JOIN (SELECT * FROM core_stores WHERE deleted_at IS NULL) AS stores
                   ON ((merchant_stores.policie_store_id IS NULL AND policie_merchant_id = stores.merchant_id) OR
                       (merchant_stores.policie_store_id = stores.id))
         LEFT JOIN core_merchants merchants ON stores.merchant_id = merchants.id
         INNER JOIN core_roles roles ON merchant_stores.role_id = roles.id AND roles.deleted_at IS NULL
WHERE stores.id IS NOT NULL;