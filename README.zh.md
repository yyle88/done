# done
package name "done" means function() return success. also means "assert".

most of the time we write a go code, we need to handle the error return value.

it wastes a lot of time and code lines. 

```
res, err := xxx()
if err != nil {
	fatal(...)
}
if res == wrong-value {
	panic(...)
}
// next logic
num, err := yyy(res) //use the res as an arg to call this function
if err != nil {
	fatal(...)
}
if num <= 0 {
	panic(...)
}
// then use the num in next logic
```

so why not just panic when an error occurs(especially in small informal projects).

so this utils can help some developers to override/rush/skip/ignore this problem.
```
res := done.VCE(xxx()).Nice()
num := done.VCE(yyy(res)).Gt(0)
// then use the num in next logic
```
so it can write in one line.
```
num := done.VCE(  yyy(  done.VCE(xxx()).Nice()  )  ).Gt(0)
// then use the num in next logic
```
a ha ha ha! a genius design.

这个包的主要功能就是断言 "assert"， 让你在99.99%不会出错的地方直接用个断言代替 err != nil 的错误处理，而且对于返回 res, err 两个返回值的函数非常有效。

通常我们调用第三方包时，往往就是 `cli, err := xx.NewXx()` 接下来判断没有错误，在调用 `res, err:= cli.Abc()` 的操作，接下来使用 `res.Xyz` 获得信息，但这个函数依然可能是会返回个错误的，毕竟根据go的标准，或者说习惯，err都是一层一层向外返回的。

我当然知道我们开发代码需要遵守规范，但是，在某些非正式的场景里，我们确实希望代码能尽量短些我们就可以使用
```
cli := done.VCE( xx.NewXx() ).Nice()
xyz := done.VCE( cli.Abc() ).Nice().Xyz
```
这样确实是能带来极大的便利。

活学活用，别太死板，而且像 gin 和 kratos 框架在 HandleFunc 和 service 层面通常都是有 recover middleware 的，在几乎不可能出错的地方用断言也没啥问题。

该 panic 就 panic 吧。

code:
```
v_e means check res, err and the result, check with some interface{}/any logic.
vae means check res, err and the result, check with array/slice logic.
vbe means check res, err and the result, check with bool logic.
vce means check res, err and the result, some comparable logic.
vse means check res, err and the result, some function: HasPrefix HasSuffix Contains
vne means check res, err and the result, some function: Gt Lt Gte Lte
```

common function:
```
Done means when param is (res, err) it only check if err is nil(logic is if err != nil {panic}), means success.
Sure/Nice/Some means error is nil and the result is not zero value(0/""/false/nil object/empty slice) and then return the res.
Good/Fine/Safe means error is nil and the result is not zero. only check. not return.
Zero/None/Void means error is nil and the result is zero value.
```

使用几个同义词的目的是，这样开发者就可以根据自己的爱好自由选择，而几乎所有的单词都是四个字母的.

当然也是为了让代码更简洁(其实是强迫症在作祟)，但最终还是没能完全只用四个字母的单词（破功啦）。

Give me stars. thank you!!!
