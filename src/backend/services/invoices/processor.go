package invoices

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/danmurphy1217/invoice-generator/models"
	"github.com/danmurphy1217/invoice-generator/repo"
	"gorm.io/gorm"
)

type Processor struct {
    Title       string
    Description string
    Price       string
    ImageURL string
    SupportEmail string
}

var _ InvoiceProcessor = &Processor{}

func (p *Processor) Build(ctx context.Context) ([]byte, error) {
    html, err := p.generateHTML()
    if err != nil {
        return []byte{}, fmt.Errorf("failed to generate HTML for PDF: %w", err)
    }

    pdfBytes, err := p.convertHTMLToPDF(html)
    if err != nil {
		return []byte{}, fmt.Errorf("failed to convert HTML to PDF: %w", err)
    }

	return pdfBytes, nil
}

func (p *Processor) Store(ctx context.Context, pdf []byte, title string) error {
    invoiceDBId, err := repo.FindByTitle(title)
    if err != nil {

        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return fmt.Errorf("failed to find invoice by title: %v", err)
        }

        // if error is record not found, noop
        return nil
    }

    // noop if found title but no error was returned
    if invoiceDBId != "" {
        return nil
    }

    // otherwise, we didn't find the title and should continue
    log.Println("did not find matching title in db, continuing")
    invoiceDTO := &models.Invoice{Title: title, PDFBytes: pdf}
    // Insert the user into the database
    result, err := repo.Insert(invoiceDTO)
    if err != nil {
        return fmt.Errorf("failed to save invoice to database: %v", err)
    }
    log.Printf("created invoice with ID %s", result)

    return nil
}

func (p *Processor) generateHTML() (string, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return "", err
    }
    templatePath := filepath.Join(cwd, "services/invoices/templates", "invoice.html")

    // Parse the HTML template file
    tpl, err := template.ParseFiles(templatePath)
    if err != nil {
        return "", err
    }

    var buf bytes.Buffer
    err = tpl.Execute(&buf, p)
    if err != nil {
        return "", err
    }

    return buf.String(), nil
}

func (p *Processor) convertHTMLToPDF(html string) ([]byte, error) {
    cmd := exec.Command("weasyprint", "-", "-")
    cmd.Stdin = bytes.NewReader([]byte(html))

    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        return nil, err
    }

    return out.Bytes(), nil
}