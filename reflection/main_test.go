package main

import (
	"reflect"
	"testing"
)

type Profile struct {
	City string
	Age  int
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "input string",
			Input: struct {
				Name string
			}{
				Name: "Elaris is a cat",
			},
			ExpectedCalls: []string{"Elaris is a cat"},
		},
		{
			Name: "Multiple string fields",
			Input: struct {
				Name        string
				Description string
			}{
				Name:        "Bob",
				Description: "Bob is a mechanic",
			},
			ExpectedCalls: []string{"Bob", "Bob is a mechanic"},
		},
		{
			Name: "Input with non string value",
			Input: struct {
				Name string
				Age  int
			}{
				Name: "Elaris",
				Age:  25,
			},
			ExpectedCalls: []string{"Elaris"},
		},
		{
			Name: "Non flat struct shape",
			Input: Person{
				Name:    "Elaris",
				Profile: Profile{City: "Chiapas", Age: 33},
			},
			ExpectedCalls: []string{"Elaris", "Chiapas"},
		},
		{
			Name: "Pointer as Input",
			Input: &Person{
				Name: "Carlee",
				Profile: Profile{
					City: "Gastonia",
					Age:  1,
				},
			},
			ExpectedCalls: []string{"Carlee", "Gastonia"},
		},
		{
			Name: "Slices",
			Input: []Person{
				{
					Name: "Jilemon",
					Profile: Profile{
						City: "Tampa",
						Age:  100,
					},
				},
				{
					Name: "Carlee",
					Profile: Profile{
						City: "Gastonia",
						Age:  1,
					},
				},
				{
					Name: "Elaris",
					Profile: Profile{
						City: "Compton",
						Age:  13,
					},
				},
			},
			ExpectedCalls: []string{"Jilemon", "Tampa", "Carlee", "Gastonia", "Elaris", "Compton"},
		},
		{
			Name: "Arrays",
			Input: [2]Person{
				{
					Name: "Elaris",
					Profile: Profile{
						City: "Compton",
						Age:  12,
					},
				},
				{
					Name: "Carlee",
					Profile: Profile{
						City: "Seattle",
						Age:  32,
					},
				},
			},
			ExpectedCalls: []string{"Elaris", "Compton", "Carlee", "Seattle"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string

			walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
			}
		})
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{City: "Berlin", Age: 33}
			aChannel <- Profile{City: "Katowice", Age: 34}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{City: "Berlin", Age: 33}, Profile{City: "Katowice", Age: 34}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
