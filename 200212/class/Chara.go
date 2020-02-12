package class

//Chara Class
type Chara struct{
	no	int
	hp	int
	ap	int
	af	int
	aw	int
	s1	string
	s2	string
	s3	string
	s4	string
	s5	string
}
func (t Chara) aaa() {
}

//NewChara test
func NewChara () *Chara {
/*
	if age < 0 {
        return nil
    }
    p := new(Person)
    p.Name = name
	p.Age = age
	*/
	p:= new(Chara)
    return p
}
