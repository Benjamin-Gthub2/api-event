SELECT schema_name, x_tenant_id
FROM tenants
WHERE (x_tenant_id = ? OR host = ?)
  AND enable = 1;
