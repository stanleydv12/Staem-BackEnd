package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/stanleydv12/gqlgen-todos/database"
	"github.com/stanleydv12/gqlgen-todos/graph/generated"
	"github.com/stanleydv12/gqlgen-todos/graph/helper"
	"github.com/stanleydv12/gqlgen-todos/graph/model"
	"strconv"
)

func (r *cartResolver) Game(ctx context.Context, obj *model.Cart) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
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
		ID:       int64(id),
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

	if result.RowsAffected == 0 {
		return "", nil
	}
	return helper.GenerateToken(int(user.ID))
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

func (r *wishlistResolver) Game(ctx context.Context, obj *model.Wishlist) (*model.Game, error) {
	var game model.Game
	db := database.GetInstance()
	db.Where("id = ?", obj.GameID).First(&game)
	return &game, nil
}

// Cart returns generated.CartResolver implementation.
func (r *Resolver) Cart() generated.CartResolver { return &cartResolver{r} }

// GameReview returns generated.GameReviewResolver implementation.
func (r *Resolver) GameReview() generated.GameReviewResolver { return &gameReviewResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// OwnedGame returns generated.OwnedGameResolver implementation.
func (r *Resolver) OwnedGame() generated.OwnedGameResolver { return &ownedGameResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Wishlist returns generated.WishlistResolver implementation.
func (r *Resolver) Wishlist() generated.WishlistResolver { return &wishlistResolver{r} }

type cartResolver struct{ *Resolver }
type gameReviewResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type ownedGameResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type wishlistResolver struct{ *Resolver }
