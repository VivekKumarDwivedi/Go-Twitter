package services

import(
	"TweetService/models"
	"TweetService/db/repositories"
	"TweetService/dto"
	"TweetService/utils"
	"fmt"
)

type TweetService interface{
	CreateTweet(payload *dto.CreateTweetRequestDTO) (*dto.TweetResponseDTO,error)
	GetTweetByID(id int)(*dto.TweetResponseDTO,error)
	GetAllTweets(limit,offset int, userID *int,hashtag *string,search *string) ([]*dto.TweetResponseDTO,error)
	UpdateTweet(id int, payload *dto.UpdateTweetRequestDTO) error
	DeleteTweet(id int) error
}

type TweetServiceImpl struct {
	tweetRepo repositories.TweetRepository
	tagRepo repositories.TagRepository
}

func NewTweetService(tRepo repositories.TweetRepository, tagRepo repositories.TagRepository) TweetService {
	return &TweetServiceImpl{
		tweetRepo: tRepo,
		tagRepo: tagRepo,
	}
}

func (s *TweetServiceImpl) CreateTweet(payload *dto.CreateTweetRequestDTO) (*dto.TweetResponseDTO, error) {
	if len(payload.Tweet) > 280 {
		return nil,fmt.Errorf("tweet exceed maximum length of 280 characters")
	}
	
	// save tweet
	tw:=models.Tweet{
		UserID: payload.UserID,
		Tweet: payload.Tweet,
	}
	savedTweet,err := s.tweetRepo.Create(tw)

	if err != nil {
		return nil,err
	}

	// Extract Hashtags
	hashtags := utils.ExtractHashtags(payload.Tweet)

	for _,tag := range hashtags {
		tagID, err := s.tagRepo.Upsert(tag)

		if err != nil {
			return nil,err
		}
		_ = s.tagRepo.AssociateTweetTag(savedTweet.ID,tagID)
	}

	//Build response DTO

	resp := &dto.TweetResponseDTO{
		ID: int(savedTweet.ID),
		UserID: savedTweet.UserID,
		Tweet: savedTweet.Tweet,
		Tags: hashtags,
		CreatedAt: savedTweet.CreatedAt,
		UpdatedAt: savedTweet.UpdatedAt,
	}
	
	return resp, nil
}

func (s *TweetServiceImpl) GetTweetByID(id int) (*dto.TweetResponseDTO, error) {
	tweet, err := s.tweetRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	return &dto.TweetResponseDTO{
		ID: int(tweet.ID),
		UserID: tweet.UserID,
		Tweet: tweet.Tweet,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}, nil
}

func (s *TweetServiceImpl) GetAllTweets(limit, offset int, userID *int, hashtag *string, search *string) ([]*dto.TweetResponseDTO, error) {
	tweets, err := s.tweetRepo.GetAll(limit, offset, userID, hashtag, search)
	if err != nil {
		return nil, err
	}
	
	var resp []*dto.TweetResponseDTO
	for _, tweet := range tweets {
		resp = append(resp, &dto.TweetResponseDTO{
			ID: int(tweet.ID),
			UserID: tweet.UserID,
			Tweet: tweet.Tweet,
			CreatedAt: tweet.CreatedAt,
			UpdatedAt: tweet.UpdatedAt,
		})
	}
	
	return resp, nil
}

func (s *TweetServiceImpl) UpdateTweet(id int, payload *dto.UpdateTweetRequestDTO) error {
	_, err := s.tweetRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.tweetRepo.UpdateTweet(id, payload.Tweet)
}

func (s *TweetServiceImpl) DeleteTweet(id int) error {
	_, err := s.tweetRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.tweetRepo.DeleteTweet(id)
}
