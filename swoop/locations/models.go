package locations

import (
	"github.com/jinzhu/gorm"
)

// Location data model
type Location struct {
	gorm.Model
	Name        string
	Description string
	Menus       []Menu
	Markers     []Marker
	Timezone    string  `json:"timezone" gorm:"column:timezone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Disabled    bool    `json:"disabled" gorm:"column:disabled"`
	Active      bool    `json:"active" gorm:"column:active"`
	Code        string  `json:"code" gorm:"column:code"`
	TaxRate     float32 `json:"tax_rate" gorm:"column:tax_rate"`
}

// Menu data model
type Menu struct {
	gorm.Model
	Name        string
	Description string
	LocationID  uint `json:"location_id" gorm:"column:location_id"`
	Location    Location
	Categories  []Category
	Disabled    bool
}

// Category data model
type Category struct {
	gorm.Model
	Name        string
	Description string
	MenuID      uint `json:"menu_id" gorm:"column:menu_id"`
	Menu        Menu
	Disabled    bool
	Items       []Item
	StartTime   int `json:"start_time" gorm:"column:start_time"`
	EndTime     int `json:"end_time" gorm:"column:end_time"`
}

// Item data model
type Item struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
	Disabled    bool
	CategoryID  uint `json:"category_id gorm:column:category_id"`
	Category    Category
	Options     []Option
}

// Option data model
type Option struct {
	gorm.Model
	Name        string
	Description string
	Min         int
	Max         int
	Disabled    bool
	ItemID      uint `json:"item_id gorm:column:item_id"`
	Item        Item
	Items       []OptionItem
}

// OptionItem data model
type OptionItem struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
	Disabled    bool
	OptionID    uint `json:"option_id gorm:column:option_id"`
	Option      Option
}

// Marker data model
type Marker struct {
	gorm.Model
	Name       string
	LocationID uint `json:"location_id" gorm:"column:location_id"`
	Location   Location
	Disabled   bool
}
