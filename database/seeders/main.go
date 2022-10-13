package seeders

import (
	"fmt"
	"golangapi/database"
	"golangapi/models"
	"golangapi/util"
)

func Run()  {
	var pass, _ = util.HashPassword("password")
	var users = []models.User{{Username: "admin", Password: pass}}
	var hotels = []models.Hotel{
		{
			Name: "Hanoi Hotel",
			Floor: 4,
			District: 4,
			Buildin: 1994,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "Lotte Hotel",
			Floor: 4,
			District: 1,
			Buildin: 2015,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "Saigon Hotel",
			Floor: 1,
			District: 12,
			Buildin: 2000,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "Hai Phong Hotel",
			Floor: 10,
			District: 4,
			Buildin: 2016,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "Da Nang Hotel",
			Floor: 2,
			District: 15,
			Buildin: 2020,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "Ha Long Hotel",
			Floor: 5,
			District: 5,
			Buildin: 2004,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "Hue Hotel",
			Floor: 2,
			District: 23,
			Buildin: 2009,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "Phu Quoc Hotel",
			Floor: 4,
			District: 10,
			Buildin: 2012,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "FLC Hotel",
			Floor: 10,
			District: 10,
			Buildin: 2018,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
		{
			Name: "My Hotel",
			Floor: 1,
			District: 1,
			Buildin: 2030,
			Image: "https://t-cf.bstatic.com/xdata/images/hotel/max1024x768/39986330.jpg?k=a637901765021f5cbccb4e33336b81c0cf1b614ddfc86ecb7c7821679af29c73&o=&hp=1",
		},
	}
	db := database.GetConnection()
	db.Where("1 = 1").Delete(&models.User{})
	db.Where("1 = 1").Delete(&models.Hotel{})
	db.Create(&users)
	db.Create(&hotels)
	fmt.Println("Done!!!")
}