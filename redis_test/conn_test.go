package redistest

import(
    "testing"
)

func TestDoGet(t *testing.T){
    DoGet("test")
}

func TestPipeGet(t *testing.T){
    PipeGet("test")
}

func BenchmarkDoGet(b *testing.B){
    for i:=0;i<b.N;i++ {
        DoGet("test")
    }
}

func BenchmarkPipeGet(b *testing.B){
    for i:=0;i<b.N;i++ {
        PipeGet("test")
    }
}
