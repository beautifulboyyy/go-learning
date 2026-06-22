package calculator

import "testing"

// 测试函数必须以 Test 开头，并接收 *testing.T。
func TestAdd(t *testing.T) {
	// 每一项代表一组测试输入和预期结果。
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "两个正数", a: 2, b: 3, want: 5},
		{name: "包含负数", a: -2, b: 3, want: 1},
		{name: "两个零", a: 0, b: 0, want: 0},
	}

	for _, test := range tests {
		// t.Run 会把每一组数据作为一个独立的子测试运行。
		t.Run(test.name, func(t *testing.T) {
			got := Add(test.a, test.b)

			if got != test.want {
				t.Errorf("Add(%d, %d) = %d，期望得到 %d", test.a, test.b, got, test.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool // 是否期望函数返回错误
	}{
		{name: "正常除法", a: 10, b: 2, want: 5, wantErr: false},
		{name: "除数为零", a: 10, b: 0, want: 0, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Divide(test.a, test.b)

			// (err != nil) 会得到 bool，再和 wantErr 比较。
			if (err != nil) != test.wantErr {
				t.Fatalf("Divide(%d, %d) 返回的错误为 %v，wantErr = %v", test.a, test.b, err, test.wantErr)
			}

			if got != test.want {
				t.Errorf("Divide(%d, %d) = %d，期望得到 %d", test.a, test.b, got, test.want)
			}
		})
	}
}
