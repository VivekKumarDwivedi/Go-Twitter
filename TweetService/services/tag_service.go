package services

import(
	"TweetService/dto"
	"TweetService/db/repositories"
)
type TagService interface {
	GetAllTags(limit, offset int) ([]dto.TagResponseDTO, error)
	GetTagWithTweets(name string) (*dto.TagWithTweetsResponseDTO, error)
	GetTopTags(limit int) ([]dto.TagResponseDTO, error)
	DeleteTag(id int) error
}

type TagServiceImpl struct {
	tagRepository repositories.TagRepository
}

func NewTagService(tagRepository repositories.TagRepository) TagService {
	return &TagServiceImpl{tagRepository: tagRepository}
}

func (s *TagServiceImpl) GetAllTags(limit, offset int) ([]dto.TagResponseDTO, error) {
	tags, err := s.tagRepository.GetAllTags(limit, offset)
	if err != nil {
		return nil, err
	}
	var tagDTOs []dto.TagResponseDTO
	for _, tag := range tags {
		tagDTOs = append(tagDTOs, dto.TagResponseDTO{
			ID:   int(tag.ID),
			Name: tag.Name,
		})
	}
	return tagDTOs, nil
}

func (s *TagServiceImpl) GetTagWithTweets(name string) (*dto.TagWithTweetsResponseDTO, error) {
	tweets, err := s.tagRepository.GetTagWithTweets(name)
	if err != nil {
		return nil, err
	}
	var tweetDTOs []dto.TweetResponseDTO
	for _, tweet := range tweets {
		tweetDTOs = append(tweetDTOs, dto.TweetResponseDTO{
			ID:        tweet.ID,
			UserID:    tweet.UserID,
			Tweet:     tweet.Tweet,
			CreatedAt: tweet.CreatedAt,
			UpdatedAt: tweet.UpdatedAt,
		})
	}
	return &dto.TagWithTweetsResponseDTO{
		Name:   name,
		Tweets: tweetDTOs,
	}, nil
}

func (s *TagServiceImpl) GetTopTags(limit int) ([]dto.TagResponseDTO, error) {
	tags, err := s.tagRepository.GetTopTags(limit)
	if err != nil {
		return nil, err
	}
	var tagDTOs []dto.TagResponseDTO
	for _, tag := range tags {
		tagDTOs = append(tagDTOs, dto.TagResponseDTO{
			ID:   int(tag.ID),
			Name: tag.Name,
		})
	}
	return tagDTOs, nil
}

func (s *TagServiceImpl) DeleteTag(id int) error {
	return s.tagRepository.DeleteTag(id)
}
