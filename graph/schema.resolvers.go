package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/stanleydv12/gqlgen-todos/database"
	"github.com/stanleydv12/gqlgen-todos/graph/generated"
	"github.com/stanleydv12/gqlgen-todos/graph/helper"
	"github.com/stanleydv12/gqlgen-todos/graph/model"
)

func (r *cartResolver) Game(ctx context.Context, obj *model.Cart) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
	return &game, nil
}

func (r *communityContentResolver) User(ctx context.Context, obj *model.CommunityContent) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *communityContentReviewResolver) Content(ctx context.Context, obj *model.CommunityContentReview) (*model.CommunityContent, error) {
	var content model.CommunityContent
	db := database.GetInstance()
	db.Where("id = ?", obj.ContentID).First(&content)
	return &content, nil
}

func (r *communityContentReviewResolver) User(ctx context.Context, obj *model.CommunityContentReview) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *communityDiscussionResolver) Game(ctx context.Context, obj *model.CommunityDiscussion) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
	return &game, nil
}

func (r *communityDiscussionResolver) User(ctx context.Context, obj *model.CommunityDiscussion) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *communityDiscussionDetailResolver) User(ctx context.Context, obj *model.CommunityDiscussionDetail) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *communityReviewResolver) User(ctx context.Context, obj *model.CommunityReview) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *communityReviewDetailResolver) User(ctx context.Context, obj *model.CommunityReviewDetail) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *friendResolver) User(ctx context.Context, obj *model.Friend) (*model.User, error) {
	var temp model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).Find(&temp)
	return &temp, nil
}

func (r *friendResolver) Friend(ctx context.Context, obj *model.Friend) (*model.User, error) {
	var temp model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.FriendID).Find(&temp)
	return &temp, nil
}

func (r *friendRequestResolver) User(ctx context.Context, obj *model.FriendRequest) (*model.User, error) {
	var temp model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).Find(&temp)
	return &temp, nil
}

func (r *friendRequestResolver) Friend(ctx context.Context, obj *model.FriendRequest) (*model.User, error) {
	var temp model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.FriendID).Find(&temp)
	return &temp, nil
}

func (r *gameResolver) Discussions(ctx context.Context, obj *model.Game) ([]*model.CommunityDiscussion, error) {
	var temp []*model.CommunityDiscussion
	db := database.GetInstance()
	db.Where("game_id = ?", obj.ID).Find(&temp)
	return temp, nil
}

func (r *gameResolver) CreatedAt(ctx context.Context, obj *model.Game) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *gameItemResolver) Game(ctx context.Context, obj *model.GameItem) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameId).First(&game)
	return &game, nil
}

func (r *gameItemResolver) Transactions(ctx context.Context, obj *model.GameItem) ([]*model.MarketTransaction, error) {
	var temp []*model.MarketTransaction
	db := database.GetInstance()
	db.Where("game_item_id = ? and created_at BETWEEN (now() - '1 month'::interval) and now()", obj.ID).Find(&temp)
	return temp, nil
}

func (r *gamePromoResolver) Game(ctx context.Context, obj *model.GamePromo) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameId).First(&game)
	return &game, nil
}

func (r *gameReviewResolver) Game(ctx context.Context, obj *model.GameReview) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameId).First(&game)
	return &game, nil
}

func (r *gameReviewResolver) User(ctx context.Context, obj *model.GameReview) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserId).First(&user)
	return &user, nil
}

func (r *marketGameItemResolver) User(ctx context.Context, obj *model.MarketGameItem) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *marketGameItemResolver) GameItem(ctx context.Context, obj *model.MarketGameItem) (*model.GameItem, error) {
	var item model.GameItem
	db := database.GetInstance()
	db.Where("id = ?", obj.GameItemID).First(&item)
	return &item, nil
}

func (r *marketListingResolver) User(ctx context.Context, obj *model.MarketListing) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *marketListingResolver) GameItem(ctx context.Context, obj *model.MarketListing) (*model.GameItem, error) {
	var item model.GameItem
	db := database.GetInstance()
	db.Where("id = ?", obj.GameItemID).First(&item)
	return &item, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.InputUser) (*model.User, error) {
	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Wallet:   0,
		Image:    "./assets/icons/user.png",
		Country:  input.Country,
	}
	db := database.GetInstance()
	db.Create(user)
	return user, nil
}

func (r *mutationResolver) UpdatePositiveReview(ctx context.Context, input string) (*model.GameReview, error) {
	db := database.GetInstance()
	var temp model.GameReview
	if err := db.First(&temp, input).Error; err != nil {
		return nil, err
	}
	temp.Positive += 1
	return &temp, db.Save(&temp).Error
}

func (r *mutationResolver) UpdateNegativeReview(ctx context.Context, input string) (*model.GameReview, error) {
	db := database.GetInstance()
	var temp model.GameReview
	if err := db.First(&temp, input).Error; err != nil {
		return nil, err
	}
	temp.Negative += 1
	return &temp, db.Save(&temp).Error
}

func (r *mutationResolver) InsertReview(ctx context.Context, input model.InputReview) (*model.GameReview, error) {
	db := database.GetInstance()
	gID, _ := strconv.Atoi(input.Gameid)
	uID, _ := strconv.Atoi(input.Userid)
	rev := &model.GameReview{
		GameId:      gID,
		UserId:      uID,
		Description: input.Description,
		Rating:      input.Rating,
		Positive:    0,
		Negative:    0,
	}
	db.Create(rev)
	return rev, nil
}

func (r *mutationResolver) InputWishlist(ctx context.Context, input model.InputWishlist) (*model.Wishlist, error) {
	db := database.GetInstance()
	var gameid int
	var userid int
	gameid, _ = strconv.Atoi(input.Gameid)
	userid, _ = strconv.Atoi(input.Userid)
	temp := &model.Wishlist{
		GameID: gameid,
		UserID: userid,
	}
	db.Create(&temp)
	return temp, nil
}

func (r *mutationResolver) DeleteWishlist(ctx context.Context, input model.InputWishlist) (*model.Wishlist, error) {
	db := database.GetInstance()
	var gameid int
	var userid int
	gameid, _ = strconv.Atoi(input.Gameid)
	userid, _ = strconv.Atoi(input.Userid)
	temp := &model.Wishlist{
		GameID: gameid,
		UserID: userid,
	}
	db.Delete(&temp)
	return temp, nil
}

func (r *mutationResolver) InputCart(ctx context.Context, input model.InputCart) (*model.Cart, error) {
	db := database.GetInstance()
	var gameid int
	var userid int
	gameid, _ = strconv.Atoi(input.Gameid)
	userid, _ = strconv.Atoi(input.Userid)
	temp := &model.Cart{
		GameID: gameid,
		UserID: userid,
	}
	db.Create(&temp)
	return temp, nil
}

func (r *mutationResolver) DeleteCartByGameID(ctx context.Context, input model.InputCart) (*model.Cart, error) {
	db := database.GetInstance()
	var gameid int
	var userid int
	gameid, _ = strconv.Atoi(input.Gameid)
	userid, _ = strconv.Atoi(input.Userid)
	temp := &model.Cart{
		GameID: gameid,
		UserID: userid,
	}
	db.Where("game_id = ?", gameid).Delete(&temp)
	return temp, nil
}

