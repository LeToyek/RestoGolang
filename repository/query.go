package repository

const (
	queryInsertUser = `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			phone,
			created_at,
			updated_at,
			user_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	queryGetAllUsers = `SELECT * FROM users`

	queryGetOneUser = `SELECT * FROM users WHERE user_id=$1;`

	queryGetAcc = `SELECT password,user_id FROM users WHERE email=$1;`

	queryInputTable = `
		INSERT INTO table_resto(
			chairs,
			isused,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4)
	`
	queryGetAllTables = `
		SELECT * FROM table_resto
	`

	queryInputFood = `
		INSERT INTO foods(
			food_id,
			name,
			price,
			image,
			category,
			created_at,
			updated_at
		)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`

	queryGetAllFoods = `
		SELECT * FROM foods
	`

	queryGetFoodsByCategory = `
		SELECT * FROM foods WHERE category=$1;
	`
	queryInputOrder = `
		INSERT INTO order_user(
			order_id,
			order_date,
			created_at,
			updated_at,
			user_id,
			table_id
		) VALUES(
			$1,$2,$3,$4,$5,$6
		) 
	`
	queryGetOrder = `
		SELECT * FROM order_user WHERE user_id=$1;
	`

	queryAddInvoice = `
		INSERT INTO invoice(
			id,
			order_id,
			pay_date,
			total
		) VALUES ($1,$2,$3,$4)
	`

	queryGetInvoice = `
	SELECT
	i.id, i.order_id, i.pay_date, i.total, o.table_id
	FROM invoice as i
	JOIN order_user AS o ON(i.order_id = o.order_id)
	WHERE i.order_id=$1
	`

	queryGetTotalPriceOrder = `
		SELECT
		fo.price*f.total_count AS total_price
	
		FROM order_user AS o
		JOIN order_food AS f ON(f.order_id = o.order_id)
		JOIN foods AS fo ON(fo.food_id = f.food_id)
		WHERE o.order_id=$1;
	`

	queryGetMoreDetailedInvoice = `
		SELECT
		fo.name,f.total_count,fo.price,
		fo.price*f.total_count AS total_price
		
		FROM order_user AS o
		JOIN order_food AS f ON(f.order_id = o.order_id)
		JOIN foods AS fo ON(fo.food_id = f.food_id)
		WHERE o.order_id=$1;
	`

	queryGetUserOrderID = `
		SELECT order_id FROM order_user WHERE user_id=$1
	`
)
