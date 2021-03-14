package model

import "github.com/stanleydv12/gqlgen-todos/database"

type ProfileBackground struct {
	ID    int `gorm:"primaryKey"`
	Path  string
	Price int
}

type AvatarBorder struct {
	ID    int `gorm:"primaryKey"`
	Path  string
	Price int
}

type Avatar struct {
	ID    int `gorm:"primaryKey"`
	Path  string
	Price int
}

type ChatSticker struct {
	ID    int `gorm:"primaryKey"`
	Path  string
	Price int
}

type MiniProfileBackground struct {
	ID    int `gorm:"primaryKey"`
	Path  string
	Price int
}

func init() {
	seedAvatarBorder()
	seedAvatar()
	seedAnimatedSticker()
	seedMiniProfileBackground()
	seedProfileBackgrouond()
}

func seedProfileBackgrouond() {
	db := database.GetInstance()
	db.DropTableIfExists(&ProfileBackground{})
	db.AutoMigrate(&ProfileBackground{})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/504400/faf8837de8be0464430ab3d6013de7d0fa30d457.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/601220/5798e88efeb3ad87745d9be18b65ee77a08d3ab0.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/504400/84b2426178f3531e3369289606c570824f97552d.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/601220/79213a222e0eb92e8ec05ec4cf3ab418e3712c2d.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/504400/d387074c0be842017227a4f25fac84871d78a849.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/570/8cda135e1cbd187d10e4a7f798213b3a2acc55d6.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1069740/01dfbab78e1ae9e2cee85958bf364de069b4a1a3.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/531510/27717b93c71f08bf780ffde0e92a1e4ace158039.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1263950/c23ea7d4ee6c8fbae783e1f4ff2955026a3d881f.webm",
		Price: 2000,
	})

	db.Create(&ProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/628750/dd00e31eafcffa1fdc2edc76d3b85e4fd4732a04.webm",
		Price: 2000,
	})
}

func seedMiniProfileBackground() {
	db := database.GetInstance()
	db.DropTableIfExists(&MiniProfileBackground{})
	db.AutoMigrate(&MiniProfileBackground{})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1098340/b0773e1a9f8fbf35fce4e6e5bb3d337ad88adca5.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/601220/cb0599011bfe9f3e28609d6e103605b5d644ba9a.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/601220/7619e7036dff69b0fddedba3ffe28875d1e2d318.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1526200/ece8d079d4b9af7ee2673ab1cddb1bae40511f5f.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1098340/01f19b0271d770fe620f3e458153a1f2e2ac6561.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1098340/599b5eff8c3c712e78caed52114707672dbf0ff6.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/619150/a7e5f47e2808e1e9fdfa0a2ca905dc2daeaaec87.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/628670/a8bea07b80fc87c41670e32df30da37ba0b01ddb.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1069740/8ec8230625bb7742e287dcf276666b3ab2767148.webm",
		Price: 2000,
	})

	db.Create(&MiniProfileBackground{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1098340/780d8aea66d619426e76ea74dc1a90cc60eac04e.webm",
		Price: 2000,
	})
}

func seedAnimatedSticker() {
	db := database.GetInstance()
	db.DropTableIfExists(&ChatSticker{})
	db.AutoMigrate(&ChatSticker{})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1091500/bc003c956445fe939986abf0ccfd903eec0de22b.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/440/3ffc77baebb36beb3e054d30a3bae2e20a9a2d51.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/730/c81eec46e75c35bcd996ae2621d124bdcfa5589d.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/730/ae23a0bc8ec6bd6659b993d053a1eb75fd26eb2f.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/440/b4114df5d49f68c585e84f41e65c229541b87002.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/532190/6cb8a731abe09539f8f2ef2655fb1b42205e5be7.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/440/176c39c5e887375ae3f3e792278c8b5916a2eaa3.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/440/48af430eb84685d08f0a1ac982e744489cbfc278.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1263950/22a76406b250da8978cf333c24f12b8ef6837724.png",
		Price: 2000,
	})

	db.Create(&ChatSticker{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/620980/ee799e8047533329b96ec83e9fb402a5c482b8d8.png",
		Price: 2000,
	})

}

func seedAvatarBorder() {
	db := database.GetInstance()
	db.DropTableIfExists(&AvatarBorder{})
	db.AutoMigrate(&AvatarBorder{})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1492660/06bb85cd5f39a963a39ae9327ea4eb7da5cd30d4.png",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/322330/46461aaea39b18a4a3da2e6d3cf253006f2d6193.png",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1098340/71f42ec23a7f80c365f0c3900a6e61bdc78733d7.png",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/601220/4957cc68eec4c2d9ee1df6de65f073951d1e94ee.png",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/212070/e39b802b3cff406590139dba9de470f31810027c.gif",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/212070/9b6b26c7a03046da283408d72319f9eec932c80a.gif",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/322330/beaee5e90d93bfafa5f5f55acb23abfd28ad180c.png",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/601220/2b817dadbde606a4773e34198e6220d8170da5c3.png",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1069740/99e0f27ac1cd7117b342b035d8b725ddb1b76d28.png",
		Price: 2000,
	})

	db.Create(&AvatarBorder{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/588650/cb7486f3edec85c92a7a544c7714497b5efa15c4.png",
		Price: 2000,
	})
}

func seedAvatar() {
	db := database.GetInstance()
	db.DropTableIfExists(&Avatar{})
	db.AutoMigrate(&Avatar{})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/620/d70a96496a8afc0602d9903859017c60e5dfc319.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/408410/e3698a58ccaf9fbe682a6b22d7ff72472c75a459.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/570/e0910978c2494bea551e70814457ea3606bf2fdc.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/650330/658d7a42b02f59a94b0ea2a97ee46b4323b66c78.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/408410/0538306fa1cafff1035d125ebbe745f1f9ce2236.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/408410/2e94146ee2c43e28ef642ea309f6eef9e939a8fd.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/570/25994d646946e0ec39a5e0e0a76cba85d03044d5.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/1069740/4712839d44f5f73e3a07f6560fb75d9de8b6969f.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/651670/cf4980a1c068d9895e738ec2ce860240e28eecbb.gif",
		Price: 2000,
	})

	db.Create(&Avatar{
		Path:  "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/items/408410/1c85fc9e8ff5d9a132080822739a7347d791bdf4.gif",
		Price: 2000,
	})
}
