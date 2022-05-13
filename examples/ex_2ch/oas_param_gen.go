// Code generated by ogen, DO NOT EDIT.

package api

type APICaptcha2chcaptchaIDGetParams struct {
	// ID доски, например, b.
	Board OptString
	// Номер треда.
	Thread OptInt
}

type APICaptcha2chcaptchaShowGetParams struct {
	// ID капчи.
	ID string
}

type APICaptchaAppIDPublicKeyGetParams struct {
	// Публичный ключ, для получения напишите admin@2ch.hk с темой
	// письма "Получение ключа для приложения" и ссылкой на
	// ваш клиент.
	PublicKey string
	// ID доски, например, b.
	Board OptString
	// Номер треда.
	Thread OptInt
}

type APICaptchaInvisibleRecaptchaIDGetParams struct {
	// ID доски, например, b.
	Board OptString
	// Номер треда.
	Thread OptInt
}

type APICaptchaRecaptchaIDGetParams struct {
	// ID доски, например, b.
	Board OptString
	// Номер треда.
	Thread OptInt
}

type APIDislikeGetParams struct {
	// ID доски, например, b.
	Board string
	// Номер поста.
	Num int
}

type APILikeGetParams struct {
	// ID доски, например, b.
	Board string
	// Номер поста.
	Num int
}

type APIMobileV2AfterBoardThreadNumGetParams struct {
	// ID доски, например, b.
	Board string
	// Номер треда.
	Thread int
	// Номер поста.
	Num int
}

type APIMobileV2InfoBoardThreadGetParams struct {
	// ID доски, например, b.
	Board string
	// Номер треда.
	Thread int
}

type APIMobileV2PostBoardNumGetParams struct {
	// ID доски, например, b.
	Board string
	// Номер поста.
	Num int
}