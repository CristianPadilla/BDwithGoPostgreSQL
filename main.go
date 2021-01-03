package main

import (
	"fmt"
	"log"

	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/product"
	"github.com/CristianPadilla/BDwithGoPostgreSQL/storage"
)

func main() {
	// //PARA CREAR ULA TABLA PRODUCTOS
	// //crear conexion con la bd
	// storage.NewPostgresDB()
	// //inicializar sistema de almacenamiento de postgres para la tabla productos (NewPsqlProduct)
	// //este es un objeto que tiene los metodos de almacenamiento (Crud y otros) y por lo tanto
	// //implementa de la interface product.storage
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// //creamos un servicio que consume el sistema de almacenamiento anteriormente creado
	// //este servicio es el que ejecuta los metodos crud mediante el objeto storageProduct
	// productService := product.NewService(storageProduct)
	// //este servicio llama a el metodo que necesite como migrar(), create() etc
	// if err := productService.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate(): %v", err)
	// }

	//CREATE
	//instanciar conexion a la bd
	storage.NewPostgresDB()
	//crear sistema de almacenamiento de la tabla productos con la bd
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	//crear servicio con el sistema de almacenamiento anterior
	productService := product.NewService(storageProduct)
	// ejecutar la creacion
	m := &product.Model{
		Name:         "Curso de BD con Go",
		Price:        60,
		Observations: "bonito curso",
	}
	if err := productService.Create(m); err != nil {
		log.Fatalf("product.MCreate(): %v", err)
	}
	fmt.Printf("%+v", m)

	// //TRAER TODOS LOS REGISTROS DE UNA TABLA
	// storage.NewPostgresDB()
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// productService := product.NewService(storageProduct)
	// ms, err := productService.GetAll()
	// if err != nil {
	// 	log.Fatalf("product.GetAll(): %v", err)
	// }
	// fmt.Println(ms)

	// //SELECT UN REGISTRO DE DE LA TABLA
	// storage.NewPostgresDB()
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// productService := product.NewService(storageProduct)
	// m, err := productService.GetByID(1)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("no se encontraron registros")
	// case err != nil:
	// 	log.Fatalf("product.GetByID(): %v", err)
	// default:
	// 	fmt.Println(m)
	// }

	// //UPDATE
	// storage.NewPostgresDB()
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// productService := product.NewService(storageProduct)
	// m := &product.Model{
	// 	ID:           13,
	// 	Name:         "Curso java",
	// 	Price:        70,
	// 	Observations: "holaa",
	// }
	// if err := productService.Update(m); err != nil {
	// 	log.Fatalf("product.Update(): %v", err)
	// }
	// fmt.Printf("%+v", m)

	// //DELETE UN REGISTRO DE DE LA TABLA
	// storage.NewPostgresDB()
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// productService := product.NewService(storageProduct)
	// if err := productService.Delete(1); err != nil {
	// 	log.Fatalf("product.Delete(): %v", err)
	// }

	// storage.NewPostgresDB()
	// storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	// storageInvoice := storage.NewPsqlInvoice(storage.Pool(), storageInvoiceHeader, storageInvoiceItem)
	// inv := &invoice.Model{
	// 	Header: &invoiceHeader.Model{Client: "carlos"},
	// 	Items: invoiceItem.Models{
	// 		&invoiceItem.Model{ProductID: 4},
	// 		&invoiceItem.Model{ProductID: 4},
	// 	},
	// }
	// invoiceService := invoice.NewService(storageInvoice)
	// if err := invoiceService.Create(inv); err != nil {
	// 	log.Fatalf("invoice.Create(): %v", err)
	// }

}
