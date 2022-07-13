package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/mirinda123/mirinda-goweb/package/middleware"
	"github.com/mirinda123/mirinda-goweb/package/mirinda"
	"github.com/stretchr/testify/assert"
)

//测试一下Reader读取太多字节是不是会被卡住
//这个测试不需要启动一个服务器
func TestBodyLimitReader(t *testing.T) {

	testSlice1 := []byte("Hello, World!")

	config := middleware.BodyLimitConfig{
		//设置大小为10bit
		Limit: 2,
	}

	reader := &middleware.BodyLimitReader{
		BodyLimitConfig: config,

		//NopCloser 的原理很简单，就是将一个不带 Close 的 Reader 封装成 ReadCloser
		//这个Close方法是空的什么也没有
		Reader: ioutil.NopCloser(bytes.NewReader(testSlice1)),
	}

	// 如果读取所有的信息，应该会报错：返回 ErrStatusRequestEntityTooLarge
	_, err := ioutil.ReadAll(reader)

	//类型转换
	he := err.(*mirinda.HTTPError)

	assert.Equal(t, http.StatusRequestEntityTooLarge, he.Code)

	// 读取2个字节，应该可以成功
	testSlice2 := make([]byte, 2)
	reader = &middleware.BodyLimitReader{
		BodyLimitConfig: config,

		//NopCloser 的原理很简单，就是将一个不带 Close 的 Reader 封装成 ReadCloser
		//这个Close方法是空的什么也没有
		Reader: ioutil.NopCloser(bytes.NewReader(testSlice2)),
	}

	_, err = ioutil.ReadAll(reader)
	assert.Equal(t, nil, err)
}
