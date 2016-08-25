package conredis

import "testing"

func TestConnect(t *testing.T){
    r,err := Connect("tcp","127.0.0.1:6379")
    if err != nil {
        t.Errorf("Connect(\"tcp\",\"127.0.0.1:6379\") is error\n",err)
    }
    defer r.Close()
}


func BenchmarkConnect(b *testing.B){
    for i := 0;i < b.N; i++{
        b.StartTimer()
        r,err := Connect("tcp","127.0.0.1:6379")
        if err != nil {
            b.Errorf()
        }
        b.StopTimer()
        r.Close()
    }
}
