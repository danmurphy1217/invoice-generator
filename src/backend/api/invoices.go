package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/danmurphy1217/invoice-generator/constants"
	"github.com/danmurphy1217/invoice-generator/custom_errors"
	"github.com/danmurphy1217/invoice-generator/gen"
	"github.com/danmurphy1217/invoice-generator/providers"
	"github.com/danmurphy1217/invoice-generator/services/invoices"
	"github.com/danmurphy1217/invoice-generator/utils"
)

func GenerateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
    /**
    handler func to generate an invoice.

    This function (1) uses our Garage "provider" to get the
    listing from the garage API, (2) parses that listing
    into our internal "Listing" data object, (3) stores
    information on that listing in our database and returns
    the bytes of that item
    */
    var req gen.GenerateInvoiceHandlerRequest

    // Decode the request body into the struct
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    truckId := req.TruckId
    if truckId == "" {
        http.Error(
            w,
            fmt.Errorf("invalid request, must include truck ID in request").Error(),
            http.StatusBadRequest,
        )
        return
    }
    
    garageProvider := providers.Garage{
        Timeout: constants.DefaultTimeout,
        Headers: map[string]string{},
        Endpoint: "getListing",
        TruckId: truckId,
    }
    respGetListing, err := garageProvider.GetListing()
    if err != nil {
        var truckNotFoundErr *custom_errors.TruckNotFoundError
        if errors.As(err, &truckNotFoundErr) {
            // if truck not found, return 404
            http.Error(
                w,
                err.Error(),
                http.StatusNotFound,
            )
            return
        }

        // otherwise 500
        http.Error(
            w,
            fmt.Errorf("failed to get listing from Garage: %w", err).Error(),
            http.StatusInternalServerError,
        )
        return
    }

    imageUrl := getImageUrl(respGetListing.Result.Listing.ImageUrls)
    invoiceProcessor := &invoices.Processor{
        Title: respGetListing.Result.Listing.ListingTitle,
		Description: respGetListing.Result.Listing.ListingDescription,
		Price: utils.FormatPrice(respGetListing.Result.Listing.SellingPrice),
        ImageURL: imageUrl,
        SupportEmail: "support@withgarage.com",
    }
    // build the invoice
    invoiceBytes, err := invoiceProcessor.Build(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // upload the invoice to our DB
    err = invoiceProcessor.Store(r.Context(), invoiceBytes, respGetListing.Result.Listing.ListingTitle)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    w.WriteHeader(http.StatusCreated)
    w.Write(invoiceBytes)
}

func getImageUrl(imageUrls []string) string {
    response := ""
    if len(imageUrls) > 0 {
        response = imageUrls[0]
    } 
    return response
}