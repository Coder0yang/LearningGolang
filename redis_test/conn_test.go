package redistest

import(
    "testing"
)

func TestDoGet(t *testing.T){
    Init("tcp","127.0.0.1:6379")
    str := DoGet("test")
    if str != "1"{
        t.Errorf("DoGet() is faild")
    }
    defer Release()
}

func TestPipeGet(t *testing.T){
    Init("tcp","127.0.0.1:6379")
    str := PipeGet("test")
    if str != "1"{
        t.Errorf("DoGet() is faild")
    }
    defer Release()
}

func BenchmarkDoGet(b *testing.B){
    b.StopTimer()
    Init("tcp","127.0.0.1:6379")
    b.StartTimer()
    for i:=0;i<b.N;i++ {
        DoGet("test")
    }
    defer Release()
}

func BenchmarkPipeGet(b *testing.B){
    b.StopTimer()
    Init("tcp","127.0.0.1:6379")
    b.StartTimer()
    for i:=0;i<b.N;i++ {
        PipeGet("test")
    }
    defer Release()
}
