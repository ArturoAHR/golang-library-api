package bookmoduletypes

type GetBooksQueryDto struct {
	Page     string `validate:"number"`
	PageSize string `validate:"number"`
}
