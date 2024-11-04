package transaction

const (
	sqlInsertTransaction = `
		INSERT INTO transactions (user_id, product_id, total_price, quantity)
		VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
)
