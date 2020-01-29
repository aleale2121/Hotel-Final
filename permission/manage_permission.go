package permission

import (
	"strings"
)

type permission struct {
	roles   []string
	methods []string
}

type authority map[string]permission

var authorities = authority{
	"/": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/profile": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/newss": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/event": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/map": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/contact": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/rooms": permission{
		roles:   []string{"USER"},
		methods: []string{"GET"},
	},
	"/rate": permission{
		roles:   []string{"USER"},
		methods: []string{"GET"},
	},
	"/about": permission{
		roles:   []string{"USER"},
		methods: []string{"GET"},
	},
	"/login": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/logout": permission{
		roles:   []string{"USER"},
		methods: []string{"POST"},
	},
	"/signup": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/order": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/home": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/user": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/admin": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
}

// HasPermission checks if a given role has permission to access a given route for a given method
func HasPermission(path string, role string, method string) bool {
	if strings.HasPrefix(path, "/admin") {
		path = "/admin"
	}
	perm := authorities[path]
	checkedRole := checkRole(role, perm.roles)
	checkedMethod := checkMethod(method, perm.methods)
	if !checkedRole || !checkedMethod {
		return false
	}
	return true
}

func checkRole(role string, roles []string) bool {
	for _, r := range roles {
		if strings.ToUpper(r) == strings.ToUpper(role) {
			return true
		}
	}
	return false
}

func checkMethod(method string, methods []string) bool {
	for _, m := range methods {
		if strings.ToUpper(m) == strings.ToUpper(method) {
			return true
		}
	}
	return false
}
