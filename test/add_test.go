package test

// testsuite test
import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AddSumSuite struct {
	suite.Suite // 第三方的
	testCount   uint32
}

// 这是入口
func TestAddSum(t *testing.T) {
	suite.Run(t, new(AddSumSuite))
}

// 只在开始时执行一次
func (suite *AddSumSuite) SetupSuite() {
	fmt.Println("这里是SetupSuite")
}

// 只在结束时执行一次
func (suite *AddSumSuite) TearDownSuite() {
	fmt.Println("这里是TearDownSuite")
}

// 加了参数才算是实现了对应的接口，加了参数后，才能点进去！怪不得当时想找找不到。
func (suite *AddSumSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
	//fmt.Println("这里是Before Test")
}

func (suite *AddSumSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
	//fmt.Println("这里是After Test")
}

// 每条case开始时都会执行
func (suite *AddSumSuite) SetupTest() {
	fmt.Printf("SetupTest test count:%d\n", suite.testCount)
	//fmt.Println("这里是SetupTest")
}

// 每条用例结束时都会执行
func (suite *AddSumSuite) TearDownTest() {
	suite.testCount++
	fmt.Printf("TearDownTest test count:%d\n", suite.testCount)
	//fmt.Println("这里是TearDownTest")

}

func (suite *AddSumSuite) TestAddSum1() {
	res := sum(10)
	fmt.Printf("sum(10):%d\n", res)
}

func (suite *AddSumSuite) TestAddSum2() {
	res := sum(20)
	fmt.Printf("sum(10):%d\n", res)
}

func (suite *AddSumSuite) TestAddSum3() {
	res := sum(30)
	fmt.Printf("sum(10):%d\n", res)
}
