UPDATE core_merchants
SET name        = TRIM(?),
    description = TRIM(?),
    code        = TRIM(?),
    phone       = TRIM(?),
    document    = TRIM(?),
    address     = TRIM(?),
    industry    = TRIM(?),
    enable      = ?
WHERE id = ?;
