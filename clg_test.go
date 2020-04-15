package main

import "testing"

func TestMain(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{`863994e,2020-04-11T22:00:54+08:00,kw7oe,Test 3 [CL] title: Integrate Stripe Payment.,
		dbc4ecd,2020-04-11T22:00:32+08:00,kw7oe,Test 2,
		fa5a5f4,2020-04-11T21:46:27+08:00,kw7oe,Hello [CL],title: Fix bugs on web onboarding`,
			`Changelog for 2020-04-11T22:00:54+08:00:

- Integrate Stripe Payment. (by kw7oe)
- Fix bugs on web onboarding (by kw7oe)`},
		{`863994e,2020-04-11T22:00:54+08:00,kw7oe,Test 3 [CL] title: Integrate Stripe Payment.,
		dbc4ecd,2020-04-11T22:00:32+08:00,kw7oe,Test 2,
		fa5a5f4,2020-04-11T21:46:27+08:00,kw7oe,Hello [CL],

		title: Fix bugs on web onboarding`, `Changelog for 2020-04-11T22:00:54+08:00:

		- Integrate Stripe Payment. (by kw7oe)
		- Fix bugs on web onboarding (by kw7oe)`},
	}

	for _, c := range cases {
		got := ExtractChangelog(c.in)
		if got != c.want {
			t.Errorf("ExtractChangelog(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