func (r *mutationResolver) InsertGame(ctx context.Context, input model.InputGame) (*model.OwnedGame, error) {
	db := database.GetInstance()
	var gameid int
	var userid int
	gameid, _ = strconv.Atoi(input.Gameid)
	userid, _ = strconv.Atoi(input.Userid)
	game := &model.OwnedGame{
		GameID: gameid,
		UserID: userid,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "sdteherag@gmail.com")
	m.SetHeader("To", "sdteherag@gmail.com")
	m.SetHeader("Subject", "OTP Code From Staem")
	m.SetBody("text", "You successfully bought a game")

	d := gomail.NewDialer("smtp.gmail.com", 587, "sdteherag@gmail.com", "smorznebygoftmdt")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	db.Create(&game)
	return game, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	db := database.GetInstance()
	var id int
	var wallet float64
	var name, email, password, country, image string
	id, _ = strconv.Atoi(input.ID)
	name = input.Name
	email = input.Email
	password = input.Password
	country = input.Country
	wallet = input.Wallet
	image = input.Image

	user := model.User{
		ID:       (id),
		Name:     name,
		Email:    email,
		Password: password,
		Wallet:   wallet,
		Image:    image,
		Country:  country,
	}

	db.Model(&user).Where("id = ?", user.ID).Updates(user)
	return &user, nil
}

func (r *mutationResolver) SavePaymentMethod(ctx context.Context, input model.InputPaymentMethod) (*model.PaymentMethod, error) {
	db := database.GetInstance()
	id, _ := strconv.Atoi(input.Userid)
	payment := model.PaymentMethod{
		UserID:      id,
		Card:        input.Card,
		CardNumber:  input.CardNumber,
		Date:        input.Date,
		Name:        input.Name,
		Address:     input.Address,
		PostalCode:  input.PostalCode,
		PhoneNumber: input.PhoneNumber,
		Country:     input.Country,
	}

	db.Create(&payment)
	return &payment, nil
}

func (r *mutationResolver) AddCommunityContent(ctx context.Context, input model.InputCommunityContent) (*model.CommunityContent, error) {
	userID, _ := strconv.Atoi(input.Userid)
	content := model.CommunityContent{
		ContentPath:        input.ContentPath,
		ContentDescription: input.ContentDescription,
		ContentType:        input.ContentType,
		UserID:             userID,
		Positive:           0,
		Negative:           0,
	}
	db := database.GetInstance()
	db.Create(&content)

	return &content, nil
}

func (r *mutationResolver) AddCommunityContentReview(ctx context.Context, input model.InputCommunityContentReview) (*model.CommunityContentReview, error) {
	contentID, _ := strconv.Atoi(input.ContentID)
	userID, _ := strconv.Atoi(input.UserID)
	review := model.CommunityContentReview{
		ContentID: contentID,
		UserID:    userID,
		Review:    input.Review,
	}

	db := database.GetInstance()
	db.Create(&review)
	return &review, nil
}

func (r *mutationResolver) UpdateCommunityContentThumbs(ctx context.Context, input model.UpdateCommunityContentThumbs) (*model.CommunityContent, error) {
	contentID, _ := strconv.Atoi(input.ContentID)
	content := model.CommunityContent{
		Positive: input.Positive,
		Negative: input.Negative,
	}

	db := database.GetInstance()
	db.Model(&content).Where("id = ?", contentID).Updates(content)
	return &content, nil
}

func (r *mutationResolver) AddCommunityReviewDetail(ctx context.Context, input model.InputCommunityReviewDetail) (*model.CommunityReviewDetail, error) {
	reviewID, _ := strconv.Atoi(input.ReviewID)
	userID, _ := strconv.Atoi(input.UserID)

	review := model.CommunityReviewDetail{
		ReviewID:      reviewID,
		ReviewContent: input.ReviewContent,
		UserID:        userID,
	}

	db := database.GetInstance()
	db.Create(&review)
	return &review, nil
}

func (r *mutationResolver) AddCommunityDiscussionDetail(ctx context.Context, input model.InputCommunityDiscussionDetail) (*model.CommunityDiscussionDetail, error) {
	discussion_id, _ := strconv.Atoi(input.DiscussionID)
	userID, _ := strconv.Atoi(input.UserID)
	gameID, _ := strconv.Atoi(input.GameID)

	detail := model.CommunityDiscussionDetail{
		GameID:            gameID,
		DiscussionID:      discussion_id,
		UserID:            userID,
		DiscussionContent: input.DiscussionContent,
	}

	db := database.GetInstance()
	db.Create(&detail)

	return &detail, nil
}

func (r *mutationResolver) CreateFriendRequest(ctx context.Context, input model.InputFriendRequest) (*model.FriendRequest, error) {
	userID, _ := strconv.Atoi(input.UserID)
	friendID, _ := strconv.Atoi(input.FriendID)

	temp := model.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	}

	db := database.GetInstance()
	db.Create(&temp)

	var user model.User

	db.First(&user, userID)
	db.Model(&user).Update("notification", user.Notification+1)

	return &temp, nil
}

func (r *mutationResolver) CreateFriend(ctx context.Context, input model.InputFriendRequest) (*model.Friend, error) {
	userID, _ := strconv.Atoi(input.UserID)
	friendID, _ := strconv.Atoi(input.FriendID)

	temp := model.Friend{
		UserID:   userID,
		FriendID: friendID,
	}

	temp1 := model.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	}

	db := database.GetInstance()
	db.Create(temp)
	db.Where("user_id = ? and friend_id = ?", temp1.UserID, temp1.FriendID).Delete(temp1)
	db.Where("user_id = ? and friend_id = ?", temp1.FriendID, temp1.UserID).Delete(temp1)

	return &temp, nil
}

func (r *mutationResolver) AddProfileComment(ctx context.Context, input model.InputProfileComment) (*model.ProfileComment, error) {
	userID, _ := strconv.Atoi(input.UserID)

	user := model.ProfileComment{
		UserID:  userID,
		Comment: input.Comment,
	}

	db := database.GetInstance()
	db.Create(&user)

	var temp model.User

	db.First(&user, userID)
	db.Model(&user).Update("notification", temp.Notification+1)

	return &user, nil
}

