package math

import (
	"testing"
)

func TestRound(t *testing.T) {

	//index第三位之後四捨五入
	//碰到9再進位
	name:=0
	got := Round(13.459999, 0.5, 3)
	want:=13.46
	if got != want {
		t.Errorf("Round error, name=%d, got=%f, want=%f", name,got, want)
	}


	//index第4位四捨五入
	name=1
	got = Round(13.459299, 0.2, 4)
	want=13.4593
	if got != want {
		t.Errorf("Round error, name=%d, got=%f, want=%f", name,got, want)
	}



	//index第2位無條件捨去
	name=2
	got = Round(13.459999, 1, 2)
	want=13.45
	if got != want {
		t.Errorf("2 Round error, name=%d, got=%f, want=%f", name ,got, want)
	}

	//index第3位無條件進位
	name=3
	got = Round(13.451999, 0, 3)
	want=13.452
	if got != want {
		t.Errorf("2 Round error, name=%d, got=%f, want=%f", name ,got, want)
	}




	//fmt.Println(int(newVal * 100))
}
func TestRandomInt(t *testing.T) {

	times:=1
	min:=1
	max:=100001
	tm := 0
	tn:=0
	tmx:=0
	tnn:=0
	for i := 0; i < times; i++ {

		r:=RandomInt(min, max)
		if r == max {
			tm += 1
		}else if r==min{
			tn+=1
		}else if r>max{
			tmx+=1
		}else if r<min{
			tnn+=1
		}
	}

	t.Logf("ran max=%d, min=%d, >max=%d, <min=%d", tm,tn,tmx,tnn)
}

func TestCryptoRandomInt(t *testing.T) {

	times:=1000000
	min:=0
	max:=100001
	tm := 0
	tn:=0
	r1:=0
	for i := 0; i < times; i++ {

		r:=CryptoRandomInt(min,max)
		if r == max-1 {
			tm += 1
		}else if r==min{
			tn+=1
		}else if r==1{
			r1+=1
		}
	}

	t.Logf("ran max-1 times=%d, min times=%d, r=1 times=%d", tm,tn,r1)
}