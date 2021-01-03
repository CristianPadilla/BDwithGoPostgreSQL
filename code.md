    ## Migrar tabla de invoiceHeader (factura)
    ```go
     storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
     invoiceHeaderService := invoiceHeader.NewService(storageInvoiceHeader)
     if err := invoiceHeaderService.Migrate(); err != nil {
     	log.Fatalf("invoiceHeader.Migrate(): %v", err)
     }

```
    ## Migrar tabla de invoiceItem (productos de factura)
	 storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	 invoiceItemService := invoiceHeader.NewService(storageInvoiceItem)
	 if err := invoiceItemService.Migrate(); err != nil {
	    log.Fatalf("invoiceItem.Migrate(): %v", err)
	 }
```