func (r *mutationResolver) BoughtGameItem(ctx context.Context, input model.InputBoughtMarketGameItem) (*model.OwnedGameItem, error) {
	var user_id, _ = strconv.Atoi(input.UserID)
	var game_item_id, _ = strconv.Atoi(input.GameItemID)
	var buyer_id, _ = strconv.Atoi(input.BuyerID)

	var market = model.MarketGameItem{
		UserID:     user_id,
		GameItemID: game_item_id,
	}

	var item = model.OwnedGameItem{
		GameItemID: game_item_id,
		UserID:     buyer_id,
	}

	var market2 model.MarketGameItem
	db := database.GetInstance()
	db.Where("user_id = ? and game_item_id = ? and type = 'offer'", user_id, game_item_id).First(&market2)

	db.Where("user_id = ? and game_item_id = ? and type = 'offer'", user_id, game_item_id).Delete(&market)
	db.Where("user_id = ? and game_item_id = ?", user_id, game_item_id).Save(&item)

	for _, socket := range r.MarketSocket[game_item_id] {
		socket <- "A user bought an item"
	}

	var user model.User

	db.First(&user, user_id)
	db.Model(&user).Update("notification", user.Notification+1)

	db.Create(&model.MarketTransaction{
		GameItemID: input.GameItemID,
		Price:      market2.Price,
	})

	m := gomail.NewMessage()
	m.SetHeader("From", "sdteherag@gmail.com")
	m.SetHeader("To", "sdteherag@gmail.com")
	m.SetHeader("Subject", "Successfully bought an item on Market")
	m.SetBody("text", "Successfully bought an item on Market")

	d := gomail.NewDialer("smtp.gmail.com", 587, "sdteherag@gmail.com", "smorznebygoftmdt")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return &item, nil
}

func (r *mutationResolver) SoldGameItem(ctx context.Context, input model.InputBoughtMarketGameItem) (*model.OwnedGameItem, error) {
	var user_id, _ = strconv.Atoi(input.UserID)
	var game_item_id, _ = strconv.Atoi(input.GameItemID)
	var seller_id, _ = strconv.Atoi(input.BuyerID)

	var market = model.MarketGameItem{
		UserID:     seller_id,
		GameItemID: game_item_id,
	}

	var item = model.OwnedGameItem{
		GameItemID: game_item_id,
		UserID:     user_id,
	}

	db := database.GetInstance()

	var market2 model.MarketGameItem
	db.First(&market2, game_item_id)

	db.Where("user_id = ? and game_item_id = ? and type = 'bid'", user_id, game_item_id).First(&market2)

	db.Where("user_id = ? and game_item_id = ? and type = 'bid'", user_id, game_item_id).Delete(&market)
	db.Where("user_id = ? and game_item_id = ?", seller_id, game_item_id).Save(&item)

	db.First(&market, game_item_id)
	db.Create(&model.MarketTransaction{
		GameItemID: input.GameItemID,
		Price:      market2.Price,
	})

	for _, socket := range r.MarketSocket[game_item_id] {
		socket <- "A user sold an item"
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "sdteherag@gmail.com")
	m.SetHeader("To", "sdteherag@gmail.com")
	m.SetHeader("Subject", "Successfully sold an item on Market")
	m.SetBody("text", "Successfully sold an item on Market")

	d := gomail.NewDialer("smtp.gmail.com", 587, "sdteherag@gmail.com", "smorznebygoftmdt")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return &item, nil
}

func (r *mutationResolver) AddUserWallet(ctx context.Context, input model.InputUpdateUserWallet) (*model.User, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)

	db := database.GetInstance()

	db.Where("id = ?", user_id).Find(&user)
	user.Wallet = user.Wallet + input.Wallet
	db.Where("id = ?", user_id).Save(&user)

	return &user, nil
}

func (r *mutationResolver) ReduceUserWallet(ctx context.Context, input model.InputUpdateUserWallet) (*model.User, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)

	db := database.GetInstance()

	db.Where("id = ?", user_id).Find(&user)
	user.Wallet = user.Wallet - input.Wallet
	db.Where("id = ?", user_id).Save(&user)

	return &user, nil
}

func (r *mutationResolver) ReducePoint(ctx context.Context, userID string, price int) (*model.User, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(userID)

	db := database.GetInstance()

	db.Where("id = ?", user_id).Find(&user)
	user.Point = user.Point - price
	db.Where("id = ?", user_id).Save(&user)

	return &user, nil
}

func (r *mutationResolver) AddMarketGameItem(ctx context.Context, input model.InputMarketGameItem) (*model.MarketGameItem, error) {
	var user_id, _ = strconv.Atoi(input.UserID)
	var item_id, _ = strconv.Atoi(input.GameItemID)
	var price = strconv.Itoa(input.Price)

	var item = model.MarketGameItem{
		UserID:     user_id,
		GameItemID: item_id,
		Price:      input.Price,
		Type:       input.Type,
	}

	var list = model.MarketListing{
		UserID:     input.UserID,
		GameItemID: input.GameItemID,
		Price:      input.Price,
		Type:       input.Type,
	}

	db := database.GetInstance()
	db.Create(&item)
	db.Create(&list)

	if input.Type == "offer" {
		for _, socket := range r.MarketSocket[item_id] {
			socket <- "A user added offer for " + price
		}
	} else {
		for _, socket := range r.MarketSocket[item_id] {
			socket <- "A user added bid for " + price
		}
	}

	return &item, nil
}

func (r *mutationResolver) IgnoreFriend(ctx context.Context, input model.InputFriendRequest) (*model.FriendRequest, error) {
	var user_id, _ = strconv.Atoi(input.UserID)
	var friend_id, _ = strconv.Atoi(input.FriendID)

	var temp model.FriendRequest

	db := database.GetInstance()
	db.Where("user_id = ? and friend_id = ?", user_id, friend_id).Delete(&temp)

	return &temp, nil
}

func (r *mutationResolver) DeclinedFriend(ctx context.Context, input model.InputFriendRequest) (*model.FriendRequest, error) {
	var user_id, _ = strconv.Atoi(input.UserID)
	var friend_id, _ = strconv.Atoi(input.FriendID)

	var temp model.FriendRequest

	db := database.GetInstance()
	db.Where("user_id = ? and friend_id = ?", user_id, friend_id).Delete(&temp)
	db.Where("user_id = ? and friend_id = ?", friend_id, user_id).Delete(&temp)

	return &temp, nil
}

func (r *mutationResolver) SendOtp(ctx context.Context, input int) (int, error) {
	if err := helper.UseCache().Set(ctx, string(input), input, 10*time.Second).Err(); err != nil {
		log.Fatal(err)
	}

	cached, _ := helper.UseCache().Get(ctx, string(input)).Result()

	m := gomail.NewMessage()
	m.SetHeader("From", "sdteherag@gmail.com")
	m.SetHeader("To", "sdteherag@gmail.com")
	m.SetHeader("Subject", "OTP Code From Staem")
	m.SetBody("text", strconv.Itoa(input))

	d := gomail.NewDialer("smtp.gmail.com", 587, "sdteherag@gmail.com", "smorznebygoftmdt")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	cache, _ := strconv.Atoi(cached)

	return cache, nil
}

func (r *mutationResolver) AddMarketList(ctx context.Context, input model.InputMarketListing) (*model.MarketListing, error) {
	list := &model.MarketListing{
		UserID:     input.UserID,
		GameItemID: input.GameItemID,
		Price:      input.Price,
		Type:       input.Type,
	}
	db := database.GetInstance()
	db.Create(list)

	return list, nil
}

