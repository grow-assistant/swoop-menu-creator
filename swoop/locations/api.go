package locations

import (
	"time"

	"swoop/pkg/db"
)

// API interface for location operations
type API interface {
	LocationAPI
	MenuAPI
	CategoryAPI
	ItemAPI
	OptionAPI
	MarkerAPI
}

type locationAPI struct {
	db *db.Database
}

// NewAPI initializes new instance of location API
func NewAPI(db *db.Database) API {
	return locationAPI{
		db: db,
	}
}

// GetLocation returns a Location by id
func (a locationAPI) GetLocation(locationID string) Location {
	var location Location
	a.db.DB.
		Preload("Markers").
		First(&location, locationID)
	return location
}

func (a locationAPI) GetLocationByCode(code string) Location {

	var location Location
	a.db.DB.
		Preload("Markers").
		Where("code = ?", code).
		First(&location)
	return location
}

// CreateLocation creates a new Location model
func (a locationAPI) CreateLocation(name, descr string) Location {
	location := Location{
		Name:        name,
		Description: descr,
		Disabled:    false,
		Active:      true,
	}
	a.db.DB.Create(&location)
	return location
}

func (a locationAPI) SearchLocations(term string) []Location {
	var locations []Location
	//a.db.DB.Select("*, LOWER(name)").Where("name ILIKE ?", "%" + search + "%").FindByUID(&locations)
	// TODO parameterize this eventually
	a.db.DB.
		Preload("Markers").
		Where("disabled = false AND active = true").
		Order("name asc").
		Find(&locations)
	return locations
}

func (a locationAPI) EnableLocation(locationID string) Location {
	panic("not implemented")
}

func (a locationAPI) DisableLocation(locationID string) Location {
	panic("not implemented")
}

func (a locationAPI) ActivateLocation(locationID string) Location {
	panic("not implemented")
}

func (a locationAPI) DeactivateLocation(locationID string) Location {
	panic("not implemented")
}

// CreateMenu creates a new Menu model
func (a locationAPI) CreateMenu(name, descr string, locationID uint) Menu {
	menu := Menu{
		Name:        name,
		Description: descr,
		LocationID:  locationID,
	}
	a.db.DB.Create(&menu)
	return menu
}

// GetMenu returns a Menu by id
func (a locationAPI) GetMenu(menuID uint) Menu {
	var menu Menu
	a.db.DB.Find(&menu, menuID)
	return menu
}

// GetMenus returns menus for a Location
func (a locationAPI) GetMenus(locationID uint) []Menu {
	var menus []Menu
	a.db.DB.Where(&Menu{LocationID: locationID}).Order("name").Find(&menus)
	return menus
}

// EnableMenu enables a Menu
func (a locationAPI) EnableMenu(menuID string) Menu {
	var menu Menu
	return menu
}

// Disable menu disables a Menu
func (a locationAPI) DisableMenu(menuID string) Menu {
	var menu Menu
	return menu
}

// CreateCategory creates a new Category model
func (a locationAPI) CreateCategory(name, desc string, menuID uint) Category {
	category := Category{
		Name:        name,
		Description: desc,
		MenuID:      menuID,
	}
	a.db.DB.Create(&category)
	return category
}

// GetCategories returns categories for a menu
func (a locationAPI) GetCategories(menuID uint, timezone string) []Category {
	var categories []Category
	a.db.DB.Where(&Category{MenuID: menuID}).Order("name").Find(&categories)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil
	}

	// convert the time to the local time of the location
	// and normalize the time to an integer consisting of hours + minutes
	// so that it can be compared with the value from the database
	// i.e. 2:30PM = (14 hours* 100) + 30 minutes = 1430
	localTime := time.Now().In(location)
	normalizedLocalTime := (localTime.Hour() * 100) + localTime.Minute()

	// if timezone is provided, filter out categories that do not fall within time range.
	// if no timezone is provided, then include in the results
	var results []Category
	for _, c := range categories {
		withinTimeRange := normalizedLocalTime >= c.StartTime && normalizedLocalTime < c.EndTime
		hasTimeRange := c.StartTime > 0 && c.EndTime > 0
		if withinTimeRange || !hasTimeRange {
			results = append(results, c)
		}
	}
	return results
}

