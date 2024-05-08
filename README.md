# done
package name "done" means function() return success.

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
val, err := yyy(res) //use the res as an arg to call this function
if err != nil {
	fatal(...)
}
if val == wrong-value {
	panic(...)
}
// then use the val in next logic
```

so why not just panic when an error occurs(especially in small informal projects).

so this utils can help some developers to override/rush/skip/ignore this problem.
```
res := done.VE(xxx()).Nice()
val := done.VE(yyy(res)).Nice()
```
so it can write in one line.
```
val := done.VE(yyy(done.VE(xxx()).Nice())).Nice()
```
a ha ha ha! a genius design.

code:
ve means check res, err and the result is an interface{}/any type.
vae means check res, err and the result is an array/slice type.
vbe means check res, err and the result is a bool type.
vce means check res, err and the result is a comparable type.
vse TODO. means string.
vne TODO. means num.

func:
Done means when param is (res, err) it only check if err is nil(logic is if err != nil {panic}), means success.
Sure/Nice/Some means error is nil and the result is not zero value(0/""/false/nil object/empty slice) and then return the res.
Good/Fine/Safe means error is nil and the result is not zero. only check. not return.
Zero/None/Void means error is nil and the result is zero value.

使用几个同义词的目的是，这样开发者就可以根据自己的爱好自由选择，而几乎所有的单词都是四个字母的，也是为了让代码更简洁(其实是强迫症在作祟)，但最终还是没能完全只用四个字母的单词（破功啦）。