func (r *mutationResolver) BuyItemAtPointShop(ctx context.Context, userID string, itemID string, typeArg string) (string, error) {
	db := database.GetInstance()

	if typeArg == "avatar" {
		db.Create(&model.OwnedAvatar{
			UserID: userID,
			ItemID: itemID,
		})
	} else if typeArg == "border" {
		db.Create(&model.OwnedAvatarBorder{
			UserID: userID,
			ItemID: itemID,
		})
	} else if typeArg == "sticker" {
		db.Create(&model.OwnedChatSticker{
			UserID: userID,
			ItemID: itemID,
		})
	} else if typeArg == "mini-background" {
		db.Create(&model.OwnedMiniProfileBackground{
			UserID: userID,
			ItemID: itemID,
		})
	} else if typeArg == "profile-background" {
		db.Create(&model.OwnedProfileBackground{
			UserID: userID,
			ItemID: itemID,
		})
	}
	return "Success", nil
}

func (r *mutationResolver) AddNewGame(ctx context.Context, input model.InputNewGame) (*model.Game, error) {
	game := model.Game{
		Name:        input.Name,
		Genre:       input.Genre,
		Price:       (int64(input.Price)),
		Description: input.Description,
		Tag:         input.Tag,
		Image:       input.Image,
		About:       input.About,
		Mature:      input.Mature,
		Hours:       0,
	}

	db := database.GetInstance()
	db.Create(&game)

	return &game, nil
}

func (r *mutationResolver) UpdateGame(ctx context.Context, input model.InputUpdateGame) (*model.Game, error) {
	id, _ := strconv.Atoi(input.ID)
	game := model.Game{
		ID:          int64(id),
		Name:        input.Name,
		Genre:       input.Genre,
		Price:       (int64(input.Price)),
		Description: input.Description,
		Tag:         input.Tag,
		Image:       input.Image,
		About:       input.About,
		Mature:      input.Mature,
	}

	db := database.GetInstance()
	db.Model(&model.Game{}).Where("id = ?", game.ID).Updates(&game)
	return &game, nil
}

func (r *mutationResolver) AddNewPromo(ctx context.Context, input model.InputPromo) (*model.GamePromo, error) {
	game_id, _ := strconv.Atoi(input.ID)
	var wish model.Wishlist
	promo := model.GamePromo{
		GameId: game_id,
		Amount: input.Amount,
	}

	db := database.GetInstance()
	result := db.Where("game_id = ? and user_id = 1", input.ID).First(&wish)
	if result.RowsAffected > 0 {
		m := gomail.NewMessage()
		m.SetHeader("From", "sdteherag@gmail.com")
		m.SetHeader("To", "sdteherag@gmail.com")
		m.SetHeader("Subject", "Discount")
		m.SetBody("text", "Your game in wishlist is on discount")

		d := gomail.NewDialer("smtp.gmail.com", 587, "sdteherag@gmail.com", "smorznebygoftmdt")

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}

	db.Create(&promo)

	return &promo, nil
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, input model.InputPromo) (*model.GamePromo, error) {
	game_id, _ := strconv.Atoi(input.ID)
	promo := model.GamePromo{
		GameId: game_id,
		Amount: input.Amount,
	}

	db := database.GetInstance()
	db.Model(&model.GamePromo{}).Where("game_id = ?", game_id).Updates(&promo)

	return &promo, nil
}

func (r *mutationResolver) CreateUnsuspensionRequest(ctx context.Context, input model.InputUnsuspensionRequest) (*model.UnsuspensionRequest, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("email = ?", input.UserEmail).First(&user)

	result := model.UnsuspensionRequest{
		UserID: (user.ID),
		Reason: input.Reason,
		Status: "pending",
	}

	db.Create(&result)

	return &result, nil
}

func (r *mutationResolver) CreateReportRequest(ctx context.Context, input model.InputRequestReport) (*model.ReportRequest, error) {
	reporter_id, _ := strconv.Atoi(input.ReporterID)
	suspected_id, _ := strconv.Atoi(input.SuspectedID)
	var req = &model.ReportRequest{
		ReporterID:  reporter_id,
		SuspectedID: suspected_id,
		Reason:      input.Reason,
	}

	db := database.GetInstance()
	db.Create(&req)
	return req, nil
}

func (r *mutationResolver) AddReported(ctx context.Context, input string) (*model.User, error) {
	var user model.User

	db := database.GetInstance()
	db.Where("id = ?", input).First(&user)
	user.Reported++
	db.Save(&user)

	return &user, nil
}

func (r *mutationResolver) CreateSuspensionList(ctx context.Context, input model.InputSuspensionList) (string, error) {
	user_id, _ := strconv.Atoi(input.UserID)
	var user model.User
	var ulist model.UnsuspensionRequest
	db := database.GetInstance()

	db.Where("id = ?", user_id).First(&user)

	var list model.SuspensionList
	if input.Suspended {
		list = model.SuspensionList{
			UserID:    user_id,
			Reason:    input.Reason,
			Suspended: true,
		}
		user.Suspended = true
	} else {
		list = model.SuspensionList{
			UserID:    user_id,
			Reason:    input.Reason,
			Suspended: false,
		}
		user.Reported = 0
		user.Suspended = false
	}

	db.Where("user_id = ?", user_id).Delete(&ulist)
	db.Save(&user)
	db.Create(&list)
	return "Success", nil
}

func (r *mutationResolver) UpdateUserProfileDetail(ctx context.Context, input model.UpdateProfileDetail) (string, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)
	db := database.GetInstance()
	db.Where("id = ?", user_id).First(&user)
	user.Name = input.Name
	user.Summary = input.Summary

	db.Save(&user)
	return "Success", nil
}

func (r *mutationResolver) UpdateAvatar(ctx context.Context, input model.UpdateProfileItem) (string, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)
	db := database.GetInstance()
	db.Where("id = ?", user_id).First(&user)
	user.Image = input.Path

	db.Save(&user)
	return "Success", nil
}

func (r *mutationResolver) UpdateAvatarBorder(ctx context.Context, input model.UpdateProfileItem) (string, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)
	db := database.GetInstance()
	db.Where("id = ?", user_id).First(&user)
	user.Border = input.Path

	db.Save(&user)
	return "Success", nil
}

func (r *mutationResolver) UpdateMiniProfileBackground(ctx context.Context, input model.UpdateProfileItem) (string, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)
	db := database.GetInstance()
	db.Where("id = ?", user_id).First(&user)
	user.MiniProfileBackground = input.Path

	db.Save(&user)
	return "Success", nil
}

func (r *mutationResolver) UpdateProfileBackground(ctx context.Context, input model.UpdateProfileItem) (string, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)
	db := database.GetInstance()
	db.Where("id = ?", user_id).First(&user)
	user.Background = input.Path

	db.Save(&user)
	return "Success", nil
}

func (r *mutationResolver) InsertChat(ctx context.Context, msg string, friendID string) (string, error) {
	friend, _ := strconv.Atoi(friendID)
	r.ChatSocket[friend] <- msg
	return msg, nil
}