// GetCategory returns a Category by id
func (a locationAPI) GetCategory(categoryID uint) Category {
	var category Category
	a.db.DB.Find(&category, categoryID)
	return category
}

// EnableCategory enables a Category
func (a locationAPI) EnableCategory(categoryID uint) Category {
	var category Category
	return category
}

// DisableCategory disables a Category
func (a locationAPI) DisableCategory(categoryID uint) Category {
	var category Category
	return category
}

// CreateItem creates a new Item for a given Category
func (a locationAPI) CreateItem(name, desc string, price float32, categoryID uint) Item {
	item := Item{
		Name:        name,
		Description: desc,
		Price:       price,
		CategoryID:  categoryID,
	}
	a.db.DB.Create(&item)
	return item
}

// GetItem returns an Item by ID
func (a locationAPI) GetItem(itemID uint) Item {
	var item Item
	a.db.DB.Find(&item, itemID)
	return item
}

// GetItems returns items matching list of provided item IDs
func (a locationAPI) GetItems(itemIds []uint) []Item {
	var items []Item
	a.db.DB.Where(itemIds).Order("name").Find(&items)
	return items
}

// GetItemsByCategoryID returns items for a given Category
func (a locationAPI) GetItemsByCategoryID(categoryID uint) []Item {
	var items []Item
	a.db.DB.Where(&Item{CategoryID: categoryID}).Order("name").Find(&items)
	return items
}

// CreateOption creates a new option
func (a locationAPI) CreateOption(name, desc string, min, max int, itemID uint, items []OptionItem) Option {
	option := Option{
		Name:        name,
		Description: desc,
		Min:         min,
		Max:         max,
		ItemID:      itemID,
	}

	a.db.DB.Create(&option)

	for _, item := range items {
		item.OptionID = option.ID
		a.db.DB.Create(&item)
	}

	option.Items = items
	return option
}

// GetOption returns an Option by ID
func (a locationAPI) GetOption(optionID uint) Option {
	var option Option
	a.db.DB.Find(&option, optionID)
	return option
}

// GetOptions returns options for an Item
func (a locationAPI) GetOptions(itemID uint) []Option {
	// TODO is there a more idiomatic way to do this?
	var options []Option
	a.db.DB.Where(map[string]interface{}{"item_id": itemID}).Order("name").Find(&options)
	return options
}

// GetOptionItem returns an OptionItem by ID
func (a locationAPI) GetOptionItem(itemID uint) OptionItem {
	var optionItem OptionItem
	a.db.DB.Find(&optionItem, itemID)
	return optionItem
}

// GetOptionItems returns OptionItems for a given Option
func (a locationAPI) GetOptionItems(optionID uint) []OptionItem {
	var optionItems []OptionItem
	a.db.DB.Where(&OptionItem{OptionID: optionID}).Order("name").Find(&optionItems)
	return optionItems
}

func (a locationAPI) GetMarker(markerID uint) (Marker, error) {
	var marker Marker
	err := a.db.DB.Find(&marker, markerID).Error
	return marker, err
}

func (a locationAPI) CreateMarker(name string) (Marker, error) {
	marker := Marker{Name: name}
	err := a.db.DB.Create(&marker).Error
	return marker, err
}

func (a locationAPI) DeleteMarker(markerID uint) error {
	var marker Marker
	err := a.db.DB.Delete(&marker, "id = ?", markerID).Error
	return err
}

func (a locationAPI) UpdateMarker(marker Marker) error {
	panic("not implemented")
}

func (a locationAPI) EnableMarker(markerID uint) error {
	panic("not implemented")
}

func (a locationAPI) DisableMarker(markerID uint) error {
	panic("not implemented")
}

func (a locationAPI) GetMarkersByLocation(locationID uint) ([]Marker, error) {
	var markers []Marker
	err := a.db.DB.Where(&Marker{LocationID: locationID}).Order("name").Find(&markers).Error
	return markers, err
}
