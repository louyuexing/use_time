# use time

## use
```
func TestUseTime(t *testing.T) {
	ut := New(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(ut.Step().Seconds())
	time.Sleep(2 * time.Second)
	fmt.Println(ut.Step().Seconds())
}
```