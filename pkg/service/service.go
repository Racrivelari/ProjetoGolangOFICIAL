package service

import (
	"fmt"
	"log"
	"github.com/faelp22/tcs_curso/stoq/entity"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
)

type ProdutoServiceInterface interface {
	GetAll() *entity.ProdutoList
	GetByID(ID *int64) *entity.Product
	Create(produto *entity.Product) int64
	Update(ID *int64, produto *entity.Product) int64
	Delete(ID *int64) int64
}

type produto_service struct {
	dbp database.DatabaseInterface
}

func NewProdutoService(dabase_pool database.DatabaseInterface) *produto_service {
	return &produto_service{
		dabase_pool,
	}
}

func (ps *produto_service) GetAll() *entity.ProdutoList {
	DB := ps.dbp.GetDB()

	rows, err := DB.Query("SELECT id_prod, name_prod, price_prod, code_prod FROM Product")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	lista_produtos := &entity.ProdutoList{}

	for rows.Next() {
		p := entity.Product{}

		if err := rows.Scan(&p.ID, &p.Name, &p.Code, &p.Price); err != nil {
			fmt.Println(err.Error())
		} else {
			lista_produtos.List = append(lista_produtos.List, &p)
		}
	}

	return lista_produtos
}

func (ps *produto_service) GetByID(ID *int64) *entity.Product {
	
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("SELECT id_prod, name_prod, price_prod, code_prod FROM Product where id_prod = ?")
	// stmt, err := DB.Prepare("SELECT id_prod, name_prod, price_prod, code_prod FROM Product where id_prod = ?", ID)
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	p := entity.Product{}

	err = stmt.QueryRow(ID).Scan(&p.ID, &p.Name, &p.Code, &p.Price)
	if err != nil {
		log.Println(err.Error())
	}
	return &p
}

func (ps *produto_service) Create(produto *entity.Product) int64 {
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("INSERT INTO Product (name_prod, price_prod, code_prod) values (?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price)
	if err != nil {
		log.Println(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	return lastId
}

func (ps *produto_service) Update(ID *int64, produto *entity.Product) int64 {
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("UPDATE Product SET name_prod = ?, price_prod = ? where id_prod = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}

func (ps *produto_service) Delete(ID *int64) int64 {
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("DELETE FROM Logs where id_prod = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	stmt.Exec(ID)

	stmt2, err := DB.Prepare("DELETE FROM Product where id_prod = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt2.Close()

	result2, err := stmt2.Exec(ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result2.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}
