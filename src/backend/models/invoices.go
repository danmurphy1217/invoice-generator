package models

// Invoice is our internal data-transportation representation of an Invoice
type Invoice struct {
    Title string
    PDFBytes []byte
}

// the struct we use for unmarshaling a listing returned from Garage
type Listing struct {
    ID                string   `json:"id"`
    CreatedAt         string   `json:"createdAt"`
    UpdatedAt         string   `json:"updatedAt"`
    ItemBrand         string   `json:"itemBrand"`
    ItemCategory      int      `json:"itemCategory"`
    ListingTitle      string   `json:"listingTitle"`
    ListingDescription string  `json:"listingDescription"`
    SellingPrice      float32      `json:"sellingPrice"`
    IsShippable       bool     `json:"isShippable"`
    ImageUrls         []string `json:"imageUrls"`
    ListingStatus     int      `json:"listingStatus"`
    Tags              []string `json:"tags"`
    Categories        []int    `json:"categories"`
    ItemWears         []string `json:"itemWears"`
    ScrapedLink       *string  `json:"scrapedLink"`
    ItemAge           int      `json:"itemAge"`
    ItemCondition     string   `json:"itemCondition"`
    ItemLength        *int     `json:"itemLength"`
    ItemWidth         *int     `json:"itemWidth"`
    ItemHeight        *int     `json:"itemHeight"`
    ItemWeight        int      `json:"itemWeight"`
    AddressPrimary    string   `json:"addressPrimary"`
    AddressSecondary  string   `json:"addressSecondary"`
    AddressCity       string   `json:"addressCity"`
    AddressState      string   `json:"addressState"`
    AddressZip        string   `json:"addressZip"`
    UserID            string   `json:"userId"`
}

// struct for the result
type Result struct {
    Listing *Listing `json:"listing"`
}

// struct for the full response
type GetListingResponse struct {
    Result *Result `json:"result"`
    Error  string `json:"error"`
}