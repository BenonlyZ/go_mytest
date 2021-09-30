
//懒汉，延迟加载，当第一次使用实例时才构建
/* type singleton struct{}
var ins *singleton
func GetIns() *singleton{
	if ins==nil{
		ins=&singleton{}
	}
	return ins
} */

//饿汉，类初始化时就加载完成，无论后面用不用这个实例
/* type singleton struct{}
var ins *singleton=&singleton{}
func GetIns() *singleton{
	return ins
} */

//懒汉枷锁
/* type singleton struct{}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
	mu.Lock()
	defer mu.Unlock()

	if ins == nil {
		ins = &singleton{}
	}
	return ins
} */

//加双重锁
/* type singleton struct{}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}
	return ins
} */

//
/* type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
} */
