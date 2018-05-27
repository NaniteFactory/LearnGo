// Go언어 기초 개인 노트

// Go언어 기초 완전 정복. 이 하나에 모두 요약.
// 다른 언어에 대한 기초가 조금 있는 상태에서,
// Go가 유별나다고 생각하는 것들을 모두 정리함.
// ref) http://golang.site/

// 패키지란 무엇인가
// 패키지는 .go 파일이나 그 바이너리를 담는 폴더를 가리킨다.
// .go 파일의 맨 위에 package 키워드와 함께 명시되는 이 패키지 이름은,
// 이 .go 파일이 속한 폴더(패키지)의 이름과 반드시 일치해야 한다.
//
// 하지만 그런 패키지 규칙에 대해서 main 패키지는 예외다.
// main 패키지는 조금 특별하다.
// main 패키지의 main()은 프로그램의 엔트리 포인트로 취급된다.
package main

// 패키지를 가져다 쓰기
// 불러오는 패키지의 이름은 폴더의 이름(경로)와 같다.
// 그리고 항상 $GOPATH/src/에 상대적이다. (%GOPATH%\src\)
import (
	"MyProject1/extern"
	"fmt"
	duh "fmt"
	"log"
	"os"
	"strings"
	"time"
	"unsafe"
)

func main() {
	// --------------------------------------------\\
	// 패키지의 별칭(alias)
	//
	// 특수한 별칭)
	// 별칭이 _ 언더스코어; 언더바인 경우 패키지의 init()만 호출하고 사용하지 않는다.
	// 별칭이 . 점인 경우 해당 패키지의 이름 없이도 해당 패키지의 함수를 호출할 수 있다.
	//
	duh.Println("hey")
	fmt.Println("hey")
	extern.SomeFunc()
	// 참고) 접근지정자
	//
	// 어떤 패키지에 정의된 것들 중,
	// 패키지 외부에서 가져다 쓸 수 있는 모든 것은,
	// 반드시 첫 글자가 대문자로 명명된다.
	//
	// Go는 접근지정을 변수명을 이용해서 표기한다.
	// 대문자로 시작하는 것: public
	// 소문자로 시작하는 것: non-public

	// --------------------------------------------\\
	// 형변환; 캐스팅; Type conversion
	//
	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)

	var aa int
	var bb uint
	aa = 1
	bb = uint(aa)

	println(str, bytes, str2, aa, bb)

	// --------------------------------------------\\
	// 가변 갯수의 인자 받기 & for range 문 테스트
	//
	foo := func(nums ...int) (sum int, first int, length int) {
		for _, e := range nums {
			sum += e
		}
		first = nums[0]
		length = len(nums)
		return
	}

	a, b, c := foo(1, 2, 3)

	println(a)
	println(b)
	println(c)

	// --------------------------------------------\\
	// 익명 함수 테스트
	//
	// 이름 있는 함수는 함수 안에서 정의할 수 없음.
	// 이름 없는 함수인즉 익명 함수만 함수 안에서 정의할 수 있음.
	// 하지만,
	// 함수 안에서 정의된 익명 함수도 (읽기 좋게) 표기상의 편의를 기했을 뿐,
	// 파이썬처럼 런타임에 함수가 동적으로 만들어지는 것은 결코 아니고,
	// C나 자바처럼 컴파일 타임에 정적으로 정의되는 것이 아닐까 생각함.

	bar := func(a int, b int) int { return a + b }

	println(foot(bar))

	// --------------------------------------------\\
	// 클로저 테스트
	//
	// 클로저는 쉽게 말해 private 멤버 변수를 가지는 객체임. (원큐정리)
	// (자바랑 다르게 JS, Swift, Go 등 언어에서는 함수도 일급객체임.)
	//
	// 함수 덩어리나 코드 블럭이라고도 부르는 클로저는 하나의 객체임.
	// 클로저는 함수를 위장했지만 본질은 그냥 객체임.

	makeClosure := func() func() int {
		var i int
		return func() int {
			i++
			return i
		}
	}

	closure1 := makeClosure()
	closure2 := makeClosure()
	closure3 := makeClosure()

	println("--------------------")
	println(closure1())
	println(closure1())
	println(closure1())
	println("--------------------")
	println(closure2())
	println(closure2())
	println("--------------------")
	println(closure3())
	println("--------------------")

	// --------------------------------------------\\
	// 슬라이스 테스트
	//
	// Go에서 슬라이스와 배열은 서로 완전히 다른 타입이다.
	// 슬라이스는 []int이며, 배열은 [...]int이다. (표기도 서로 다르다.)
	//
	// 위 요약 문장 이상의 불필요한 설명을 덧붙이자면 아래가 있다.
	// 슬라이스는 배열을 마음대로 주무르기 위해서 존재한다.
	// 반면 Go의 배열은 그 자체만으로는 나누는 등의 조작이 어려운 매우 딱딱하고 정적인 자료다.
	// C나 자바와 다르게 Go에서 배열은 크기조차 타입을 구성하는 한 요소다.
	// 다시 말해 예를 들어 [2]int와 [3]int는 서로 전혀 다른 타입이다.
	//
	// Go의 슬라이스와 슬라이싱은 파이썬의 memoryview()와 같다.
	// Go의 슬라이싱은 파이썬과 달리 내부적으로 배열을 생성하지 않는다.
	// 슬라이스는 내부의 어떤 배열에 대한 포인터 정보를 가지고 있을 뿐이다.
	// 슬라이스는 세 가지 정보를 가지고 있다: 머리 포인터, len, cap
	var slice []int
	slice = []int{1, 2, 3, 4}
	slice = slice[0:1]

	println(slice[0])

	// --------------------------------------------\\
	// 구조체(클래스)와 포인터 테스트
	var p *Person
	p = new(Person)
	p.name = "Keith"
	p.age = 10
	println(p.name, p.age)

	// Go에서 객체의 생존 테스트
	//
	// Go 객체가 얼마나 목숨이 질긴가? 얼마나 생존력이 강할까?
	// Aggregation(집합)이 아닌 Composition(구성)의 관계를 이용한 테스트
	//
	// Go는 함수에서 지역변수의 주소를 얻어서 반환하는 이상한 문법이 가능하다.
	// 지역변수의 포인터를 함수 바깥으로 내보내면,
	// 죽었어야 할 지역 개체가 죽지 않고 산다.
	// 컴파일러가 판단해서 스택에 있었어야 할 변수를 힙으로 보내 주는 것이다.
	//
	// 또한 어떤 객체에 대한 포인터가 사라졌다고 해도,
	// 멤버를 포인터로 내보냈다면, 그 객체는 죽지 않고 산다.
	// 예를 들어 의자를 등받이, 바닥, 다리로 구성된 객체로 보았을 때,
	// 의자 전체에 대한 포인터가 없어도 의자 다리를 붙잡으면 의자는 산다.
	// (의자 다리가 의자를 구성하는 요소기 때문이다.)
	//
	// Go언어에서는 프로그래머가 객체의 소멸을 직접 지시할 수는 없다.
	// 하지만 Go언어는 GC(소멸관리자)가 매우 똑똑하다.
	// 어떤 모양이든지 주소값을 떠서 변수로 가지고 있는 객체는 반드시 산다.
	println("--------------------")
	getName1 := func(person Person) *string {
		copy := person
		println("1:: ", &person, &copy, &(copy.name))
		return &(copy.name)
	}
	getName2 := func(person Person) *string {
		println("2:: ", &person, &(person.name))
		return &(person.name)
	}
	getName3 := func(person *Person) *string {
		println("3:: ", person, &(person.name), &(person.age))
		return &(person.name)
	}
	name1 := getName1(*p)
	name2 := getName2(*p)
	name3 := getName3(p)
	ptrUnsafe := unsafe.Pointer(name3)
	ptrUnsafe = unsafe.Pointer(uintptr(ptrUnsafe) + 0x10)
	//
	var ptrVoid interface{}
	ptrVoid = (*int)(ptrUnsafe)
	ptrVoid = ptrVoid.(*int) // 'Type Assertion' 예시
	//
	ptr := (*int)(ptrUnsafe) // 포인터 타입의 'Type Conversion'
	//
	p = nil // 원본 person 객체 포인터 상실시킴
	println(
		p, "\n",
		name1, name2, name3, "\n",
		ptrUnsafe, *(*int)(ptrUnsafe), "\n",
		ptr, *ptr, &ptr, "\n",
		ptrVoid,
	)
	println("--------------------")

	// --------------------------------------------\\
	// 예외처리 (에러)
	//
	// 전체요약: defer와 panic(), recover()는,
	// try-catch-finally문과 같은 용도로 사용하곤 한다.
	//
	// 요약1: error 데이터에 대한 type-switch 문은 try-catch처럼 쓰인다.
	// 1) 이는 panic()과 결합하는 경우 catch절에서 내용 직전의 분기 부분까지만이다.
	// 2) 단독으로 쓰이는 경우 try-catch 전체와 같다.
	//
	// 요약2: panic()은 throw처럼 쓰이거나,
	// catch절 분기에서 catch-finally 내용으로 향하는 통로로 쓰인다.
	// 1) panic()은 throw와 같다. throw와 조금 다른 면이 있는데, 그런 점에서는,
	// 2) panic()을 catch절의 내용과 그 뒤 finally 부분으로 향하는 통로로 봐도 좋다.
	//
	// 요약3: defer는 catch나 finally 부분에 해당한다.
	// 1) defer 단독은 finally와 쓰임새가 같다. (스트림의 Close와 같은 Clean-up 작업에 쓰인다.)
	// 2) 반면 recover()를 수반하는 defer는 catch절의 내용 부분처럼 쓰인다.

	// 환경변수 읽기
	// Go에서는 string이 자바와 달리 reference가 아닌 값이기 때문에,
	// == 연산만으로 값의 비교가 가능하다.
	// (한편 다른 언어와 마찬가지인 것은 Go에서도 string은 Immutable이라는 점이다.)
	var gopath string
	for _, e := range os.Environ() {
		splt := strings.Split(e, "=")
		if splt[0] == "GOPATH" {
			gopath = splt[1]
		}
	}

	// 파일 열기 // defer, panic(), recover()
	file, err := os.Open(gopath + "\\a.txt")
	defer file.Close() // defer: 나중에 닫아 줘. (함수 마지막으로 미룬다.) 매우 편리함.

	// defer 함수. panic 호출시 실행됨. 이것이 없으면 panic()에서 앱이 깨짐.
	defer func() { // defer for exception handling
		// exception catch clause
		// 복구 가능하다면 복구해 보고, 복구 가능한 상태가 맞아서 복구가 되었다면,
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r) // 한 줄 찍어 주고 마무리한다.
		}
	}()

	if err != nil {
		// Type Switch를 이용한 에러 체크
		switch err.(type) {
		default: // no error
			println("ok")
		case *os.PathError:
			panic(err)
			//log.Print("Log my error PathError")
			//log.Fatal(err.Error())
			//fallthrough // cannot use fallthrough in type switch
		case error:
			log.Fatal(err.Error())
		}
	}
	println(file.Name())

	// --------------------------------------------\\
	// Type Switch 문
	//
	var somevar interface{}
	somevar = 10
	switch somevar.(type) {
	case int:
		println("int")
	case string:
		println("string")
		//fallthrough // cannot use fallthrough in type switch
	case interface{}:
		println("interface{}")
	default:
		println("unknown")
	}
	println("--------------------")

	// --------------------------------------------\\
	// '채널'의 전체 요약: Go의 채널은 쉽게 말해 스트림과 같은 것이다.
	//
	// 중요성)
	// Go는 그런 스트림을 얼마나 중요하게 생각했는지,
	// 채널을 위해서 아예 특별히 문법을 만들어 두었다.
	//
	// 기본 연산의 표기)
	// 채널에서 읽거나 쓰는 것(송수신)은 <- 연산자를 이용한다.
	// 송신이냐 수신이냐에 따라서 <-의 위치만 변한다.
	// <-라는 표기만 기억하면 되니까 직관적이다.
	//
	// 파라미터로서의 채널)
	// 함수의 파라미터로 채널을 넘길 때에는, 마찬가지로 <-의 위치를 통해서,
	// 송신 전용 채널과 수신 전용 채널을 구분하여 전달할 수 있다.
	//
	// 채널의 생성과 버퍼링)
	// 버퍼가 없는 채널은 make(chan int)와 같이 생성한다.
	// 자바의 BufferedReader처럼 버퍼 채널을 만들고 싶다면,
	// 생성시에 make(chan int, 2)의 형식처럼 버퍼의 크기를 추가적인 인자로 넘긴다.
	// 버퍼가 있는 채널은 송신시에 수신 여부를 기다리지 않기 때문에 루틴을 블록시키지 않는다.
	//
	// 동기함수)
	// 채널은 기본적으로 동기식 처리를 한다.
	// 즉, 채널에서 값을 읽어올 것을 지시했다면,
	// 해당 처리흐름(루틴; 예를 들면 스레드)은 그 값이 떨어질 때까지 블록킹된다.
	// 그렇기 때문에 자주 Go루틴과 조합되어 쓰인다.
	//
	// select문)
	// 특이하게 채널을 위해서 select문이라는 특수 문법을 지원하는데,
	// 소켓 프로그래밍 등에서 사용하는 select()와 거의 같다.
	// 네트워크 프로그래밍에서는 select()가 여러 연결들을 돌봤다.
	// 반면 Go의 select문은 여러 채널을 돌보고 입력이 들어온 쪽을 봐 준다.
	// switch문 모양이랑 비슷해서 select()의 활용을 구현시에 가독성이 좋다.
	//
	// --------------------------------------------\\
	// 채널 테스트 1
	// - 버퍼가 있는 채널 생성하기.
	// - 채널을 닫기.
	// - 채널에서 읽어들일 때 두 번째 반환값을 통해서 채널이 닫혔는지 확인하기.
	// - (번외) goto와 break의 차이.
	ch1 := make(chan int, 2)
	ch1 <- 1 // 여기서 데드락이 걸릴 것 같지만 그렇지 않는 이유는 버퍼가 있음이다.
	ch1 <- 2
	close(ch1)
