package custom_errors

const ErrorNotFound = "truck not found on garage"

// TruckNotFoundError is a custom error used when a listing is not properly
// returned from the Garage External API
type TruckNotFoundError struct {
    Message string
}

func (e *TruckNotFoundError) Error() string {
    return e.Message
}

func NewTruckNotFoundError() error {
    return &TruckNotFoundError{Message: ErrorNotFound}
}