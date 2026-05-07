SELECT schema_name, x_tenant_id
FROM tenants
WHERE (x_tenant_id = $1 OR host = $2)
  AND enable = true;
