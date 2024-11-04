package product

const (
	sqlInsertProduct = `
		INSERT INTO products (supplier_id, name, stock, price)
		VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`

	sqlGetProducts = `
		SELECT id, supplier_id, name, stock, price 
		FROM products 
		ORDER BY id DESC;
	`

	sqlGetProductByID = `
		SELECT id, supplier_id, name, stock, price 
		FROM products 
		WHERE id = $1
	`

	sqlUpdateProduct = `
		UPDATE products
		SET supplier_id = $1, name = $2, stock = $3, price = $4
		WHERE id = $5
	`

	sqlUpdateProductStock = `
		UPDATE products
		SET stock = $1
		WHERE id = $2
	`
)
