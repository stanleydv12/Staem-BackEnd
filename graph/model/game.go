package model

import (
	"github.com/stanleydv12/gqlgen-todos/database"
)

type Game struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex"`
	Genre       string
	Price       int64
	Description string
	Tag         string
	Banner      string
	Image       string
	About       string
	Mature      string
	Hours       int
}

type GameImage struct {
	ID     int64 `gorm:"primaryKey"`
	GameId int64
	Image  string
}

type GameVideo struct {
	ID        int64 `gorm:"primaryKey"`
	GameId    int64
	VideoLink string
}

type RecommendedRequirement struct {
	ID        int64 `gorm:"primaryKey"`
	GameId    int64
	OS        string
	Processor string
	Memory    string
	Graphics  string
	Storage   string
}

type MinimumRequirement struct {
	ID        int64 `gorm:"primaryKey"`
	GameId    int64
	OS        string
	Processor string
	Memory    string
	Graphics  string
	Storage   string
}

func init() {

	seedGame()
	seedGameImage()
	seedGameVideo()
	seedRecommendedRequirement()
	seedMinimumRequirement()

}

func seedMinimumRequirement() {
	db := database.GetInstance()
	db.DropTableIfExists(&MinimumRequirement{})
	db.AutoMigrate(&MinimumRequirement{})

	db.Create(&MinimumRequirement{
		GameId:    1,
		OS:        "Windows 7/8.1/10 (64-bit versions)",
		Processor: "Intel Core i5-6600K or AMD Ryzen 5 1400 or better",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 970 / AMD RX 480 8GB or better",
		Storage:   "15 GB available space",
	})

	db.Create(&MinimumRequirement{
		GameId:    2,
		OS:        "Windows 7/8.1/10 (64-bit versions)",
		Processor: "Intel Core i5-6600K or AMD Ryzen 5 1400 or better",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 970 / AMD RX 480 8GB or better",
		Storage:   "15 GB available space",
	})

	db.Create(&MinimumRequirement{
		GameId:    3,
		OS:        "Windows 7/8.1/10 (64-bit versions)",
		Processor: "Intel Core i5-6600K or AMD Ryzen 5 1400 or better",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 970 / AMD RX 480 8GB or better",
		Storage:   "15 GB available space",
	})

	db.Create(&MinimumRequirement{
		GameId:    4,
		OS:        "Windows 7/8.1/10 (64-bit versions)",
		Processor: "Intel Core i5-6600K or AMD Ryzen 5 1400 or better",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 970 / AMD RX 480 8GB or better",
		Storage:   "15 GB available space",
	})

	db.Create(&MinimumRequirement{
		GameId:    5,
		OS:        "Windows 7/8.1/10 (64-bit versions)",
		Processor: "Intel Core i5-6600K or AMD Ryzen 5 1400 or better",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 970 / AMD RX 480 8GB or better",
		Storage:   "15 GB available space",
	})

	db.Create(&MinimumRequirement{
		GameId:    6,
		OS:        "Windows 7/8.1/10 (64-bit versions)",
		Processor: "Intel Core i5-6600K or AMD Ryzen 5 1400 or better",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 970 / AMD RX 480 8GB or better",
		Storage:   "15 GB available space",
	})

}

func seedRecommendedRequirement() {
	db := database.GetInstance()
	db.DropTableIfExists(&RecommendedRequirement{})
	db.AutoMigrate(&RecommendedRequirement{})

	db.Create(&RecommendedRequirement{
		GameId:    1,
		OS:        "Windows 10 (64-bit)",
		Processor: "Intel Core i7-4790 or AMD Ryzen 5 1500X",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1070 8GB / AMD RX Vega 56 8GB",
		Storage:   "15 GB available space",
	})

	db.Create(&RecommendedRequirement{
		GameId:    2,
		OS:        "Windows 10 (64-bit)",
		Processor: "Intel Core i7-4790 or AMD Ryzen 5 1500X",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1070 8GB / AMD RX Vega 56 8GB",
		Storage:   "15 GB available space",
	})

	db.Create(&RecommendedRequirement{
		GameId:    3,
		OS:        "Windows 10 (64-bit)",
		Processor: "Intel Core i7-4790 or AMD Ryzen 5 1500X",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1070 8GB / AMD RX Vega 56 8GB",
		Storage:   "15 GB available space",
	})

	db.Create(&RecommendedRequirement{
		GameId:    4,
		OS:        "Windows 10 (64-bit)",
		Processor: "Intel Core i7-4790 or AMD Ryzen 5 1500X",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1070 8GB / AMD RX Vega 56 8GB",
		Storage:   "15 GB available space",
	})

	db.Create(&RecommendedRequirement{
		GameId:    5,
		OS:        "Windows 10 (64-bit)",
		Processor: "Intel Core i7-4790 or AMD Ryzen 5 1500X",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1070 8GB / AMD RX Vega 56 8GB",
		Storage:   "15 GB available space",
	})

	db.Create(&RecommendedRequirement{
		GameId:    6,
		OS:        "Windows 10 (64-bit)",
		Processor: "Intel Core i7-4790 or AMD Ryzen 5 1500X",
		Memory:    "8 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1070 8GB / AMD RX Vega 56 8GB",
		Storage:   "15 GB available space",
	})
}

