package model

import (
	"encoding/json"
	"fmt"
)

const (
	orderFilename = "order_%d.md"
)

type Order struct {
	Base
	Amount uint64 `json:"amount"`
	// Relationships
	// foreign key
	DrinkID uint  `json:"drink_id" gorm:"not null"`
	Drink   Drink `json:"drink"`
}

func (o *Order) ToMarkdown() string {
	// Format the CreatedAt timestamp in your preferred layout
	createdAt := o.CreatedAt.Format("Jan 02 15:04:05") // e.g., "Nov 12 17:12:39"

	// Generate markdown table
	markdown := fmt.Sprintf(`# Order: %d

| Created At      | Drink ID | Amount |
|-----------------|----------|--------|
| %s | %d        | %d      |

Thanks for drinking with us!
`, o.ID, createdAt, o.DrinkID, o.Amount)

	return markdown
}

// ToJSON returns a JSON representation of the order. This is used when we write
// the order "markdown" file to storage but want the file to contain only JSON.
func (o *Order) ToJSON() string {
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(b)
}

func (o *Order) GetFilename() string {
	return fmt.Sprintf(orderFilename, o.ID)
}
