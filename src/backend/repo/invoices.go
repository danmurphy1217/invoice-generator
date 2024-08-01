package repo

import (
	"fmt"

	"github.com/danmurphy1217/invoice-generator/db"
	"github.com/danmurphy1217/invoice-generator/db/collections"
	"github.com/danmurphy1217/invoice-generator/models"
)

func Insert(invoiceDTO *models.Invoice) (string, error) {
	_, dbDriver := db.Connect()

    invoiceDB := collections.Invoice{
        Title: invoiceDTO.Title, 
        PDFBytes: invoiceDTO.PDFBytes,
    }
    // Insert the user into the database
    result := dbDriver.Create(&invoiceDB)
    if result.Error != nil {
        return "", fmt.Errorf("failed to save invoice to database: %v", result.Error)
    }
    // commit txn
    result.Commit()
    
    return invoiceDB.ID, nil
}

func FindByTitle(title string) (string, error) {
    /*
    find an invoice by title (which is used, naively, 
    as our unique constraint)
    */
	_, dbDriver := db.Connect()

    var invoiceDB collections.Invoice
    // find invoice with title
    result := dbDriver.Where("title = ?", title).First(&invoiceDB)
    if result.Error != nil {
        return "", result.Error
    }

    return invoiceDB.ID, nil
}