package main

import "testing"

func TestMain(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{`863994e,2020-04-11T22:00:54+08:00,kw7oe,Test 3 [CL] title: Integrate Stripe Payment.,<!-- end -->
		dbc4ecd,2020-04-11T22:00:32+08:00,kw7oe,Test 2,<!-- end -->
		fa5a5f4,2020-04-11T21:46:27+08:00,kw7oe,Hello [CL],title: Fix bugs on web onboarding<!-- end -->`,
			`Changelog for 2020-04-11T22:00:54+08:00:

- Integrate Stripe Payment. (by kw7oe)
- Fix bugs on web onboarding (by kw7oe)`},
		{`863994e,2020-04-11T22:00:54+08:00,kw7oe,Test 3 [CL] title: Integrate Stripe Payment.,<!-- end -->
		dbc4ecd,2020-04-11T22:00:32+08:00,kw7oe,Test 2,<!-- end -->
		fa5a5f4,2020-04-11T21:46:27+08:00,kw7oe,Hello [CL],

		title: Fix bugs on web onboarding<!-- end -->`, `Changelog for 2020-04-11T22:00:54+08:00:

- Integrate Stripe Payment. (by kw7oe)
- Fix bugs on web onboarding (by kw7oe)`},
		{`ca0249a,2020-04-18T14:07:31+08:00,kw7oe,Hello [CL],title: Hello world

title: Hellow world 2
<!-- end -->
fb55031,2020-04-18T14:02:19+08:00,kw7oe,Testing [CL],title: New integration
title: Hello World
<!-- end -->
863994e,2020-04-11T22:00:54+08:00,kw7oe,Test 3 [CL] title: Integrate Stripe Payment.,<!-- end -->
dbc4ecd,2020-04-11T22:00:32+08:00,kw7oe,Test 2,<!-- end -->
fa5a5f4,2020-04-11T21:46:27+08:00,kw7oe,Hello [CL],title: Fix bugs on web onboarding
<!-- end -->`, `Changelog for 2020-04-18T14:07:31+08:00:

- Hello world (by kw7oe)
- New integration (by kw7oe)
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
