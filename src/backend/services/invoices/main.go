package invoices

import "context"

type InvoiceProcessor interface {
	/*
	interface for processing invoices

	serves as a basic abstraction for building
	invoices and storing them in S3 and our DB
	*/

	// Build method builds the underlying Invoice
	Build(ctx context.Context) ([]byte, error)
	// Store method stores the underlying Invoice in db, s3, or both
	Store(ctx context.Context, pdf []byte, title string) error
}