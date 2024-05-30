package playerRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/liangkhwai/go-shop/modules/player"
	"github.com/liangkhwai/go-shop/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PlayerRepositoryService interface {
		IsUniquePlayer(pctx context.Context, email, username string) bool
		InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error)
		FindOnePlayerProfile (pctx context.Context, playerId string) (*player.PlayerProfileBson, error)
		InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error
	}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepositoryService {
	return &playerRepository{db: db}
}

func (r *playerRepository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}

func (r *playerRepository) IsUniquePlayer(pctx context.Context, email, username string) bool {

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player")

	player := new(player.Player)
	if err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
		},
	).Decode(player); err != nil {
		log.Printf("Error: IsUniquePlayer: %s", err.Error())
		return true
	}

	return false

}

func (r *playerRepository) InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)

	col := db.Collection("player")

	playerId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayer: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one player failed")

	}

	return playerId.InsertedID.(primitive.ObjectID), nil
}


func (r *playerRepository) FindOnePlayerProfile (pctx context.Context, playerId string) (*player.PlayerProfileBson, error){
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player")

	result := new(player.PlayerProfileBson)

	if err := col.FindOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(playerId)},
		options.FindOne().SetProjection(bson.M{
			"_id": 1,
			"email":1,
			"username":1,
			"created_at":1,
			"updated_at":1,
		},
	),
	).Decode(result); err != nil{
		log.Printf("Error: FindOnePlayerProfile: %s", err.Error())
		return nil, errors.New("error: player profile not found")
	}

	return result, nil
}


func (r *playerRepository) InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)

	col := db.Collection("player_transactions")

	result,err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayerTransaction: %s", err.Error())
		return  errors.New("error: insert one player transaction failed")

	}
	log.Printf("Result: InsertOnePlayerTransaction: %s", result.InsertedID)

	return  nil
}
