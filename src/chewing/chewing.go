package chewing

type Chewing struct {
}

func New() (chewing *Chewing) {
    chewing = new(Chewing)
    return chewing
}