func seedGameVideo() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameVideo{})
	db.AutoMigrate(&GameVideo{})

	db.Create(&GameVideo{
		GameId:    1,
		VideoLink: "https://cdn.akamai.steamstatic.com/steam/apps/256740604/movie480.webm?t=1551460377",
	})

	db.Create(&GameVideo{
		GameId:    2,
		VideoLink: "https://cdn.akamai.steamstatic.com/steam/apps/256711431/movie480.webm?t=1521039659",
	})

	db.Create(&GameVideo{
		GameId:    3,
		VideoLink: "https://cdn.akamai.steamstatic.com/steam/apps/256740604/movie480.webm?t=1551460377",
	})

	db.Create(&GameVideo{
		GameId:    4,
		VideoLink: "https://cdn.akamai.steamstatic.com/steam/apps/256711431/movie480.webm?t=1521039659",
	})

	db.Create(&GameVideo{
		GameId:    5,
		VideoLink: "https://cdn.akamai.steamstatic.com/steam/apps/256740604/movie480.webm?t=1551460377",
	})

	db.Create(&GameVideo{
		GameId:    6,
		VideoLink: "https://cdn.akamai.steamstatic.com/steam/apps/256711431/movie480.webm?t=1521039659",
	})

}

func seedGameImage() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameImage{})
	db.AutoMigrate(&GameImage{})

	db.Create(&GameImage{
		GameId: 1,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_f50e0ac29c65ceba9a6e61faa81751229b493a4c.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 1,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_784f12a1c9818045adef23c4b713f3039d803e48.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 1,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_4b63ee11e347ea8d431425405c2af518d905d4e3.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 1,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_85997da66bf40cb0172a47b001a062bb9e7fd1ed.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 2,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_8439fefb3f8efdefae1dfb47aabba1c0ae616ef4.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 2,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_3737964d741715847d62317e6598c86847cb1cfa.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 2,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_5a9d0f4c3a98873de13e64934e6bf46058aabb25.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 2,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_d3fac3f1b058881c7f208cfd8bd53a04e39837c9.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 3,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_f50e0ac29c65ceba9a6e61faa81751229b493a4c.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 3,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_784f12a1c9818045adef23c4b713f3039d803e48.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 3,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_4b63ee11e347ea8d431425405c2af518d905d4e3.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 3,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_85997da66bf40cb0172a47b001a062bb9e7fd1ed.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 4,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_8439fefb3f8efdefae1dfb47aabba1c0ae616ef4.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 4,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_3737964d741715847d62317e6598c86847cb1cfa.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 4,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_5a9d0f4c3a98873de13e64934e6bf46058aabb25.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 4,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_d3fac3f1b058881c7f208cfd8bd53a04e39837c9.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 5,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_f50e0ac29c65ceba9a6e61faa81751229b493a4c.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 5,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_784f12a1c9818045adef23c4b713f3039d803e48.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 5,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_4b63ee11e347ea8d431425405c2af518d905d4e3.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 5,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/823500/ss_85997da66bf40cb0172a47b001a062bb9e7fd1ed.600x338.jpg?t=1581381377",
	})

	db.Create(&GameImage{
		GameId: 6,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_8439fefb3f8efdefae1dfb47aabba1c0ae616ef4.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 6,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_3737964d741715847d62317e6598c86847cb1cfa.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 6,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_5a9d0f4c3a98873de13e64934e6bf46058aabb25.600x338.jpg?t=1564002443",
	})

	db.Create(&GameImage{
		GameId: 6,
		Image:  "https://cdn.akamai.steamstatic.com/steam/apps/611670/ss_d3fac3f1b058881c7f208cfd8bd53a04e39837c9.600x338.jpg?t=1564002443",
	})

}

