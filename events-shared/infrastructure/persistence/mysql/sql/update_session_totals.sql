UPDATE sessions
SET total_reg  = COALESCE(?, total_reg),
    total_pay  = COALESCE(?, total_pay),
    total_pres = COALESCE(?, total_pres)
WHERE id = ?;