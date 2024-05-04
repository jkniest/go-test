package reflection

import (
	"slices"
	"testing"
)

type Profile struct {
	Picture      string
	CreationYear int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with multiple string fields",
			struct {
				FirstName string
				LastName  string
			}{"Jordan", "Kniest"},
			[]string{"Jordan", "Kniest"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Jordan", 24},
			[]string{"Jordan"},
		},
		{
			"nested fields",
			struct {
				Name    string
				Profile Profile
			}{"Jordan", Profile{"me.png", 1999}},
			[]string{"Jordan", "me.png"},
		},
		{
			"pointer",
			&struct {
				Name string
			}{"Jordan"},
			[]string{"Jordan"},
		},
		{
			"slices",
			[]Profile{
				{"Jordan", 1999},
				{"Jeremy", 2024},
			},
			[]string{"Jordan", "Jeremy"},
		},
		{
			"arrays",
			[2]Profile{
				{"Jordan", 1999},
				{"Jeremy", 2024},
			},
			[]string{"Jordan", "Jeremy"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})

		t.Run("maps", func(t *testing.T) {
			aMap := map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			}

			var got []string
			walk(aMap, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, "Moo")
			assertContains(t, got, "Moo")
		})

		t.Run("channels", func(t *testing.T) {
			aChannel := make(chan Profile)

			go func() {
				aChannel <- Profile{"Jordan", 1999}
				aChannel <- Profile{"Jeremy", 2024}
				close(aChannel)
			}()

			var got []string
			want := []string{"Jordan", "Jeremy"}

			walk(aChannel, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("function", func(t *testing.T) {
			aFunction := func() (Profile, Profile) {
				return Profile{"Jordan", 1999}, Profile{"Jeremy", 2024}
			}

			var got []string
			want := []string{"Jordan", "Jeremy"}

			walk(aFunction, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, want) {
				t.Errorf("got %v, but wanted %v", got, want)
			}
		})
	}
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	for _, x := range haystack {
		if x == needle {
			return
		}
	}

	t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
}