func (r *mutationResolver) UpdateTheme(ctx context.Context, input model.UpdateProfileItem) (string, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)
	db := database.GetInstance()
	db.Where("id = ?", user_id).First(&user)
	user.Theme = input.Path

	db.Save(&user)
	return "Success", nil
}

func (r *mutationResolver) UpdateBadges(ctx context.Context, input model.UpdateProfileItem) (string, error) {
	var user model.User
	var user_id, _ = strconv.Atoi(input.UserID)
	db := database.GetInstance()
	db.Where("id = ?", user_id).First(&user)
	user.Badge = input.Path

	db.Save(&user)
	return "Success", nil
}

func (r *ownedAvatarResolver) Item(ctx context.Context, obj *model.OwnedAvatar) (*model.Avatar, error) {
	var item model.Avatar
	db := database.GetInstance()
	db.Where("id = ?", obj.ItemID).Find(&item)
	return &item, nil
}

func (r *ownedAvatarBorderResolver) Item(ctx context.Context, obj *model.OwnedAvatarBorder) (*model.AvatarBorder, error) {
	var item model.AvatarBorder
	db := database.GetInstance()
	db.Where("id = ?", obj.ItemID).Find(&item)
	return &item, nil
}

func (r *ownedBadgeResolver) Game(ctx context.Context, obj *model.OwnedBadge) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
	return &game, nil
}

func (r *ownedBadgeResolver) Badge(ctx context.Context, obj *model.OwnedBadge) (*model.Badges, error) {
	var badge model.Badges
	db := database.GetInstance()
	db.Where("id = ?", obj.GameBadgeID).Find(&badge)
	return &badge, nil
}

func (r *ownedChatStickerResolver) Item(ctx context.Context, obj *model.OwnedChatSticker) (*model.ChatSticker, error) {
	var item model.ChatSticker
	db := database.GetInstance()
	db.Where("id = ?", obj.ItemID).Find(&item)
	return &item, nil
}

func (r *ownedGameResolver) Game(ctx context.Context, obj *model.OwnedGame) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
	return &game, nil
}

func (r *ownedGameResolver) User(ctx context.Context, obj *model.OwnedGame) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *ownedGameResolver) GameItems(ctx context.Context, obj *model.OwnedGame, currentPage int) ([]*model.OwnedGameItem, error) {
	var items []*model.OwnedGameItem
	db := database.GetInstance()
	db.Scopes(helper.Paginate(currentPage)).Joins("join game_items gi on owned_game_items.game_item_id = gi.id").Where("game_id = ?", obj.GameID).Find(&items)
	return items, nil
}

func (r *ownedGameItemResolver) GameItem(ctx context.Context, obj *model.OwnedGameItem) (*model.GameItem, error) {
	var item model.GameItem
	db := database.GetInstance()
	db.Where("id = ?", obj.GameItemID).Find(&item)
	return &item, nil
}

func (r *ownedMiniProfileBackgroundResolver) Item(ctx context.Context, obj *model.OwnedMiniProfileBackground) (*model.MiniProfileBackground, error) {
	var item model.MiniProfileBackground
	db := database.GetInstance()
	db.Where("id = ?", obj.ItemID).Find(&item)
	return &item, nil
}

func (r *ownedProfileBackgroundResolver) Item(ctx context.Context, obj *model.OwnedProfileBackground) (*model.ProfileBackground, error) {
	var item model.ProfileBackground
	db := database.GetInstance()
	db.Where("id = ?", obj.ItemID).Find(&item)
	return &item, nil
}

func (r *ownedTradingCardResolver) Game(ctx context.Context, obj *model.OwnedTradingCard) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
	return &game, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db := database.GetInstance()
	var u []*model.User
	db.Find(&u)
	return u, nil
}

func (r *queryResolver) Countries(ctx context.Context) ([]*model.Country, error) {
	db := database.GetInstance()
	var u []*model.Country
	db.Find(&u)
	return u, nil
}

func (r *queryResolver) Login(ctx context.Context, input model.LoginUser) (string, error) {
	var user model.User
	db := database.GetInstance()

	result := db.Where("email = ? and password = ?", input.Email, input.Password).First(&user)
	if user.Reported >= 5 {
		user.Suspended = true
		return "Suspended", nil
	}

	if result.RowsAffected == 0 {
		return "", nil
	}
	return helper.GenerateToken(int(user.ID))
}

func (r *queryResolver) LoginAdmin(ctx context.Context, input model.LoginUser) (string, error) {
	var user model.User

	if input.Email == "admin@admin.com" && input.Password == "admin" {
		db := database.GetInstance()
		result := db.Where("email = ? and password = ?", input.Email, input.Password).First(&user)

		if result.RowsAffected == 0 {
			return "", nil
		}

		return helper.GenerateToken(int(user.ID))
	}
	return "", errors.New("NOT ADMIN")
}

func (r *queryResolver) Auth(ctx context.Context, input string) (int, error) {
	var temp, err = helper.ParseToken(input)
	return int(temp), err
}

func (r *queryResolver) GetUserByID(ctx context.Context, input string) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", input).First(&user)

	return &user, nil
}

func (r *queryResolver) GetImageSlider(ctx context.Context) ([]*model.Game_Slider, error) {
	var games []*model.Game_Slider
	db := database.GetInstance()
	db.Find(&games)
	return games, nil
}

func (r *queryResolver) GetGameByTag(ctx context.Context, input string) ([]*model.Game, error) {
	var games []*model.Game
	db := database.GetInstance()
	if input == "Featured" {
		db.Order("hours desc").Limit(4).Find(&games)
	}
	return games, nil
}

func (r *queryResolver) GetGameByID(ctx context.Context, input string) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", input).First(&game)
	return &game, nil
}

func (r *queryResolver) GetGameImageByID(ctx context.Context, input string) ([]*model.GameImage, error) {
	var images []*model.GameImage
	db := database.GetInstance()
	db.Where("game_id = ?", input).Find(&images)
	return images, nil
}

func (r *queryResolver) GetGameVideoByID(ctx context.Context, input string) (*model.GameVideo, error) {
	var video model.GameVideo
	db := database.GetInstance()
	db.Where("game_id = ?", input).First(&video)
	return &video, nil
}

func (r *queryResolver) GetRecommendedRequirement(ctx context.Context, input string) (*model.RecommendedRequirement, error) {
	var req model.RecommendedRequirement
	db := database.GetInstance()
	db.Where("game_id = ?", input).First(&req)
	return &req, nil
}

func (r *queryResolver) GetMinimumRequirement(ctx context.Context, input string) (*model.MinimumRequirement, error) {
	var req model.MinimumRequirement
	db := database.GetInstance()
	db.Where("game_id = ?", input).First(&req)
	return &req, nil
}

func (r *queryResolver) GetMostHelpfulReview(ctx context.Context, input string) ([]*model.GameReview, error) {
	var rev []*model.GameReview
	db := database.GetInstance()
	db.Where("game_id = ? and positive > 1 ", input).Limit(5).Find(&rev)
	return rev, nil
}

func (r *queryResolver) GetRecentReview(ctx context.Context, input string) ([]*model.GameReview, error) {
	var rev []*model.GameReview
	db := database.GetInstance()
	db.Where("game_id = ?", input).Limit(5).Find(&rev)
	return rev, nil
}

