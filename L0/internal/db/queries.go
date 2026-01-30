package db

import (
	"database/sql"
	"log"
	"microservice/internal/models"
)

// GetOrderByUID возвращает заказ по order_uid
func GetOrderByUID(db *sql.DB, orderUID string) (*models.Order, error) {
	order := &models.Order{}

	var entry, locale, internalSig, customerID, deliveryService, shardKey, oofShard sql.NullString
	var smID sql.NullInt64
	var dateCreated sql.NullTime

	// Таблица orders
	err := db.QueryRow(`
        SELECT order_uid, track_number,
               entry, locale, internal_signature,
               customer_id, delivery_service, shardkey, sm_id,
               date_created, oof_shard
        FROM orders WHERE order_uid = $1`, orderUID).
		Scan(&order.OrderUID, &order.TrackNumber, &entry, &locale, &internalSig,
			&customerID, &deliveryService, &shardKey, &smID, &dateCreated, &oofShard)
	if err != nil {
		return nil, err
	}

	order.Entry = entry.String
	order.Locale = locale.String
	order.InternalSig = internalSig.String
	order.DeliveryService = deliveryService.String
	order.ShardKey = shardKey.String
	order.OofShard = oofShard.String
	if customerID.Valid {
		order.CustomerID = customerID.String
	}
	if smID.Valid {
		order.SmID = int(smID.Int64)
	}
	if dateCreated.Valid {
		order.DateCreated = dateCreated.Time.Format("2006-01-02 15:04:05")
	}

	// Таблица delivery
	var name, phone, zip, city, address, region, email sql.NullString
	err = db.QueryRow(`
        SELECT name, phone, zip, city, address, region, email
        FROM delivery WHERE order_uid = $1`, orderUID).
		Scan(&name, &phone, &zip, &city, &address, &region, &email)
	if err != nil {
		return nil, err
	}
	order.Delivery = models.Delivery{
		Name:    name.String,
		Phone:   phone.String,
		Zip:     zip.String,
		City:    city.String,
		Address: address.String,
		Region:  region.String,
		Email:   email.String,
	}

	// Таблица payment
	var transaction, requestID, currency, provider, bank sql.NullString
	var amount, deliveryCost, goodsTotal, customFee sql.NullInt64
	var paymentDt sql.NullInt64
	err = db.QueryRow(`
        SELECT transaction, request_id, currency, provider,
               amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
        FROM payment WHERE order_uid = $1`, orderUID).
		Scan(&transaction, &requestID, &currency, &provider,
			&amount, &paymentDt, &bank, &deliveryCost, &goodsTotal, &customFee)
	if err != nil {
		return nil, err
	}
	order.Payment = models.Payment{
		Transaction:  transaction.String,
		RequestID:    requestID.String,
		Currency:     currency.String,
		Provider:     provider.String,
		Amount:       int(amount.Int64),
		PaymentDt:    paymentDt.Int64,
		Bank:         bank.String,
		DeliveryCost: int(deliveryCost.Int64),
		GoodsTotal:   int(goodsTotal.Int64),
		CustomFee:    int(customFee.Int64),
	}

	// Таблица items
	rows, err := db.Query(`
        SELECT chrt_id, track_number, price, rid, name, sale,
               size, total_price, nm_id, brand, status
        FROM items WHERE order_uid = $1`, orderUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chrtID, price, sale, totalPrice, nmID, status sql.NullInt64
		var trackNumber, rid, name, size, brand sql.NullString

		if err := rows.Scan(&chrtID, &trackNumber, &price, &rid, &name, &sale,
			&size, &totalPrice, &nmID, &brand, &status); err != nil {
			return nil, err
		}

		item := models.Item{
			ChrtID:      int(chrtID.Int64),
			TrackNumber: trackNumber.String,
			Price:       int(price.Int64),
			Rid:         rid.String,
			Name:        name.String,
			Sale:        int(sale.Int64),
			Size:        size.String,
			TotalPrice:  int(totalPrice.Int64),
			NmID:        int(nmID.Int64),
			Brand:       brand.String,
			Status:      int(status.Int64),
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

// InsertOrder вставляет заказ или обновляет (UPSERT)
func InsertOrder(db *sql.DB, order *models.Order) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,
	                    customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	ON CONFLICT (order_uid) DO UPDATE
	SET track_number = EXCLUDED.track_number,
	    entry = EXCLUDED.entry,
	    locale = EXCLUDED.locale,
	    internal_signature = EXCLUDED.internal_signature,
	    customer_id = EXCLUDED.customer_id,
	    delivery_service = EXCLUDED.delivery_service,
	    shardkey = EXCLUDED.shardkey,
	    sm_id = EXCLUDED.sm_id,
	    date_created = EXCLUDED.date_created,
	    oof_shard = EXCLUDED.oof_shard
	`, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSig,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID,
		order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	ON CONFLICT (order_uid) DO UPDATE
	SET name = EXCLUDED.name,
	    phone = EXCLUDED.phone,
	    zip = EXCLUDED.zip,
	    city = EXCLUDED.city,
	    address = EXCLUDED.address,
	    region = EXCLUDED.region,
	    email = EXCLUDED.email
	`, order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount,
	                     payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	ON CONFLICT (order_uid) DO UPDATE
	SET transaction = EXCLUDED.transaction,
	    request_id = EXCLUDED.request_id,
	    currency = EXCLUDED.currency,
	    provider = EXCLUDED.provider,
	    amount = EXCLUDED.amount,
	    payment_dt = EXCLUDED.payment_dt,
	    bank = EXCLUDED.bank,
	    delivery_cost = EXCLUDED.delivery_cost,
	    goods_total = EXCLUDED.goods_total,
	    custom_fee = EXCLUDED.custom_fee
	`, order.OrderUID, order.Payment.Transaction, order.Payment.RequestID,
		order.Payment.Currency, order.Payment.Provider, order.Payment.Amount,
		order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost,
		order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		_, err = tx.Exec(`
		INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale,
		                   size, total_price, nm_id, brand, status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		ON CONFLICT (order_uid, chrt_id) DO UPDATE
		SET track_number = EXCLUDED.track_number,
		    price = EXCLUDED.price,
		    rid = EXCLUDED.rid,
		    name = EXCLUDED.name,
		    sale = EXCLUDED.sale,
		    size = EXCLUDED.size,
		    total_price = EXCLUDED.total_price,
		    nm_id = EXCLUDED.nm_id,
		    brand = EXCLUDED.brand,
		    status = EXCLUDED.status
		`, order.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid,
			item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID,
			item.Brand, item.Status)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetAllOrders возвращает все заказы
func GetAllOrders(db *sql.DB) ([]*models.Order, error) {
	rows, err := db.Query(`SELECT order_uid FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order

	for rows.Next() {
		var orderUID string
		if err := rows.Scan(&orderUID); err != nil {
			return nil, err
		}

		order, err := GetOrderByUID(db, orderUID)
		if err != nil {
			log.Printf("Ошибка чтения заказа %s из БД: %v", orderUID, err)
			continue
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
