package model

type student struct {
	Name  string
	Score int
}

func NewStudent(name string, score int) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}

// 可以用getter，setter来设置
