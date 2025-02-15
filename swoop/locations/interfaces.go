package locations

// LocationAPI interface for locations
type LocationAPI interface {
	CreateLocation(name, desc string) Location

	GetLocation(locationID string) Location
	GetLocationByCode(code string) Location
	SearchLocations(term string) []Location

	EnableLocation(locationID string) Location
	DisableLocation(locationID string) Location

	ActivateLocation(locationID string) Location
	DeactivateLocation(locationID string) Location
}

// MenuAPI interface for menus
type MenuAPI interface {
	CreateMenu(name, desc string, locationID uint) Menu

	GetMenu(menuID uint) Menu
	GetMenus(locationID uint) []Menu

	EnableMenu(menuID string) Menu
	DisableMenu(menuID string) Menu
}

// CategoryAPI interface for categories
type CategoryAPI interface {
	CreateCategory(name, desc string, menuID uint) Category

	GetCategories(menuID uint, timezone string) []Category
	GetCategory(categoryID uint) Category

	EnableCategory(categoryID uint) Category
	DisableCategory(categoryID uint) Category
}

// ItemAPI interface for items
type ItemAPI interface {
	CreateItem(name, desc string, price float32, categoryID uint) Item

	GetItem(itemID uint) Item
	GetItems(itemIds []uint) []Item
	GetItemsByCategoryID(categoryID uint) []Item
}

// OptionAPI interface for item options
type OptionAPI interface {
	CreateOption(name, desc string, min, max int, itemID uint, items []OptionItem) Option

	GetOption(optionID uint) Option
	GetOptions(itemID uint) []Option

	GetOptionItem(itemID uint) OptionItem
	GetOptionItems(optionID uint) []OptionItem
}

// MArker interface for location markers
type MarkerAPI interface {
	GetMarker(markerID uint) (Marker, error)
	CreateMarker(name string) (Marker, error)
	DeleteMarker(markerID uint) error
	UpdateMarker(marker Marker) error
	EnableMarker(markerID uint) error
	DisableMarker(markerID uint) error
	GetMarkersByLocation(locationID uint) ([]Marker, error)
}
