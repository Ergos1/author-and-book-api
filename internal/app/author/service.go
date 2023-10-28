package author

type AuthorService struct {
	repo AuthorRepo
}

func NewAuthorService(repo AuthorRepo) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}
