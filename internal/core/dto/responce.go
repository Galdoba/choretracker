package dto

type FromServiceResponce interface {
	Content() ChoreContent
	Id() ChoreIdentity
}
