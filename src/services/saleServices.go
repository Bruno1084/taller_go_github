package services

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"go_test/src/db"
	"go_test/src/domain"
	"math/rand"
	"strings"
	"time"
)

func SearchSales(userID, status string) ([]domain.Sale, map[string]interface{}, error) {
	var rows *sql.Rows
	var err error

	if status != "" {
		rows, err = db.DB.Query(
			`SELECT id, userId, amount, status, createdAt, updatedAt, version FROM sales WHERE userId = ? AND status = ?`,
			userID, status)
	} else {
		rows, err = db.DB.Query(
			`SELECT id, userId, amount, status, createdAt, updatedAt, version FROM sales WHERE userId = ?`,
			userID)
	}

	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var sales []domain.Sale
	// Metadata
	var quantity, approved, rejected, pending int
	var totalAmount float32

	for rows.Next() {
		var sale domain.Sale
		err := rows.Scan(&sale.ID, &sale.UserID, &sale.Amount, &sale.Status, &sale.CreatedAt, &sale.UpdatedAt, &sale.Version)
		if err != nil {
			return nil, nil, err
		}

		// Contar en metadata
		quantity++
		totalAmount += sale.Amount

		switch sale.Status {
		case "approved":
			approved++
		case "rejected":
			rejected++
		case "pending":
			pending++
		}

		sales = append(sales, sale)
	}

	metadata := map[string]interface{}{
		"quantity":     quantity,
		"approved":     approved,
		"rejected":     rejected,
		"pending":      pending,
		"total_amount": totalAmount,
	}

	return sales, metadata, nil
}

func GetAllSales() ([]domain.Sale, error) {
	rows, err := db.DB.Query("SELECT id, userId, amount, status, version FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []domain.Sale

	for rows.Next() {
		var sale domain.Sale
		err := rows.Scan(&sale.ID, &sale.UserID, &sale.Amount, &sale.Status, &sale.Version)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func GetSaleById(id string) (domain.Sale, error) {
	var sale domain.Sale
	row := db.DB.QueryRow("SELECT id, userId, amount, status, version FROM sales WHERE id = ?", id)
	err := row.Scan(&sale.ID, &sale.UserID, &sale.Amount, &sale.Status, &sale.Version)

	return sale, err
}

func CreateSale(sale *domain.Sale) (*domain.Sale, error) {
	// Validar que user_id exista
	_, err := GetUserById(sale.UserID)
	if err != nil {
		return nil, err
	}

	// Validar que amount no sea cero
	if sale.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0 - createSale")
	}

	// Asignar estado aleatorio
	//Rand es obsoleto
	statuses := []string{"pending", "approved", "rejected"}
	rand.Seed(time.Now().UnixNano())
	sale.Status = statuses[rand.Intn(len(statuses))]

	// Setear campos automáticos
	sale.ID = uuid.NewString()
	sale.CreatedAt = time.Now()
	sale.UpdatedAt = time.Now()
	sale.Version = 1

	// Insertar en la base de datos
	query := `INSERT INTO sales(id, userId, amount, status, createdAt, updatedAt, version)
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err = db.DB.Exec(query, sale.ID, sale.UserID, sale.Amount, sale.Status, sale.CreatedAt, sale.UpdatedAt, sale.Version)
	if err != nil {
		return nil, err
	}

	return sale, nil
}

func PatchSale(id string, sale *domain.SaleUpdateFields) (*domain.Sale, error) {
	existingSale, err := GetSaleById(id)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(existingSale.Status) != "pending" {
		return nil, errors.New("solo se pueden actualizar ventas en estado 'pending'")
	}

	if sale.Status == nil || (strings.ToLower(*sale.Status) != "approved" && strings.ToLower(*sale.Status) != "rejected") {
		return nil, errors.New("el nuevo estado debe ser 'approved' o 'rejected'")
	}

	now := time.Now()
	existingSale.Status = *sale.Status
	existingSale.UpdatedAt = now
	existingSale.Version++

	query := `UPDATE sales SET status = ?, updatedAt = ?, version = ? WHERE id = ?`
	_, err = db.DB.Exec(query, existingSale.Status, existingSale.UpdatedAt, existingSale.Version, existingSale.ID)
	if err != nil {
		return nil, err
	}

	return &existingSale, nil
}
