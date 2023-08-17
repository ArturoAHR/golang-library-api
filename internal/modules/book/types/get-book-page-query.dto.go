package bookmoduletypes

type GetBookPageQueryDto struct {
	Page string `validate:"numeric"`
}
