package util

import (
    "testing"
    "math/rand"
)

func TestEncodeAndDecode(t *testing.T) {
    type args struct {
        num int
        n   int
    }

    randNum := rand.Intn(16)
    randN := rand.Intn(16) % 62
    tests := []struct {
        name string
        args args
        want int
    }{
       {name:"64", args:args{123456, 64}, want: 123456},
       {name:"64", args:args{123456, 63}, want: 123456},
       {name:"64", args:args{123456, 32}, want: 123456},
       {name:"64", args:args{341878932, 16}, want: 341878932},
       {name:"64", args:args{341878932, 18}, want: 341878932},
       {name:"64", args:args{33, 31}, want: 33},
       {name:"64", args:args{randNum, randN}, want: randNum},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            encoded := DecimalToAny(tt.args.num, tt.args.n)
            decoded := AnyToDecimal(encoded, tt.args.n)

            if decoded != tt.want {
                t.Errorf("AnyToDecimal() = %v, want %v", decoded, tt.want)
            }
        })
    }
   
}