func (r *queryResolver) CheckOwnedGame(ctx context.Context, input model.InputOwnedGame) (bool, error) {
	db := database.GetInstance()
	var temp model.OwnedGame
	result := db.Where("game_id = ? and user_id = ?", input.Gameid, input.Userid).First(&temp)

	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (r *queryResolver) CheckWishlist(ctx context.Context, input model.InputWishlist) (bool, error) {
	db := database.GetInstance()
	var temp model.Wishlist
	result := db.Where("game_id = ? and user_id = ?", input.Gameid, input.Userid).First(&temp)

	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (r *queryResolver) GetCartByID(ctx context.Context, input string) ([]*model.Cart, error) {
	db := database.GetInstance()
	var temp []*model.Cart
	db.Where("user_id = ?", input).Find(&temp)

	return temp, nil
}

func (r *queryResolver) GetWishlists(ctx context.Context, input string) ([]*model.Wishlist, error) {
	db := database.GetInstance()
	var temp []*model.Wishlist
	db.Where("user_id = ?", input).Find(&temp)

	return temp, nil
}

func (r *queryResolver) GetCommunityContent(ctx context.Context) ([]*model.CommunityContent, error) {
	db := database.GetInstance()
	var temp []*model.CommunityContent
	db.Find(&temp)

	return temp, nil
}

func (r *queryResolver) GetCommunityReview(ctx context.Context) ([]*model.CommunityReview, error) {
	db := database.GetInstance()
	var temp []*model.CommunityReview
	db.Find(&temp)

	return temp, nil
}

func (r *queryResolver) GetCommunityContentDetail(ctx context.Context, contentID string, paginator int) ([]*model.CommunityContentReview, error) {
	var content []*model.CommunityContentReview
	db := database.GetInstance()
	db.Scopes(helper.Paginate(paginator)).Where("content_id = ?", contentID).Find(&content)

	return content, nil
}

func (r *queryResolver) GetCommunityReviewDetail(ctx context.Context, reviewID string, paginator int) ([]*model.CommunityReviewDetail, error) {
	var review []*model.CommunityReviewDetail
	db := database.GetInstance()
	db.Scopes(helper.Paginate(paginator)).Where("review_id = ?", reviewID).Find(&review)

	return review, nil
}

func (r *queryResolver) GetCommunityDiscussionGame(ctx context.Context) ([]*model.Game, error) {
	var game []*model.Game
	db := database.GetInstance()
	db.Find(&game)
	return game, nil
}

func (r *queryResolver) GetCommunityDiscussioDetail(ctx context.Context, input string) ([]*model.CommunityDiscussionDetail, error) {
	var detail []*model.CommunityDiscussionDetail
	db := database.GetInstance()
	db.Where("discussion_id = ?", input).Find(&detail)
	return detail, nil
}

func (r *queryResolver) GetProfileComment(ctx context.Context, id string, paginator int) ([]*model.ProfileComment, error) {
	var comment []*model.ProfileComment
	db := database.GetInstance()
	db.Scopes(helper.Paginate(paginator)).Where("user_id = ?", id).Find(&comment)
	return comment, nil
}

func (r *queryResolver) GetCommunityDiscussionByID(ctx context.Context, input model.ProfilePaginate) ([]*model.CommunityDiscussion, error) {
	var disc []*model.CommunityDiscussion
	db := database.GetInstance()
	userId, _ := strconv.Atoi(input.UserID)
	db.Where("user_id = ?", userId).Offset(input.Offset).Limit(2).Find(&disc)

	return disc, nil
}

func (r *queryResolver) GetCommunityReviewByID(ctx context.Context, input model.ProfilePaginate) ([]*model.CommunityReview, error) {
	var disc []*model.CommunityReview
	db := database.GetInstance()
	userId, _ := strconv.Atoi(input.UserID)
	db.Where("user_id = ?", userId).Offset(input.Offset).Limit(2).Find(&disc)
	return disc, nil
}

func (r *queryResolver) GetOwnedGames(ctx context.Context, input string) ([]*model.OwnedGame, error) {
	var game []*model.OwnedGame

	db := database.GetInstance()
	db.Where("user_id = ?", input).Find(&game)
	return game, nil
}

func (r *queryResolver) GetGameItem(ctx context.Context, input int) ([]*model.GameItem, error) {
	var gameItem []*model.GameItem
	db := database.GetInstance()
	db.Scopes(helper.Paginate(input)).Find(&gameItem)
	return gameItem, nil
}

func (r *queryResolver) GetGameItemByID(ctx context.Context, input string) (*model.GameItem, error) {
	var item model.GameItem
	db := database.GetInstance()
	db.Where("id = ?", input).First(&item)
	return &item, nil
}

func (r *queryResolver) GetMarketGameItemByID(ctx context.Context, input string) ([]*model.MarketGameItem, error) {
	var item []*model.MarketGameItem
	db := database.GetInstance()
	db.Where("game_item_id = ?", input).Find(&item)
	return item, nil
}

func (r *queryResolver) CheckOwnedGameItem(ctx context.Context, input model.InputOwnedGameItem) (bool, error) {
	db := database.GetInstance()
	var temp model.OwnedGameItem
	result := db.Where("game_item_id = ? and user_id = ?", input.GameItemID, input.UserID).First(&temp)

	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (r *queryResolver) CheckFriend(ctx context.Context, input model.InputFriendRequest) (bool, error) {
	db := database.GetInstance()
	var temp model.Friend
	var temp1 model.Friend
	result := db.Where("friend_id = ? and user_id = ?", input.UserID, input.FriendID).First(&temp)
	result1 := db.Where("friend_id = ? and user_id = ?", input.FriendID, input.UserID).First(&temp1)

	if result.RowsAffected == 0 && result1.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (r *queryResolver) CheckWallet(ctx context.Context, input model.InputTopUpWallet) (bool, error) {
	var wallet model.WalletCode

	db := database.GetInstance()
	result := db.Where("code = ?", input.Code).First(&wallet)

	if result.RowsAffected == 0 {
		return false, nil
	} else {
		var user model.User

		db.First(&user, input.UserID)
		user.Wallet = user.Wallet + wallet.Amount
		db.Save(&user)
		return true, nil
	}
}

func (r *queryResolver) GetMostRecentGames(ctx context.Context) ([]*model.Game, error) {
	var games []*model.Game

	db := database.GetInstance()
	db.Find(&games).Order("id desc").Limit(10).Find(&games)
	return games, nil
}

func (r *queryResolver) GetPersonalizedGames(ctx context.Context) ([]*model.Game, error) {
	var games []*model.Game

	db := database.GetInstance()
	db.Find(&games).Order("random()").Limit(10).Find(&games)
	return games, nil
}

func (r *queryResolver) GetMarketListing(ctx context.Context, input *model.InputGetMarketListing) ([]*model.MarketListing, error) {
	var list []*model.MarketListing

	db := database.GetInstance()
	db.Where("user_id = ? and game_item_id = ?", input.UserID, input.GameItemID).Find(&list)

	return list, nil
}

func (r *queryResolver) GetGameItemByUserID(ctx context.Context, input string) ([]*model.OwnedGameItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAvatar(ctx context.Context) ([]*model.Avatar, error) {
	var item []*model.Avatar

	db := database.GetInstance()
	db.Find(&item)

	return item, nil
}

func (r *queryResolver) GetAvatarBorder(ctx context.Context) ([]*model.AvatarBorder, error) {
	var item []*model.AvatarBorder

	db := database.GetInstance()
	db.Find(&item)

	return item, nil
}

func (r *queryResolver) GetProfileBackground(ctx context.Context) ([]*model.ProfileBackground, error) {
	var item []*model.ProfileBackground

	db := database.GetInstance()
	db.Find(&item)

	return item, nil
}

func (r *queryResolver) GetMiniProfileBackground(ctx context.Context) ([]*model.MiniProfileBackground, error) {
	var item []*model.MiniProfileBackground

	db := database.GetInstance()
	db.Find(&item)

	return item, nil
}

func (r *queryResolver) GetChatSticker(ctx context.Context) ([]*model.ChatSticker, error) {
	var item []*model.ChatSticker
	db := database.GetInstance()
	db.Find(&item)
	return item, nil
}

func (r *queryResolver) GetAllGames(ctx context.Context, paginator int) ([]*model.Game, error) {
	var game []*model.Game
	db := database.GetInstance()
	db.Scopes(helper.Paginate(paginator)).Find(&game)
	return game, nil
}

func (r *queryResolver) GetGames(ctx context.Context) ([]*model.Game, error) {
	var game []*model.Game
	db := database.GetInstance()
	db.Find(&game)
	return game, nil
}

func (r *queryResolver) GetAllPromo(ctx context.Context, paginator int) ([]*model.GamePromo, error) {
	var promo []*model.GamePromo
	db := database.GetInstance()
	db.Scopes(helper.Paginate(paginator)).Find(&promo)
	return promo, nil
}

func (r *queryResolver) DeleteGame(ctx context.Context, gameID string) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", gameID).Delete(&game)

	return &game, nil
}

func (r *queryResolver) DeletePromo(ctx context.Context, promoID string) (*model.GamePromo, error) {
	var game model.GamePromo
	db := database.GetInstance()
	db.Where("game_id = ?", promoID).Delete(&game)

	return &game, nil
}

func (r *queryResolver) GetAllUsersPaginate(ctx context.Context, paginator int) ([]*model.User, error) {
	var user []*model.User
	db := database.GetInstance()
	db.Scopes(helper.Paginate(paginator)).Find(&user)

	return user, nil
}

func (r *queryResolver) GetReportRequest(ctx context.Context) ([]*model.ReportRequest, error) {
	var req []*model.ReportRequest

	db := database.GetInstance()
	db.Find(&req)
	return req, nil
}

func (r *queryResolver) GetUnsuspensionRequest(ctx context.Context) ([]*model.UnsuspensionRequest, error) {
	var req []*model.UnsuspensionRequest

	db := database.GetInstance()
	db.Find(&req)
	return req, nil
}

func (r *queryResolver) GetSuspensionList(ctx context.Context) ([]*model.SuspensionList, error) {
	var list []*model.SuspensionList

	db := database.GetInstance()
	db.Find(&list)

	return list, nil
}

func (r *queryResolver) DeleteReport(ctx context.Context, id string) (string, error) {
	var report model.ReportRequest

	db := database.GetInstance()
	db.Where("id = ?", id).Delete(&report)
	return "Success", nil
}

func (r *queryResolver) GetOwnedAvatar(ctx context.Context, id string) ([]*model.OwnedAvatar, error) {
	var item []*model.OwnedAvatar

	db := database.GetInstance()
	db.Where("user_id = ?", id).Find(&item)

	return item, nil
}

func (r *queryResolver) GetOwnedAvatarBorder(ctx context.Context, id string) ([]*model.OwnedAvatarBorder, error) {
	var item []*model.OwnedAvatarBorder

	db := database.GetInstance()
	db.Where("user_id = ?", id).Find(&item)

	return item, nil
}

func (r *queryResolver) GetOwnedProfileBackground(ctx context.Context, id string) ([]*model.OwnedProfileBackground, error) {
	var item []*model.OwnedProfileBackground

	db := database.GetInstance()
	db.Where("user_id = ?", id).Find(&item)

	return item, nil
}

func (r *queryResolver) GetOwnedMiniProfile(ctx context.Context, id string) ([]*model.OwnedMiniProfileBackground, error) {
	var item []*model.OwnedMiniProfileBackground

	db := database.GetInstance()
	db.Where("user_id = ?", id).Find(&item)

	return item, nil
}

func (r *queryResolver) GetOwnedChatSticker(ctx context.Context, id string) ([]*model.OwnedChatSticker, error) {
	var item []*model.OwnedChatSticker

	db := database.GetInstance()
	db.Where("user_id = ?", id).Find(&item)

	return item, nil
}

func (r *queryResolver) GetOwnedBadges(ctx context.Context, id string) ([]*model.OwnedBadge, error) {
	var badge []*model.OwnedBadge

	db := database.GetInstance()
	db.Where("user_id = ?", id).Find(&badge)

	return badge, nil
}

func (r *queryResolver) GetBadges(ctx context.Context) ([]*model.Badges, error) {
	var badges []*model.Badges
	db := database.GetInstance()
	db.Find(&badges)
	return badges, nil
}

func (r *queryResolver) GetTradingCards(ctx context.Context) ([]*model.TradingCard, error) {
	var badges []*model.TradingCard
	db := database.GetInstance()
	db.Find(&badges)
	return badges, nil
}

func (r *reportRequestResolver) Reporter(ctx context.Context, obj *model.ReportRequest) (*model.User, error) {
	var user model.User

	db := database.GetInstance()
	db.Where("id = ?", obj.ReporterID).First(&user)
	return &user, nil
}

func (r *reportRequestResolver) Suspected(ctx context.Context, obj *model.ReportRequest) (*model.User, error) {
	var user model.User

	db := database.GetInstance()
	db.Where("id = ?", obj.SuspectedID).First(&user)
	return &user, nil
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, itemID int) (<-chan string, error) {
	event := make(chan string, 1)
	r.MarketSocket[itemID] = append(r.MarketSocket[itemID], event)
	return event, nil
}

func (r *subscriptionResolver) PrivateChatAdded(ctx context.Context, userID string) (<-chan string, error) {
	event := make(chan string, 1)
	var user, _ = strconv.Atoi(userID)
	r.ChatSocket[user] = event
	return event, nil
}

func (r *suspensionListResolver) User(ctx context.Context, obj *model.SuspensionList) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *unsuspensionRequestResolver) User(ctx context.Context, obj *model.UnsuspensionRequest) (*model.User, error) {
	var user model.User
	db := database.GetInstance()
	db.Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.Friend, error) {
	var request []*model.Friend
	db := database.GetInstance()
	db.Where("user_id = ?", obj.ID).Find(&request)
	return request, nil
}

func (r *userResolver) FriendRequest(ctx context.Context, obj *model.User) ([]*model.FriendRequest, error) {
	var request []*model.FriendRequest
	db := database.GetInstance()
	db.Where("user_id = ?", obj.ID).Find(&request)
	return request, nil
}

func (r *wishlistResolver) Game(ctx context.Context, obj *model.Wishlist) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
	return &game, nil
}

// Cart returns generated.CartResolver implementation.
func (r *Resolver) Cart() generated.CartResolver { return &cartResolver{r} }

// CommunityContent returns generated.CommunityContentResolver implementation.
func (r *Resolver) CommunityContent() generated.CommunityContentResolver {
	return &communityContentResolver{r}
}

// CommunityContentReview returns generated.CommunityContentReviewResolver implementation.
func (r *Resolver) CommunityContentReview() generated.CommunityContentReviewResolver {
	return &communityContentReviewResolver{r}
}

// CommunityDiscussion returns generated.CommunityDiscussionResolver implementation.
func (r *Resolver) CommunityDiscussion() generated.CommunityDiscussionResolver {
	return &communityDiscussionResolver{r}
}

// CommunityDiscussionDetail returns generated.CommunityDiscussionDetailResolver implementation.
func (r *Resolver) CommunityDiscussionDetail() generated.CommunityDiscussionDetailResolver {
	return &communityDiscussionDetailResolver{r}
}

// CommunityReview returns generated.CommunityReviewResolver implementation.
func (r *Resolver) CommunityReview() generated.CommunityReviewResolver {
	return &communityReviewResolver{r}
}

// CommunityReviewDetail returns generated.CommunityReviewDetailResolver implementation.
func (r *Resolver) CommunityReviewDetail() generated.CommunityReviewDetailResolver {
	return &communityReviewDetailResolver{r}
}

// Friend returns generated.FriendResolver implementation.
func (r *Resolver) Friend() generated.FriendResolver { return &friendResolver{r} }

// FriendRequest returns generated.FriendRequestResolver implementation.
func (r *Resolver) FriendRequest() generated.FriendRequestResolver { return &friendRequestResolver{r} }

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// GameItem returns generated.GameItemResolver implementation.
func (r *Resolver) GameItem() generated.GameItemResolver { return &gameItemResolver{r} }

// GamePromo returns generated.GamePromoResolver implementation.
func (r *Resolver) GamePromo() generated.GamePromoResolver { return &gamePromoResolver{r} }

// GameReview returns generated.GameReviewResolver implementation.
func (r *Resolver) GameReview() generated.GameReviewResolver { return &gameReviewResolver{r} }

// MarketGameItem returns generated.MarketGameItemResolver implementation.
func (r *Resolver) MarketGameItem() generated.MarketGameItemResolver {
	return &marketGameItemResolver{r}
}

// MarketListing returns generated.MarketListingResolver implementation.
func (r *Resolver) MarketListing() generated.MarketListingResolver { return &marketListingResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// OwnedAvatar returns generated.OwnedAvatarResolver implementation.
func (r *Resolver) OwnedAvatar() generated.OwnedAvatarResolver { return &ownedAvatarResolver{r} }

// OwnedAvatarBorder returns generated.OwnedAvatarBorderResolver implementation.
func (r *Resolver) OwnedAvatarBorder() generated.OwnedAvatarBorderResolver {
	return &ownedAvatarBorderResolver{r}
}

// OwnedBadge returns generated.OwnedBadgeResolver implementation.
func (r *Resolver) OwnedBadge() generated.OwnedBadgeResolver { return &ownedBadgeResolver{r} }

// OwnedChatSticker returns generated.OwnedChatStickerResolver implementation.
func (r *Resolver) OwnedChatSticker() generated.OwnedChatStickerResolver {
	return &ownedChatStickerResolver{r}
}

// OwnedGame returns generated.OwnedGameResolver implementation.
func (r *Resolver) OwnedGame() generated.OwnedGameResolver { return &ownedGameResolver{r} }

// OwnedGameItem returns generated.OwnedGameItemResolver implementation.
func (r *Resolver) OwnedGameItem() generated.OwnedGameItemResolver { return &ownedGameItemResolver{r} }

// OwnedMiniProfileBackground returns generated.OwnedMiniProfileBackgroundResolver implementation.
func (r *Resolver) OwnedMiniProfileBackground() generated.OwnedMiniProfileBackgroundResolver {
	return &ownedMiniProfileBackgroundResolver{r}
}

// OwnedProfileBackground returns generated.OwnedProfileBackgroundResolver implementation.
func (r *Resolver) OwnedProfileBackground() generated.OwnedProfileBackgroundResolver {
	return &ownedProfileBackgroundResolver{r}
}

// OwnedTradingCard returns generated.OwnedTradingCardResolver implementation.
func (r *Resolver) OwnedTradingCard() generated.OwnedTradingCardResolver {
	return &ownedTradingCardResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// ReportRequest returns generated.ReportRequestResolver implementation.
func (r *Resolver) ReportRequest() generated.ReportRequestResolver { return &reportRequestResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// SuspensionList returns generated.SuspensionListResolver implementation.
func (r *Resolver) SuspensionList() generated.SuspensionListResolver {
	return &suspensionListResolver{r}
}

// UnsuspensionRequest returns generated.UnsuspensionRequestResolver implementation.
func (r *Resolver) UnsuspensionRequest() generated.UnsuspensionRequestResolver {
	return &unsuspensionRequestResolver{r}
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// Wishlist returns generated.WishlistResolver implementation.
func (r *Resolver) Wishlist() generated.WishlistResolver { return &wishlistResolver{r} }

type cartResolver struct{ *Resolver }
type communityContentResolver struct{ *Resolver }
type communityContentReviewResolver struct{ *Resolver }
type communityDiscussionResolver struct{ *Resolver }
type communityDiscussionDetailResolver struct{ *Resolver }
type communityReviewResolver struct{ *Resolver }
type communityReviewDetailResolver struct{ *Resolver }
type friendResolver struct{ *Resolver }
type friendRequestResolver struct{ *Resolver }
type gameResolver struct{ *Resolver }
type gameItemResolver struct{ *Resolver }
type gamePromoResolver struct{ *Resolver }
type gameReviewResolver struct{ *Resolver }
type marketGameItemResolver struct{ *Resolver }
type marketListingResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type ownedAvatarResolver struct{ *Resolver }
type ownedAvatarBorderResolver struct{ *Resolver }
type ownedBadgeResolver struct{ *Resolver }
type ownedChatStickerResolver struct{ *Resolver }
type ownedGameResolver struct{ *Resolver }
type ownedGameItemResolver struct{ *Resolver }
type ownedMiniProfileBackgroundResolver struct{ *Resolver }
type ownedProfileBackgroundResolver struct{ *Resolver }
type ownedTradingCardResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reportRequestResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type suspensionListResolver struct{ *Resolver }
type unsuspensionRequestResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type wishlistResolver struct{ *Resolver }
