package service

type Service struct {
	repo  Repository
	kafka Kafka
}

type Repository interface {
}

type Kafka interface {
}
