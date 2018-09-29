# gtee

极验证golang Sdk

不需要任何第三方依赖


`go get github.com/zcshan/gtee`


# 依赖
```
package main.
import (
    "github.com/zcshan/gtee"
)

var gt3id string = "xxx"
var gt3key string = "xxx"
```

# 注册
```




gteeobj := gtee.NewGeetest(gt3id, gt3key)

gteeobj.Register("unknnow", "unknnow", func(b *gtee.Register_result, str string) {
    if b.Success == 0 {
        fallback = true
    } else {
        fallback = false
    }
    
   

})
```

# 验证
```

geetest_challenge := "xxx"
geetest_validate := "xxx"
geetest_seccode := "xxx"
gteeobj := gtee.NewGeetest(gt3id, gt3key)
fallback = "xxx" //注册方法后的fallback  (建议session存储)
that := ctx
gteeobj.Validate(fallback, geetest_challenge, geetest_validate, geetest_seccode, func(b bool) {
    //b 是否成功
})



```

