package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

type Handler struct {
	DB *pgx.Conn
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(context.Background(), "SELECT id, name, price, description, stock FROM products")
	if err != nil {
		http.Error(w, "error querying", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Stock)
		products = append(products, p)
	}
	data, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "error parsing to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id_str := r.PathValue("id")
	product_id, err := strconv.Atoi(id_str)
	if err != nil {
		http.Error(w, "error parsing id", http.StatusInternalServerError)
		return
	}
	my_query := "SELECT id,name,price,description,stock FROM products WHERE id=$1"
	row := h.DB.QueryRow(context.Background(), my_query, product_id)
	var p Product
	err = row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Stock)
	if err == pgx.ErrNoRows {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(p) // filled Product with Scan()!
	if err != nil {
		http.Error(w, "error parsing to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
