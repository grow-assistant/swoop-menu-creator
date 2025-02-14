package main

import (
	"log"

	"swoop/locations"
	"swoop/pkg/config"
	database "swoop/pkg/db"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Seeding Database")
	db := database.Connect(config.DB())
	api := locations.NewAPI(db)
	// seed locations
	log.Println("Seeding *Location* data")
	gatesFourGolfandCountryClub := api.CreateLocation("Gates Four Golf & Country Club", "Fayetteville, NC")
	log.Println("Seeding *Menus* data")
	// seed menus
	gatesFourGolfandCountryClubJPsBarandGrill := api.CreateMenu("JPs Bar and Grill", "JPs Bar and Grill", gatesFourGolfandCountryClub.ID)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrill)
	// seed categories
	log.Println("Seeding *Categories* data")
	gatesFourGolfandCountryClubJPsBarandGrillAppetizers := api.CreateCategory("Appetizers", "Appetizers", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads := api.CreateCategory("Soups & Salads", "Soups & Salads", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillSpecialties := api.CreateCategory("Specialties", "Specialties", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillSandwiches := api.CreateCategory("Sandwiches", "Sandwiches", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillBurgers := api.CreateCategory("Burgers", "Burgers", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillSides := api.CreateCategory("Sides", "Sides", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillAppetizers)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillSpecialties)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillSandwiches)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillBurgers)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillSides)
	// seed item
	log.Println("Seeding *Items* data")
	_ = api.CreateItem("Calamari", "A mix of tentacles and rings, lightly coated and fried, served with a side of cocktail.", 12, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	gatesFourGolfandCountryClubChickenWings := api.CreateItem("Chicken Wings", "Fried jumbo chicken wings, your choice of six or twelve, tossed in one of our signature sauces: Buffalo, BBQ, Blazing BBQ, Incinerator, Garlic Parm.", 7, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	gatesFourGolfandCountryClubClubNachos := api.CreateItem("Club Nachos", "Your choice of chicken or beef, atop warm corn tortilla chips with jalapenos, onions, and cheddar cheese.", 12, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	_ = api.CreateItem("Loaded Skins", "Six halves, loaded down with cheddar cheese, bacon, and green onion with a side of sour cream.", 10, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	_ = api.CreateItem("Mozzarella Sticks", "Six of our breaded mozzarella sticks, fried golden and served with marinara.", 8, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	_ = api.CreateItem("Par Three Platter", "A half portion of quesadilla, three mozzarella sticks, and six of our stuffed mushrooms, served with marinara, salsa, and Cajun aioli.", 13, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	_ = api.CreateItem("BBQ Pork Quesadilla", "A full 12\" quesadilla stuffed with smoked pork, pepper jack, onions, and jalapenos.", 12, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	_ = api.CreateItem("Loaded Quesadilla", "A full 12\" quesadilla loaded with peppers, onions, and chicken. Served with salsa and sour cream.", 11, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	_ = api.CreateItem("Stuffed Mushrooms", "A plate of our baked button mushrooms stuffed with our blend of seasonings and cream cheese.", 8, gatesFourGolfandCountryClubJPsBarandGrillAppetizers.ID)
	_ = api.CreateItem("Soup of the Day (Cup)", "Chef’s daily soup selection.", 4, gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads.ID)
	_ = api.CreateItem("Soup of the Day (Bowl)", "Chef’s daily soup selection.", 6, gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads.ID)
	_ = api.CreateItem("Side House or Caesar Salad", "Small portion of house or Caesar salad.", 5, gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads.ID)
	gatesFourGolfandCountryClubHouseSalad := api.CreateItem("House Salad", "Fresh mixed greens with tomato, cucumber, mushrooms, onions, and cheddar.", 10, gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads.ID)
	gatesFourGolfandCountryClubCaesarSalad := api.CreateItem("Caesar Salad", "Chopped romaine tossed with parmesan, croutons, and Caesar dressing.", 9, gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads.ID)
	gatesFourGolfandCountryClubClubSalad := api.CreateItem("Club Salad", "Fresh mixed greens with tomato, cucumber, mushrooms, onions, cheddar, smoked ham and turkey, and bacon.", 13, gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads.ID)
	gatesFourGolfandCountryClubGatesFourSpinachSalad := api.CreateItem("Gates Four Spinach Salad", "Baby leaf spinach tossed with candied nuts, craisins, bacon, onion, blue cheese crumbles, and balsamic dressing.", 14, gatesFourGolfandCountryClubJPsBarandGrillSoupsAndSalads.ID)
	gatesFourGolfandCountryClubAhiTunaSteak := api.CreateItem("Ahi Tuna Steak", "A yellowfin tuna steak crusted with sesame and lightly seared and sliced. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillSpecialties.ID)
	_ = api.CreateItem("Bang Bang Tempura", "Eight battered jumbo shrimp tossed in our sweet and sour sauce. Served with your choice of side.", 15, gatesFourGolfandCountryClubJPsBarandGrillSpecialties.ID)
	_ = api.CreateItem("Fish and Chips", "Two battered cod filets fried golden and served with an order of French fries and tartar sauce.", 16, gatesFourGolfandCountryClubJPsBarandGrillSpecialties.ID)
	gatesFourGolfandCountryClubChickenFajitas := api.CreateItem("Chicken Fajitas", "Two 6\" flour tortillas with grilled fajita chicken, grilled peppers and onions, and cotija cheese. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillSpecialties.ID)
	_ = api.CreateItem("Bacon Mac & Cheese", "Penne pasta tossed in our homemade bacon cheddar sauce.", 9, gatesFourGolfandCountryClubJPsBarandGrillSpecialties.ID)
	gatesFourGolfandCountryClubPortobelloBurger := api.CreateItem("Portobello Burger", "A portobello mushroom cap seasoned and grilled, served on a seeded gluten-free bun with lettuce and tomato. Served with your choice of side.", 12, gatesFourGolfandCountryClubJPsBarandGrillSpecialties.ID)
	gatesFourGolfandCountryClubIrongateDeli := api.CreateItem("Irongate Deli", "Your choice of sliced deli meat, bread or wrap, and cheese with lettuce, tomato, and onion. Served with your choice of side.", 10, gatesFourGolfandCountryClubJPsBarandGrillSandwiches.ID)
	gatesFourGolfandCountryClubPhillyCheesesteak := api.CreateItem("Philly Cheesesteak", "Sliced roasted steak with peppers, onions, and mushrooms, topped with provolone cheese in a footlong hoagie roll. Served with your choice of side.", 16, gatesFourGolfandCountryClubJPsBarandGrillSandwiches.ID)
	gatesFourGolfandCountryClubTheItalian := api.CreateItem("The Italian", "Ham, capicola, and pepperoni with lettuce, oil and vinegar, and banana peppers in a footlong hoagie roll. Served with your choice of side.", 15, gatesFourGolfandCountryClubJPsBarandGrillSandwiches.ID)
	gatesFourGolfandCountryClubBlackenedSalmonSandwich := api.CreateItem("Blackened Salmon Sandwich", "A filet of North Atlantic caught salmon seared and topped with lettuce and tomato on a potato bun. Served with your choice of side.", 18, gatesFourGolfandCountryClubJPsBarandGrillSandwiches.ID)
	gatesFourGolfandCountryClubJPsClubSandwich := api.CreateItem("JP’s Club Sandwich", "Black forest ham, roasted turkey, and smoked bacon with Swiss and cheddar cheeses, crisp leaf lettuce, tomato, and mayo on three slices of your choice of bread. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillSandwiches.ID)
	gatesFourGolfandCountryClubBuffaloChickenSandwich := api.CreateItem("Buffalo Chicken Sandwich", "A fried 8oz chicken breast, tossed in buffalo sauce, topped with pepper jack cheese, bacon, leaf lettuce, and tomato on a potato bun. Served with your choice of side.", 17, gatesFourGolfandCountryClubJPsBarandGrillSandwiches.ID)
	gatesFourGolfandCountryClubJPsAce := api.CreateItem("JP’s Ace", "Angus hotdog with chili, coleslaw, mustard, diced onions. Side Optional", 4, gatesFourGolfandCountryClubJPsBarandGrillSandwiches.ID)
	gatesFourGolfandCountryClubGatesFourBurger := api.CreateItem("Gates Four Burger", "A half-pound hamburger grilled with American cheese, lettuce, onion, and tomato. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillBurgers.ID)
	gatesFourGolfandCountryClubBigSkyBurger := api.CreateItem("Big Sky Burger", "Our half-pound burger grilled with barbecue sauce, applewood smoked bacon, cheddar cheese, and crispy fried onions. Served with your choice of side.", 15, gatesFourGolfandCountryClubJPsBarandGrillBurgers.ID)
	gatesFourGolfandCountryClubCarolinaBurger := api.CreateItem("Carolina Burger", "Our half-pound burger grilled with chili, American cheese, coleslaw, mustard, and onions. Served with your choice of sides.", 15, gatesFourGolfandCountryClubJPsBarandGrillBurgers.ID)
	_ = api.CreateItem("French Fries", "Crispy, golden, and perfectly salted.", 2, gatesFourGolfandCountryClubJPsBarandGrillSides.ID)
	_ = api.CreateItem("Sweet Potato Fries", "Sweet, crispy, and addicting.", 2, gatesFourGolfandCountryClubJPsBarandGrillSides.ID)
	_ = api.CreateItem("Fruit", "Seasonal, fresh, and refreshing.", 2, gatesFourGolfandCountryClubJPsBarandGrillSides.ID)
	_ = api.CreateItem("Potato Salad", "Creamy, classic, and loaded with flavor.", 2, gatesFourGolfandCountryClubJPsBarandGrillSides.ID)
	_ = api.CreateItem("Pasta Salad", "Chilled, zesty, and packed with veggies.", 2, gatesFourGolfandCountryClubJPsBarandGrillSides.ID)
	_ = api.CreateItem("Onion Rings", "Thick-cut, crispy, and golden brown.", 3, gatesFourGolfandCountryClubJPsBarandGrillSides.ID)
	// seed item options
	log.Println("Seeding *Options* data")
	_ = api.CreateOption("Wings Sauce", "Wings Sauce", 1, 1, gatesFourGolfandCountryClubChickenWings.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Buffalo", Description: "Buffalo", Price: 0},
		locations.OptionItem{Name: "BBQ", Description: "BBQ", Price: 0},
		locations.OptionItem{Name: "Blazing BBQ", Description: "Blazing BBQ", Price: 0},
		locations.OptionItem{Name: "Incinerator", Description: "Incinerator", Price: 0},
		locations.OptionItem{Name: "Garlic Parm", Description: "Garlic Parm", Price: 0}})
	_ = api.CreateOption("Wings Dipping Sauce", "Wings Dipping Sauce", 1, 1, gatesFourGolfandCountryClubChickenWings.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubChickenWings.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Six Wings", Description: "Six Wings", Price: 0},
		locations.OptionItem{Name: "Twelve Wings", Description: "Twelve Wings", Price: 5}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubClubNachos.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 0},
		locations.OptionItem{Name: "Beef", Description: "Beef", Price: 0},
		locations.OptionItem{Name: "No Protein", Description: "No Protein", Price: 0}})
	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, gatesFourGolfandCountryClubHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Balsalmic", Description: "Balsalmic", Price: 0},
		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "Bleu Cheese", Description: "Bleu Cheese", Price: 0},
		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0},
		locations.OptionItem{Name: "Thousand Island", Description: "Thousand Island", Price: 0},
		locations.OptionItem{Name: "Balsalmic Vinaigrette", Description: "Balsalmic Vinaigrette", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Cucumbers", Description: "No Cucumbers", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Mushrooms", Description: "No Mushrooms", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubCaesarSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubCaesarSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Croutons", Description: "No Croutons", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubClubSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubClubSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Ham", Description: "No Ham", Price: 0},
		locations.OptionItem{Name: "No Turkey", Description: "No Turkey", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubGatesFourSpinachSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubGatesFourSpinachSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Craisins", Description: "No Craisins", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Candied Nuts", Description: "No Candied Nuts", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Blue Cheese", Description: "No Blue Cheese", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubAhiTunaSteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubChickenFajitas.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubChickenFajitas.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Peppers", Description: "No Peppers", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Cojita Cheese", Description: "No Cojita Cheese", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubPortobelloBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubPortobelloBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0}})
	_ = api.CreateOption("Choice of Meat", "Choice of Meat", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ham", Description: "Ham", Price: 0},
		locations.OptionItem{Name: "Roast Beef", Description: "Roast Beef", Price: 0},
		locations.OptionItem{Name: "Capicola", Description: "Capicola", Price: 0},
		locations.OptionItem{Name: "Turkey", Description: "Turkey", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Choice of Bread", "Choice of Bread", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "White", Description: "White", Price: 0},
		locations.OptionItem{Name: "Wheat", Description: "Wheat", Price: 0},
		locations.OptionItem{Name: "Flour Wrap", Description: "Flour Wrap", Price: 0},
		locations.OptionItem{Name: "Garlic Wrap", Description: "Garlic Wrap", Price: 0},
		locations.OptionItem{Name: "Gluten Free Wrap", Description: "Gluten Free Wrap", Price: 1}})
	_ = api.CreateOption("Choice of Cheese", "Choice of Cheese", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "American", Description: "American", Price: 0},
		locations.OptionItem{Name: "Cheddar", Description: "Cheddar", Price: 0},
		locations.OptionItem{Name: "Swiss", Description: "Swiss", Price: 0},
		locations.OptionItem{Name: "Provolone", Description: "Provolone", Price: 0},
		locations.OptionItem{Name: "Pepperjack", Description: "Pepperjack", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubPhillyCheesesteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubPhillyCheesesteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Peppers", Description: "No Peppers", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Mushrooms", Description: "No Mushrooms", Price: 0},
		locations.OptionItem{Name: "No Provolone", Description: "No Provolone", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubTheItalian.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubTheItalian.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Oil & Vinegar", Description: "No Oil & Vinegar", Price: 0},
		locations.OptionItem{Name: "No Banana Peppers", Description: "No Banana Peppers", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubBlackenedSalmonSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubBlackenedSalmonSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubJPsClubSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Choice of Bread", "Choice of Bread", 1, 1, gatesFourGolfandCountryClubJPsClubSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "White", Description: "White", Price: 0},
		locations.OptionItem{Name: "Wheat", Description: "Wheat", Price: 0},
		locations.OptionItem{Name: "Flour Wrap", Description: "Flour Wrap", Price: 0},
		locations.OptionItem{Name: "Garlic Wrap", Description: "Garlic Wrap", Price: 0},
		locations.OptionItem{Name: "Gluten Free Wrap", Description: "Gluten Free Wrap", Price: 1}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubJPsClubSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Mayo", Description: "No Mayo", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Swiss", Description: "No Swiss", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubBuffaloChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubBuffaloChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 0, 0, gatesFourGolfandCountryClubJPsAce.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 2},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 2},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 2},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 2},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 2},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 2},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 2}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubJPsAce.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Chili", Description: "No Chili", Price: 0},
		locations.OptionItem{Name: "No Coleslaw", Description: "No Coleslaw", Price: 0},
		locations.OptionItem{Name: "No Mustard", Description: "No Mustard", Price: 0},
		locations.OptionItem{Name: "No Diced Onions", Description: "No Diced Onions", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubGatesFourBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubGatesFourBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubGatesFourBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubBigSkyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubBigSkyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubBigSkyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Cheddar", Description: "No Cheddar", Price: 0},
		locations.OptionItem{Name: "No Crispy Onions", Description: "No Crispy Onions", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubCarolinaBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubCarolinaBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubCarolinaBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Mustard", Description: "No Mustard", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Coleslaw", Description: "No Coleslaw", Price: 0}})
	// seed orders
	log.Println("Seeding Completed")
}