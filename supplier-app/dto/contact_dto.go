package dto

type ContactRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	JobPosition string `json:"job_position"`
	Mobile      string `json:"mobile"`
	SupplierID  uint   `json:"supplier_id"`
}
