UPDATE workshops
SET   total_pres = COALESCE(?, total_pres)
WHERE id = ?;