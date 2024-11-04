package supplier

const (
	sqlInsertSupplier = `
		INSERT INTO suppliers (name, leadtime_max, leadtime_avg) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`

	sqlGetSuppliers = `
		SELECT id, name, leadtime_max, leadtime_avg
		FROM suppliers
		ORDER BY id DESC;
	`

	sqlGetSupplierByID = `
		SELECT id, name, leadtime_max, leadtime_avg
		FROM suppliers
		WHERE id = $1;
	`
)