EXIT:
	for {
		if element, bSuccessfullyReceived := <-ch1; bSuccessfullyReceived { // 채널 열림 감지
			println(element)
		} else { // 채널 닫힘 감지
			println(element, "exit")
			// goto <-> break 예시
			//goto EXIT // 무한 루프 유발
			break EXIT // goto와 마찬가지로 점프하지만, goto와 다르게 현재 for 루프는 건너뛰게 함
		}
	}
	println("--------------------")
	// --------------------------------------------\\
	// 채널 range문 테스트
	//
	// 채널 전용의 for range문의 용법을 알아두어야 한다.
	ch2 := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 1
		ch2 <- 2
		close(ch2)
	}()
	// for element, bSuccessfullyReceived := range ch2 { // 이거 에러 남.
	// 채널 닫힘 감지시 종료한다. 채널 닫힘 감지시까지만 반복수행한다.
	for element := range ch2 {
		println(element)
	}

	// 끝
} // End of main()

// 델리게이트 테스트
//
// C#에서 유래하는 델리게이트는,
// 함수를 인자로 넘길 수 있도록 하기 위해서,
// 특정한 함수 원형을 타입으로 만드는 것을 의미한다.
//
// 델리게이트는 콜백이나 자바의 리스너 패턴에 비교된다.
// 자바의 리스너 패턴과 구현법은 다르지만 용도는 완전히 동일하다.
// Go는 물론 델리게이트 외에도 인터페이스를 이용하여 이를 자바처럼 구현할 수 있다.
type bam func(int, int) int

