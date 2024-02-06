package person

// Male 根目录执行 mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock
type Male interface {
	Get(id int64) error
}
