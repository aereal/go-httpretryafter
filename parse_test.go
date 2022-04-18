package httpretryafter

import (
	"reflect"
	"testing"
	"time"
)

func TestParseSeconds(t *testing.T) {
	cases := []struct {
		name    string
		args    string
		want    time.Duration
		wantErr bool
	}{
		{"ok", "60", time.Minute * 1, false},
		{"invalid", "", time.Duration(0), true},
		{"negative", "-10", time.Duration(0), true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseSeconds(c.args)
			if (err != nil) != c.wantErr {
				t.Errorf("error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestParseHTTPDate(t *testing.T) {
	now := time.Now().Round(time.Second)
	aMinuteLater := now.Add(time.Minute)

	cases := []struct {
		name    string
		args    string
		want    time.Time
		wantErr bool
	}{
		{"ok", aMinuteLater.Format(time.RFC1123), aMinuteLater, false},
		{"invalid format", "2020-01-02", time.Time{}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseHTTPDate(c.args)
			if (err != nil) != c.wantErr {
				t.Errorf("error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if !eqTimeInSeconds(got, c.want) {
				t.Errorf("got = %s, want %s", got, c.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	now := time.Now().Round(time.Second)
	orig := nowFunc
	nowFunc = func() time.Time { return now }
	defer func() {
		nowFunc = orig
	}()
	aMinuteLater := now.Add(time.Minute)

	cases := []struct {
		name    string
		args    string
		want    time.Time
		wantErr bool
	}{
		{"seconds/ok", "60", aMinuteLater, false},
		{"http date/ok", aMinuteLater.Format(time.RFC1123), aMinuteLater, false},
		{"invalid", "", time.Time{}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Parse(c.args)
			if (err != nil) != c.wantErr {
				t.Errorf("error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if !eqTimeInSeconds(got, c.want) {
				t.Errorf("got = %s, want %s", got, c.want)
			}
		})
	}
}

func eqTimeInSeconds(x, y time.Time) bool {
	return x.Round(time.Second).Equal(y.Round(time.Second))
}