// 함수를 인자로 넘겨받는 함수 정의 테스트
//
// C나 Go에서 함수를 인자로 넘길 때에는,
// 넘기는 함수 원형이 엄격하게(Strict) 정의되어 있어야 한다.
//
// 자바스크립트나 파이썬은 변수 타입을 엄격하게 확인하지 않는데,
// 그런 언어들은, 변수의 타입은 아무래도 좋고,
// 변수에 어떤 것이 할당될지가 런타임에 수시로 변할 수 있고,
// 그것 때문에 런타임에 에러 나던 말던,
// 프로그래머의 재량에 맡긴다는 태도다.
//
// 반면 자바에서는 그냥 객체를 넘기기 때문에 가장 쿨한 편이다.
// 자바는 그런 점 때문에,
// (자바 언어 스펙상 자바로 쓴 표현이) 조금 장황해 보이기는 하지만,
// 진정한 객체지향이라고 할 수 있다.
// 자바는 관념적으로 쉽고 어떠한 예외적인 규칙도 없는 직관적인 디자인을 채택했다.
//
//
// 번외) The Strictly Typed와 := 문법에 관한 초보적 고찰.
//
// Go에서 타입을 엄격하게 명시하지 않는 경우는 := 문법뿐인데,
// := 사용시는 해당 변수의 타입을 굳이 적지 않아도 알 수 있기 때문이다.
//
// 그런데 := 문법이 무엇이 좋아서 쓰는 것이며 왜 필요할까?
// := 문법이 없었다면 C처럼 매번 변수 타입을 상단에 선언해 주어야 했을 것이다.
// 이는 매우 장황하고 C에서 오히려 가독성을 떨어뜨리는 원인이었다.
// 매번 선언과 할당이 자기 위치를 지킬 필요는 없다.
// 변수는 항상 타입이 엄격한 것이 알아보기 쉽고, 거기에,
// 변수에 뭔가가 처음으로 할당되는 경우에만 작은 표시를 하는 쪽이,
// 훨씬 편한 것이다.
// :=의 도입으로 간이적인 변수 사용이 훨씬 읽고 쓰기 쉬워졌다.
//
// 참고) Go에서 파이썬과 같은 동적 타입 변수를 필요로 할 때에는,
// interface{}로 표기하는 타입인 빈 인터페이스를 쓸 수 있다.
// 이는 자바의 Object, C/C++의 void*와 같다.
func foot(f bam) int {
	return f(1, 2)
}

// 구조체(클래스)
type Person struct {
	name string
	age  int
}