func seedGame() {
	db := database.GetInstance()
	db.DropTableIfExists(&Game{})
	db.AutoMigrate(&Game{})

	db.Create(&Game{
		Name:        "Boneworks",
		Genre:       "Action",
		Price:       15,
		Description: "BONEWORKS is an Experimental Physics VR Adventure. Use found physics weapons, tools, and objects to fight across dangerous playscapes and mysterious architecture.",
		Tag:         "Featured",
		Banner:      "",
		Image:       "./assets/image/small/boneworks.jpg",
		About:       "BONEWORKS Is a narrative VR action adventure using advanced experimental physics mechanics. Dynamically navigate through environments, engage in physics heavy combat, and creatively approach puzzles with physics. Advanced Physics: Designed entirely for consistent universal rules, the advanced physics mechanics encourage players to confidently and creatively interact with the virtual world however you want. Combat: Approach combat in any number of ways you can think of following the physical rules of the game's universe. Melee weapons, firearms, physics traps, environments, can all be used to aid you in fights with enemy entities. Weapons, lots of weapons: Boneworks provides players with a plethora of physics based weaponry; guns, swords, axes, clubs, spears, hammers, experimental energy weapons, nonsensical mystery tools, and anomalous physics weapons. Interaction: Hyper realistic VR object and environment interaction. Story: Play through the game's mysterious narrative and explore the deep inner workings of the Monogon Industries' artificial intelligence operating system; Myth OS. Character Bodies: Accurate full IK body systems built from the ground up provide a realistic looking body presence and allow for a maximum level of immersion with physical",
		Mature:      "True",
		Hours:       12,
	})

	db.Create(&Game{
		Name:        "Skyrim",
		Genre:       "Casual",
		Price:       20,
		Description: "A true, full-length open-world game for VR has arrived from Bethesda Game Studios. Skyrim VR reimagines the complete epic fantasy masterpiece with an unparalleled sense of scale, depth, and immersion. Skyrim VR also includes all official add-ons.",
		Tag:         "Featured",
		Banner:      "",
		Image:       "./assets/image/small/skyrim.jpg",
		About:       "A true, full-length open-world game for VR has arrived from award-winning developers, Bethesda Game Studios. Skyrim VR reimagines the complete epic fantasy masterpiece with an unparalleled sense of scale, depth, and immersion. From battling ancient dragons to exploring rugged mountains and more, Skyrim VR brings to life a complete open world for you to experience any way you choose. Skyrim VR includes the critically-acclaimed core game and official add-ons – Dawnguard, Hearthfire, and Dragonborn. Dragons, long lost to the passages of the Elder Scrolls, have returned to Tamriel and the future of the Empire hangs in the balance. As Dragonborn, the prophesied hero born with the power of The Voice, you are the only one who can stand amongst them.",
		Mature:      "True",
		Hours:       15,
	})

	db.Create(&Game{
		Name:        "Resident Evil",
		Genre:       "Action",
		Price:       15,
		Description: "BONEWORKS is an Experimental Physics VR Adventure. Use found physics weapons, tools, and objects to fight across dangerous playscapes and mysterious architecture.",
		Tag:         "Featured",
		Banner:      "",
		Image:       "./assets/image/small/boneworks.jpg",
		About:       "BONEWORKS Is a narrative VR action adventure using advanced experimental physics mechanics. Dynamically navigate through environments, engage in physics heavy combat, and creatively approach puzzles with physics. Advanced Physics: Designed entirely for consistent universal rules, the advanced physics mechanics encourage players to confidently and creatively interact with the virtual world however you want. Combat: Approach combat in any number of ways you can think of following the physical rules of the game's universe. Melee weapons, firearms, physics traps, environments, can all be used to aid you in fights with enemy entities. Weapons, lots of weapons: Boneworks provides players with a plethora of physics based weaponry; guns, swords, axes, clubs, spears, hammers, experimental energy weapons, nonsensical mystery tools, and anomalous physics weapons. Interaction: Hyper realistic VR object and environment interaction. Story: Play through the game's mysterious narrative and explore the deep inner workings of the Monogon Industries' artificial intelligence operating system; Myth OS. Character Bodies: Accurate full IK body systems built from the ground up provide a realistic looking body presence and allow for a maximum level of immersion with physical",
		Mature:      "True",
		Hours:       20,
	})

	db.Create(&Game{
		Name:        "World War Z",
		Genre:       "Casual",
		Price:       40,
		Description: "A true, full-length open-world game for VR has arrived from Bethesda Game Studios. Skyrim VR reimagines the complete epic fantasy masterpiece with an unparalleled sense of scale, depth, and immersion. Skyrim VR also includes all official add-ons.",
		Tag:         "Featured",
		Banner:      "",
		Image:       "./assets/image/small/skyrim.jpg",
		About:       "A true, full-length open-world game for VR has arrived from award-winning developers, Bethesda Game Studios. Skyrim VR reimagines the complete epic fantasy masterpiece with an unparalleled sense of scale, depth, and immersion. From battling ancient dragons to exploring rugged mountains and more, Skyrim VR brings to life a complete open world for you to experience any way you choose. Skyrim VR includes the critically-acclaimed core game and official add-ons – Dawnguard, Hearthfire, and Dragonborn. Dragons, long lost to the passages of the Elder Scrolls, have returned to Tamriel and the future of the Empire hangs in the balance. As Dragonborn, the prophesied hero born with the power of The Voice, you are the only one who can stand amongst them.",
		Mature:      "True",
		Hours:       11,
	})

	db.Create(&Game{
		Name:        "Pokemon Go",
		Genre:       "Action",
		Price:       50,
		Description: "BONEWORKS is an Experimental Physics VR Adventure. Use found physics weapons, tools, and objects to fight across dangerous playscapes and mysterious architecture.",
		Tag:         "Featured",
		Banner:      "",
		Image:       "./assets/image/small/boneworks.jpg",
		About:       "BONEWORKS Is a narrative VR action adventure using advanced experimental physics mechanics. Dynamically navigate through environments, engage in physics heavy combat, and creatively approach puzzles with physics. Advanced Physics: Designed entirely for consistent universal rules, the advanced physics mechanics encourage players to confidently and creatively interact with the virtual world however you want. Combat: Approach combat in any number of ways you can think of following the physical rules of the game's universe. Melee weapons, firearms, physics traps, environments, can all be used to aid you in fights with enemy entities. Weapons, lots of weapons: Boneworks provides players with a plethora of physics based weaponry; guns, swords, axes, clubs, spears, hammers, experimental energy weapons, nonsensical mystery tools, and anomalous physics weapons. Interaction: Hyper realistic VR object and environment interaction. Story: Play through the game's mysterious narrative and explore the deep inner workings of the Monogon Industries' artificial intelligence operating system; Myth OS. Character Bodies: Accurate full IK body systems built from the ground up provide a realistic looking body presence and allow for a maximum level of immersion with physical",
		Mature:      "True",
		Hours:       14,
	})

	db.Create(&Game{
		Name:        "Forza Horizon 4",
		Genre:       "Casual",
		Price:       25,
		Description: "A true, full-length open-world game for VR has arrived from Bethesda Game Studios. Skyrim VR reimagines the complete epic fantasy masterpiece with an unparalleled sense of scale, depth, and immersion. Skyrim VR also includes all official add-ons.",
		Tag:         "Featured",
		Banner:      "",
		Image:       "./assets/image/small/skyrim.jpg",
		About:       "A true, full-length open-world game for VR has arrived from award-winning developers, Bethesda Game Studios. Skyrim VR reimagines the complete epic fantasy masterpiece with an unparalleled sense of scale, depth, and immersion. From battling ancient dragons to exploring rugged mountains and more, Skyrim VR brings to life a complete open world for you to experience any way you choose. Skyrim VR includes the critically-acclaimed core game and official add-ons – Dawnguard, Hearthfire, and Dragonborn. Dragons, long lost to the passages of the Elder Scrolls, have returned to Tamriel and the future of the Empire hangs in the balance. As Dragonborn, the prophesied hero born with the power of The Voice, you are the only one who can stand amongst them.",
		Mature:      "True",
		Hours:       130,
	})
}
